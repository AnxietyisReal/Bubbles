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
		/* bot.WithHTTPServerConfigOpts("50e0cd86cf9f69644afc22050e8f6fec41691e958cf4156d1b749634ab0e785e",
			httpserver.WithAddress(":8080"),
			httpserver.WithURL("/api/bot"),
		), */

		bot.WithEventListenerFunc(botEvents.Ready),
		bot.WithEventListeners(&events.ListenerAdapter{OnApplicationCommandInteraction: botEvents.ListenForCommand}),
	)
	if err != nil {
		panic(err)
	}
	/* if err = client.OpenHTTPServer(); err != nil {
		panic(err)
	} */
	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), commandList); err != nil {
		log.Errorf("failed to set global commands: %v", err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
