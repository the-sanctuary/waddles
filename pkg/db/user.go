package db

type User struct {
	DiscordID string `gorm:"uniqueIndex"`
}
