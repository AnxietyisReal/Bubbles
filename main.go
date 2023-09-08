package main

import (
	botEvents "bubbles/bot/handlers/events"
	"bubbles/bot/loaders"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/json"
	"github.com/disgoorg/log"
)

var (
	BotIntents  = gateway.IntentGuilds | gateway.IntentGuildWebhooks
	commandList = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "ping",
			Description: "Check latency to Discord API from bot",
		},
	}
)

func main() {
	client, err := disgo.New(loaders.LoadBotToken("tokens.json"),
		bot.WithGatewayConfigOpts(gateway.WithIntents(BotIntents)),

		bot.WithEventListenerFunc(botEvents.Ready),
	)
	if err != nil {
		panic(err)
	}
	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), commandList); err != nil {
		log.Errorf("failed to set global commands: %v", err)
	}

	listenForInterrupt()
}

func listenForInterrupt() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

func listenForCommand(e *events.ApplicationCommandInteractionCreate) {
	switch name := e.Data.CommandName(); name {
	case "ping":
		updateInter(e, discord.MessageUpdate{
			Content: json.Ptr("Pong!"),
		})
	}
}

func updateInter(e *events.ApplicationCommandInteractionCreate, messageUpdate discord.MessageUpdate) {
	if _, err := e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate); err != nil {
		log.Errorf("failed to update interaction response: %v", err)
	}
}

func updateErr(e *events.ApplicationCommandInteractionCreate, message string) {
	updateInter(e, discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			{
				Title:       "There was an attempt...",
				Description: message,
				Color:       0x560000,
			},
		},
	})
}
