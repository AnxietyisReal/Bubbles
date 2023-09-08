package main

import "github.com/disgoorg/disgo/discord"

func CommandsJSON() (commandList []discord.ApplicationCommandCreate) {
	commandList = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "ping",
			Description: "Check latency to Discord API from bot",
		},
	}
	return commandList
}
