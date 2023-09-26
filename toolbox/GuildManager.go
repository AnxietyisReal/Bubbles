package toolbox

import (
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/snowflake/v2"
)

func RESTGuild_Name(ID snowflake.ID, BotToken string) string {
	guild, _ := rest.Guilds.GetGuild(rest.NewGuilds(rest.NewClient(BotToken)), ID, false)
	return guild.Name
}
