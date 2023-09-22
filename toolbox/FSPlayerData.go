package toolbox

import (
	"bubbles/bot/structures"
	"fmt"
	"strings"
	"time"
)

func formatUptime(uptime int) string {
	dur := time.Duration(uptime) * time.Minute
	return dur.String()
}

func GetPlayerInfo(playerArray []structures.DSS_PlayerArray) string {
	var playerInfo string
	for _, player := range playerArray {
		adminStatus := ""
		if player.IsAdmin {
			adminStatus = "- **ADMIN**"
		}
		if player.IsUsed == false {
			continue
		}
		playerInfo += fmt.Sprintf("%v - `%v` %v\n", player.Name, strings.TrimSuffix(formatUptime(player.Uptime), "0s"), adminStatus)
	}
	return playerInfo
}
