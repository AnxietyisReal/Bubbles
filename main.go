package main

import (
	botEvents "bubbles/bot/events"
	"bubbles/bot/loaders"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/httpserver"
	"github.com/disgoorg/log"
)

func main() {
	loaders.ConnectToDatabase(loaders.TokenLoader("database"))
	client, err := disgo.New(loaders.TokenLoader("bot"),
		bot.WithHTTPServerConfigOpts(loaders.TokenLoader("botPublicKey"),
			httpserver.WithAddress(":8080"),
			httpserver.WithURL("/api/interactions"),
		),
		bot.WithEventListeners(&events.ListenerAdapter{OnApplicationCommandInteraction: botEvents.ListenForCommand}),
	)
	if err != nil {
		panic(err)
	}

	if err = client.OpenHTTPServer(); err != nil {
		panic(err)
	}

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), CommandsJSON()); err != nil {
		log.Errorf("failed to set global commands: %v", err)
	}

	fmt.Printf("Client ready!\n")
	fmt.Printf("Running Disgo %v\n", disgo.Version)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
