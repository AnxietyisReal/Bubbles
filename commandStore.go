package main

import "github.com/disgoorg/disgo/discord"

func CommandsJSON() (commandList []discord.ApplicationCommandCreate) {
	commandList = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "ping",
			Description: "Returns the Discord API latency",
		},
		discord.SlashCommandCreate{
			Name:        "stats",
			Description: "Returns host statistics like system usage, etc",
		},
		discord.SlashCommandCreate{
			Name:        "mp",
			Description: "Returns the server data via embed",
		},
	}
	return commandList
}
