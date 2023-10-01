package main

import (
	botEvents "bubbles/bot/events"
	"bubbles/bot/loaders"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/httpserver"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
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

	if loaders.IsCmdsDeployable() {
		log.Infof("Commands deployment is enabled, deploying...")
		if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), commandsJSON()); err != nil {
			log.Errorf("failed to set global commands: %v", err)
		}
	} else {
		log.Infof("Commands deployment is disabled")
	}

	fmt.Printf("Client ready!\n")
	fmt.Printf("Running Disgo %v & Go %v\n", disgo.Version, strings.TrimPrefix(runtime.Version(), "go"))
	client.Rest().CreateWebhookMessage(snowflake.MustParse(loaders.TokenLoader("hookId")), loaders.TokenLoader("hookToken"), discord.WebhookMessageCreate{
		Content: "Container has been restarted.",
	}, true, 0)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
