package db

import (
	"time"

	"gorm.io/gorm"
)

type UserActivity struct {
	gorm.Model
	UserID                     string `gorm:"uniqueIndex"`
	MessageCount               int
	CommandCount               int
	VoiceCount                 int
	LastChannelVoiceAppearence *time.Time
	LastChannelTextAppearence  *time.Time
}

//NitroUserChannel represents a voice channel registered to a nitro booster
type NitroUserChannel struct {
	gorm.Model
	Name      string
	UserID    string
	ChannelID string
	Active    bool
	LastUsed  *time.Time
	Enabled   bool
}
