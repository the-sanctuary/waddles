package db

import (
	"time"

	"gorm.io/gorm"
)

type UserActivity struct {
	gorm.Model
	CommandCount int
	MessageCount int
	VoiceCount   int
	User
	LastChannelVoiceAppearence *time.Time
	LastChannelTextAppearence  *time.Time
}
