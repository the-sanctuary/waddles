package db

//User holds a DiscordID to be embeded in other models
type User struct {
	DiscordID string `gorm:"uniqueIndex"`
}
