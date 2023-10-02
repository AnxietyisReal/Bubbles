package toolbox

import (
	"bubbles/bot/structures"
	"fmt"
	"strings"
	"time"
)

func formatUptime(uptime int) string {
	if uptime == 0 {
		return "just now"
	}
	dur := time.Duration(uptime) * time.Minute
	return dur.String()
}

func GetPlayerInfo(playerArray []structures.DSS_PlayerArray) string {
	var playerInfo string
	for _, player := range playerArray {
		adminStatus := ""
		if player.IsAdmin {
			adminStatus = "- **Admin**"
		}
		if player.IsUsed == false {
			continue
		}
		playerInfo += fmt.Sprintf("%v - `%v` %v\n", player.Name, strings.TrimSuffix(formatUptime(player.Uptime), "0s"), adminStatus)
	}
	return playerInfo
}

func FormatDaytime(dayTime int) string {
	hour := dayTime / 3600 / 1000
	minute := dayTime / 60 / 1000 % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}
