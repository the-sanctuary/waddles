package db

import (
	"gorm.io/gorm"
)

const (
	ChannelTypeText  = iota
	ChannelTypeVoice = iota
)

type Topic struct {
	gorm.Model
	DiscordID   string
	Channels    []TopicChannel `gorm:"foreignKey:TopicID"`
	TopicUsers  []TopicUser    `gorm:"foreignKey:TopicID"`
	Tags        []*TopicTag    `gorm:"many2many:topic_tag;"`
	Archived    bool           `gorm:"index"`
	Slug        string         `gorm:"unique"`
	Name        string
	Description string
}

type TopicTag struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
	Topics      []*Topic `gorm:"many2many:topic_tag;"`
}

type TopicChannel struct {
	gorm.Model
	ChannelID string
	Type      int
	TopicID   int
	Topic     Topic
}

type TopicUser struct {
	gorm.Model
	User
	TopicID int
	Topic   Topic
}
