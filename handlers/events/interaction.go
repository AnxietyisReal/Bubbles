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
		println("im here, ayoo")
		SendInter(e, discord.InteractionResponse{
			Type: discord.InteractionResponseTypeCreateMessage,
			Data: discord.NewMessageCreateBuilder().SetContent("Pong!").Build(),
		})
		break
	}
}

func SendInter(e *events.ApplicationCommandInteractionCreate, interResponse discord.InteractionResponse) {
	if err := e.Client().Rest().CreateInteractionResponse(e.ApplicationID(), e.Token(), interResponse); err != nil {
		log.Errorf("failed to send interaction response: %v", err)
	}
}

func UpdateInter(e *events.ApplicationCommandInteractionCreate, messageUpdate discord.MessageUpdate) {
	fmt.Print(e.Token())
	if _, err := e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate); err != nil {
		log.Errorf("failed to update interaction response: %v", err)
	}
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
