package botEvents

import (
	"bubbles/bot/loaders"
	"bubbles/bot/structures"
	"bubbles/bot/toolbox"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
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

var (
	mainEmbedColor = 0xf8add8
	noPermText     = "You need to have a role with `Manage Server` permission to use this command."
)

func ListenForCommand(e *events.ApplicationCommandInteractionCreate) {
	TRUE := true
	switch name := e.Data.CommandName(); name {
	case "host-stats":
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
					Description: fmt.Sprintf("**OS:** %s", si.OS.Name),
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
							Value:  fmt.Sprintf("Disgo %v\nGo %v", disgo.Version, strings.TrimPrefix(runtime.Version(), "go")),
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
	case "stats":
		httpClient := &http.Client{Timeout: time.Second * 10}
		apiEndpoint := loaders.RetrieveFSServerURL(*e.GuildID())

		requ, err := http.NewRequest("GET", apiEndpoint, nil)
		if err != nil {
			log.Errorf("failed to parse the API data: %v", err.Error())
		}
		requ.Header.Add("User-Agent", fmt.Sprintf("Bubbles/Go %v", strings.TrimPrefix(runtime.Version(), "go")))
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
			res.Server.Version == "" {
			emptyField := "····"
			res.Server.Name = "Server is offline"
			res.Server.MapName += emptyField
			res.Server.Version += emptyField
		}

		playerArray := string("")
		if res.Slots.Used < 1 {
			playerArray = "*No players online*"
		} else {
			playerArray = toolbox.GetPlayerInfo(res.Slots.Players)
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
			DumpErrToInteraction(e, fmt.Errorf("Connection timed out while waiting for a response from the server."))
		}
		break
	case "link":
		if !isGuildManager(e) {
			e.CreateMessage(discord.MessageCreate{
				Content: noPermText,
			})
			return
		}
		str, _ := e.SlashCommandInteractionData().OptString("panel-url")
		doc := loaders.CreateSettings(*e.GuildID(), str)
		if doc != nil {
			log.Infof("Created server settings for %v", *e.GuildID())
			e.CreateMessage(discord.MessageCreate{
				Content: "I have stored the URL for this server.",
			})
		} else {
			log.Infof("Server settings already exist for %v", *e.GuildID())
			e.CreateMessage(discord.MessageCreate{
				Content: "I already have an existing URL stored for this server.",
			})
		}
		break
	case "unlink":
		if !isGuildManager(e) {
			e.CreateMessage(discord.MessageCreate{
				Content: noPermText,
			})
			return
		}
		if doc := loaders.DeleteSettings(*e.GuildID()); doc != nil {
			log.Infof("Deleted server settings for %v", *e.GuildID())
			e.CreateMessage(discord.MessageCreate{
				Content: "I have deleted the URL for this server.",
			})
		} else {
			log.Infof("Server settings do not exist for %v", *e.GuildID())
			e.CreateMessage(discord.MessageCreate{
				Content: "I do not have an existing URL stored for this server.",
			})
		}
	}
}

func DumpErrToConsole(err error) {
	log.Errorf("failed to send interaction response: %v", err.Error())
}

func DumpErrToInteraction(e *events.ApplicationCommandInteractionCreate, err error) {
	if err := e.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Title:       "There was an attempt...",
				Description: fmt.Sprintf("```%v```", err.Error()),
				Color:       0x560000,
			},
		},
	}); err != nil {
		DumpErrToConsole(err)
	}
}

func isGuildManager(e *events.ApplicationCommandInteractionCreate) bool {
	member := e.Member()
	if member.Permissions.Has(discord.PermissionManageGuild) {
		return true
	} else {
		return false
	}
}
