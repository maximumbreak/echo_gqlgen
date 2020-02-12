package models

import (
	"time"
)

// Stamp current time to token
func (x *Token) Stamp() {
	t := time.Now()
	x.LastAccessAt = t.Format("2006-01-02 15:04:05")
	if x.CreatedAt == "" {
		x.CreatedAt = x.LastAccessAt
	}
}
