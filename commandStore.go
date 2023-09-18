package main

import "github.com/disgoorg/disgo/discord"

/*
	This file is used to store the commands for the bot
	and will be sent to Discord API on startup

	If you came here expecting to see code for the commands,
	they are in events/interaction.go
*/

var FALSE = false

func CommandsJSON() (commandList []discord.ApplicationCommandCreate) {
	commandList = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:         "host-stats",
			Description:  "Returns host statistics like system usage, etc",
			DMPermission: &FALSE,
		},
		discord.SlashCommandCreate{
			Name:         "stats",
			Description:  "Returns the FS server information like players, etc",
			DMPermission: &FALSE,
			/* Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "server",
					Description: "Which FS server to get the stats for",
					Required:    true,
				},
			}, */
		},
		discord.SlashCommandCreate{
			Name:         "link",
			Description:  "Link the FS server to the bot",
			DMPermission: &FALSE,
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "panel-url",
					Description: "Your FS server's \"Link XML\" URL",
					Required:    true,
				},
			},
		},
		discord.SlashCommandCreate{
			Name:         "unlink",
			Description:  "Unlink the FS server from the bot",
			DMPermission: &FALSE,
		},
	}
	return commandList
}
