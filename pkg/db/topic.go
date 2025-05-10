package db

import (
	"gorm.io/gorm"
)

const (
	ChannelTypeText = iota
	ChannelTypeVoice
)

type Topic struct {
	gorm.Model
	DiscordRoleID string
	Channels      []TopicChannel `gorm:"foreignKey:TopicID"`
	TopicUsers    []TopicUser    `gorm:"foreignKey:TopicID"`
	Tags          []*TopicTag    `gorm:"many2many:topictag_topic;"`
	Archived      bool           `gorm:"index,default:false;"`
	Slug          string         `gorm:"unique"`
	Name          string
	Description   string
}

type TopicTag struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	Topics      []*Topic `gorm:"many2many:topictag_topic;"`
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
	Active  bool `gorm:"index,default:true;"`
}
