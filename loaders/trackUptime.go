package loaders

import "time"

var botUptime time.Time

func BotUptimeTracker() time.Time {
	botUptime = time.Now()
	return botUptime
}
