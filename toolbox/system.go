package toolbox

import "time"

var timerStart = time.Now()

func GetUptime() string {
	uptime := time.Since(timerStart)
	return uptime.Round(time.Second).String()
}

// I dunno how accurate this is, but it's good enough for me.
