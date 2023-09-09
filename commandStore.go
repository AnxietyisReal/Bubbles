package main

import "github.com/disgoorg/disgo/discord"

func CommandsJSON() (commandList []discord.ApplicationCommandCreate) {
	commandList = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "ping",
			Description: "Check latency to Discord API from bot",
		},
		discord.SlashCommandCreate{
			Name:        "stats",
			Description: "Returns bot statistics like uptime and etc",
		},
		discord.SlashCommandCreate{
			Name:        "mp",
			Description: "Returns the server data via embed",
		},
	}
	return commandList
}
