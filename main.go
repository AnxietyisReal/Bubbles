package main

import (
	botEvents "bubbles/bot/events"
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

func main() {
	loaders.ConnectToDatabase(loaders.TokenLoader("database"))
	client, err := disgo.New(loaders.TokenLoader("bot"),
		bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentGuilds)),

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
