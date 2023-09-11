package botEvents

import (
	"context"
	"fmt"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
)

func Ready(e *events.Ready) {
	fmt.Printf("Ready as %s#%s\n", e.User.Username, e.User.Discriminator)
	fmt.Printf("Cached %d guilds\n", len(e.Guilds))
	fmt.Printf("Running Disgo %v\n", disgo.Version)

	e.Client().SetPresence(context.TODO(), gateway.WithCustomActivity("Streaming your server data"))
}
