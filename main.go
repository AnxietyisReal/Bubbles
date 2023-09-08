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
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
)

var BotIntents = gateway.IntentGuilds | gateway.IntentGuildWebhooks

func main() {
	client, err := disgo.New(loaders.LoadBotToken("tokens.json"),
		bot.WithGatewayConfigOpts(gateway.WithIntents(BotIntents)),

		bot.WithEventListenerFunc(botEvents.Ready),
		bot.WithEventListeners(&events.ListenerAdapter{OnApplicationCommandInteraction: botEvents.ListenForCommand}),
	)
	if err != nil {
		panic(err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), CommandsJSON()); err != nil {
		log.Errorf("failed to set global commands: %v", err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
