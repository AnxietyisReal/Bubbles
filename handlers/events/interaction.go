package botEvents

import (
	"bubbles/bot/structures"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/log"
	"github.com/dustin/go-humanize"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zcalusic/sysinfo"
)

var mainEmbedColor = 0xf8add8

func ListenForCommand(e *events.ApplicationCommandInteractionCreate) {
	fmt.Printf("Command requested by %s\n", e.Member().User.Username)
	TRUE := true
	switch name := e.Data.CommandName(); name {
	case "ping":
		if err := e.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("Pong! `%vms`", e.Client().Gateway().Latency().Milliseconds()),
		}); err != nil {
			DumpErrToConsole(err)
		}
		break
	case "stats":
		before, err := cpu.Get()
		if err != nil {
			DumpErrToConsole(err)
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			DumpErrToConsole(err)
			return
		}
		total := float64(after.Total - before.Total)

		memInfo, _ := mem.VirtualMemory()
		var si sysinfo.SysInfo
		si.GetSysInfo()

		if err := e.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				{
					Description: fmt.Sprintf("**OS:** %s\n**Arch:** %v", si.OS.Name, runtime.GOARCH),
					Fields: []discord.EmbedField{
						{
							Name:   "CPU Usage",
							Value:  fmt.Sprintf("User: %.1f %%\nSys: %.1f %%", float64(after.User-before.User)/total*100, float64(after.System-before.System)/total*100),
							Inline: &TRUE,
						},
						{
							Name:   "Memory",
							Value:  fmt.Sprintf("Used: %v\nTotal: %v", humanize.IBytes(memInfo.Used), humanize.IBytes(memInfo.Total)),
							Inline: &TRUE,
						},
						{
							Name:   "\u200b",
							Value:  "\u200b",
							Inline: &TRUE,
						},
						{
							Name:   "Version",
							Value:  fmt.Sprintf("Disgo %v\n%v", disgo.Version, runtime.Version()),
							Inline: &TRUE,
						},
						{
							Name:   "Goroutines",
							Value:  fmt.Sprintf("%v", runtime.NumGoroutine()),
							Inline: &TRUE,
						},
					},
					Color: mainEmbedColor,
				},
			},
		}); err != nil {
			DumpErrToConsole(err)
		}
		break
	case "mp":
		httpClient := &http.Client{
			Timeout: time.Second * 15,
		}

		apiEndpoint := ""
		requ, err := http.NewRequest("GET", apiEndpoint, nil)
		if err != nil {
			log.Errorf("failed to parse the API data: %v", err.Error())
		}
		requ.Header.Add("User-Agent", fmt.Sprintf("Bubbles/%v", runtime.Version()))
		r, err := httpClient.Do(requ)
		if err != nil {
			log.Error(err.Error())
		}
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)

		var res structures.FSAPIRawDataResponse
		json.Unmarshal(body, &res)

		if res.Server.Name == "" ||
			res.Server.MapName == "" ||
			res.Server.Game == "" ||
			res.Server.Version == "" {

			res.Server.Name = "Server is offline"
			res.Server.MapName = "· null ·"
			res.Server.Version = "0.0.0.0"
		}

		playerArray := string("")
		for _, player := range res.Slots.Players {
			if playerArray == "" && res.Slots.Used < 1 {
				playerArray = "*No players online*"
			}
			if player.IsAdmin {
				playerArray = fmt.Sprintf("%v\n%v **[ADMIN]**", playerArray, player.Name)
			} else {
				playerArray = fmt.Sprintf("%v\n%v", playerArray, player.Name)
			}
		}

		if err := e.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				{
					Author: &discord.EmbedAuthor{
						Name: fmt.Sprint(res.Server.Name),
					},
					Fields: []discord.EmbedField{
						{
							Name:   "Version",
							Value:  fmt.Sprintf("%v", res.Server.Version),
							Inline: &TRUE,
						},
						{
							Name:   "Map",
							Value:  fmt.Sprintf("%v", res.Server.MapName),
							Inline: &TRUE,
						},
						{
							Name:  "Players",
							Value: fmt.Sprintf("%v", playerArray),
						},
					},
					Footer: &discord.EmbedFooter{
						Text: fmt.Sprintf("%v/%v", res.Slots.Used, res.Slots.Capacity),
					},
					Color: mainEmbedColor,
				},
			},
		}); err != nil {
			DumpErrToConsole(err)
		}
		break
	}
}

func UpdateInter(e *events.ApplicationCommandInteractionCreate, messageUpdate discord.MessageUpdate) {
	fmt.Print(e.Token())
	if _, err := e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate); err != nil {
		log.Errorf("failed to update interaction response: %v", err)
	}
}

func DumpErrToConsole(err error) {
	log.Errorf("failed to send interaction response: %v", err.Error())
}

func UpdateErr(e *events.ApplicationCommandInteractionCreate, message string) {
	UpdateInter(e, discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			{
				Title:       "There was an attempt...",
				Description: message,
				Color:       0x560000,
			},
		},
	})
}
