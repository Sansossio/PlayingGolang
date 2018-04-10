package results

import (
	"time"
)

// Global properties
var (
	totalMsg, printInterval int
	totalTime, start        time.Time
	messageCallback         func(string, string)
)
