package db

import (
	"time"

	"gorm.io/gorm"
)

//UserActivity stores any per-user statistics
type UserActivity struct {
	gorm.Model
	CommandCount int
	MessageCount int
	VoiceCount   int
	User
	LastChannelVoiceAppearence *time.Time
	LastChannelTextAppearence  *time.Time
}

//NicknameUpdate stores ever nickname change in the server
type NicknameUpdate struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"` //we want to easyily keep track of the last one
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User
	Nickname string
}
