package botEvents

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/log"
)

func ListenForCommand(e *events.ApplicationCommandInteractionCreate) {
	fmt.Printf("Command requested by %s\n", e.Member().User.Username)
	switch name := e.Data.CommandName(); name {
	case "ping":
		if err := e.CreateMessage(discord.MessageCreate{
			Content: "Pong!",
		}); err != nil {
			DumpErrToConsole(err)
		}
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
