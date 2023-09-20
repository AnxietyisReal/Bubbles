package toolbox

import (
	"bubbles/bot/structures"
	"fmt"
	"time"
)

func formatUptime(uptime int) string {
	dur := time.Duration(uptime) * time.Second
	formatted := fmt.Sprintf("%d", int(dur.Seconds()))
	fmt.Println("formatted:", formatted, "dur:", dur)
	return formatted
}

func GetPlayerInfo(playerArray []structures.FSAPI_PlayerArr) string {
	var playerInfo string
	for _, player := range playerArray {
		adminStatus := ""
		if player.IsAdmin {
			adminStatus = "- **[ADMIN]**"
		}
		if player.IsUsed == false {
			continue
		}
		playerInfo += fmt.Sprintf("%v - `%v` %v\n", player.Name, formatUptime(player.Uptime), adminStatus)
	}
	return playerInfo
}
