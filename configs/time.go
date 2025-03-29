package configs

import "time"

var StartedAt *time.Time

func TimeStart() {
	now := time.Now()
	StartedAt = &now
}
