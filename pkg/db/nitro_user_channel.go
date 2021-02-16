package db

import (
	"time"

	"gorm.io/gorm"
)

//NitroUserChannel represents a voice channel registered to a nitro booster
type NitroUserChannel struct {
	gorm.Model
	User
	Name      string
	ChannelID string
	Active    bool
	LastUsed  *time.Time
	Enabled   bool
}
