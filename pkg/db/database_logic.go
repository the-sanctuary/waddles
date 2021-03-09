package db

func TopicFindAll(db *WadlDB) []Topic {
	var topics []Topic

	db.
		Find(&topics)

	return topics
}

func TopicFindAllForUser(db *WadlDB, user *User) []Topic {
	var topicUsers []TopicUser

	db.
		Debug().
		Where("topic_users.discord_id = ?", user.DiscordID).
		Preload("Topic").
		Find(&topicUsers)

	var topics []Topic
	for _, tu := range topicUsers {
		topics = append(topics, tu.Topic)
	}
	return topics
}

func TopicFindById(db *WadlDB, id int) Topic {
	var topic Topic

	db.
		First(&topic, id)

	return topic
}

func TopicFindBySlug(db *WadlDB, slug string) Topic {
	var topic Topic

	db.
		Where("slug = ?", slug).
		First(&topic)

	return topic
}
