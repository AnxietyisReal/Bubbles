package main

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
)

/*
	This file is used to store the commands for the bot
	and will be sent to Discord API on startup

	If you came here expecting to see code for the commands,
	they are in events/interaction.go
*/

var (
	FALSE      = false
	MinServers = 0
	MaxServers = 20
)

func commandsJSON() (commandList []discord.ApplicationCommandCreate) {
	commandList = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "host-stats",
			Description: "Returns host statistics like OS, CPU Usage, etc",
		},
		discord.SlashCommandCreate{
			Name:         "stats",
			Description:  "Returns the FS server information like players, etc",
			DMPermission: &FALSE,
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionInt{
					Name:        "server",
					Description: "Server ID to get stats for",
					MinValue:    &MinServers,
					MaxValue:    &MaxServers,
					Required:    true,
				},
			},
		},
		discord.SlashCommandCreate{
			Name:         "link",
			Description:  "Link the FS server to the bot",
			DMPermission: &FALSE,
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "server-url",
					Description: "Your FS server's \"Link XML\" URL",
					Required:    true,
				},
				discord.ApplicationCommandOptionInt{
					Name:        "id",
					Description: fmt.Sprintf("Server ID (Pick any number between %v-%v, it's just for identification)", MinServers, MaxServers),
					MinValue:    &MinServers,
					MaxValue:    &MaxServers,
					Required:    true,
				},
			},
		},
		discord.SlashCommandCreate{
			Name:         "unlink",
			Description:  "Unlink the FS server from the bot",
			DMPermission: &FALSE,
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionInt{
					Name:        "id",
					Description: "Server ID to be unlinked",
					MinValue:    &MinServers,
					MaxValue:    &MaxServers,
					Required:    true,
				},
			},
		},
		discord.SlashCommandCreate{
			Name:        "invite",
			Description: "Add the bot to your community server",
		},
		discord.SlashCommandCreate{
			Name:         "database",
			Description:  "[Developer] View the list of server IDs and their respective URLs for this guild",
			DMPermission: &FALSE,
		},
	}
	return commandList
}
