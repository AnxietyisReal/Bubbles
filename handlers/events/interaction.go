package botEvents

import (
	"fmt"
	"net/http"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/log"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

var mainEmbedColor = 0xf8add8

func ListenForCommand(e *events.ApplicationCommandInteractionCreate) {
	fmt.Printf("Command requested by %s\n", e.Member().User.Username)
	switch name := e.Data.CommandName(); name {
		case "ping":
			if err := e.CreateMessage(discord.MessageCreate{
				Content: "Pong!",
			}); err != nil {
				DumpErrToConsole(err)
			}
			break
		case "stats":
			cpuInfo, _ := cpu.Info()
			memInfo, _ := mem.VirtualMemory()
			TRUE := true

			if err := e.CreateMessage(discord.MessageCreate{
				Embeds: []discord.Embed{
					{
						Title: "System statistics",
						Fields: []discord.EmbedField{
							{
								Name:   "Processor",
								Value:  fmt.Sprintf("%s", cpuInfo[0].ModelName),
								Inline: &TRUE,
							},
							{
								Name:  "Memory",
								Value: fmt.Sprintf("Used: %v\nTotal: %v", humanize.IBytes(memInfo.Used), humanize.IBytes(memInfo.Total)),
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
			if err := e.CreateMessage(discord.MessageCreate{
				Content: "console.mp4",
			}); err != nil {
				DumpErrToConsole(err)
			}
			pullDataFromAPI("")
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

func pullDataFromAPI(url string) {
	res, err := http.Get(url); if err != nil {
		log.Errorf("failed to pull data from API: %v", err.Error())
	}
	defer res.Body.Close()
	log.Infof("Content-Type: %v", res.Header.Get("Content-Type"))
}