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
					Description: "Server ID (Pick any number, it's just for identification)",
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
					Required:    true,
				},
			},
		},
		discord.SlashCommandCreate{
			Name:        "invite",
			Description: "Add the bot to your community server",
		},
		discord.SlashCommandCreate{
			Name:         "fields",
			Description:  "Returns a list of fields on the FS server",
			DMPermission: &FALSE,
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionInt{
					Name:        "id",
					Description: "Server ID to get fields data",
					Required:    true,
				},
				discord.ApplicationCommandOptionBool{
					Name:        "display-owned",
					Description: "Only show owned fields",
					Required:    false,
				},
			},
		},
	}
	return commandList
}
