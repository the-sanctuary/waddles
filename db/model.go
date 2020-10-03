package db

import (
	"time"

	"gorm.io/gorm"
)

type UserActivity struct {
	gorm.Model
	UserID                     string `gorm:"uniqueIndex"`
	MessageCount               int
	LastChannelVoiceAppearence *time.Time
	LastChannelTextAppearence  *time.Time
}
