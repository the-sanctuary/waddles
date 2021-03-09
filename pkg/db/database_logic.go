package db

func TopicFindAll(db *WadlDB) ([]*Topic, error) {
	var topics []*Topic

	tx := db.
		Find(&topics)

	return topics, tx.Error
}

func TopicFindAllForUser(db *WadlDB, user *User) ([]*Topic, error) {
	var topicUsers []*TopicUser

	tx := db.
		Debug().
		Where("topic_users.discord_id = ?", user.DiscordID).
		Preload("Topic").
		Find(&topicUsers)

	var topics []*Topic

	for _, tu := range topicUsers {
		topics = append(topics, &tu.Topic)
	}

	return topics, tx.Error
}

func TopicFindById(db *WadlDB, id int) (*Topic, error) {
	var topic *Topic

	tx := db.
		First(&topic, id)

	return topic, tx.Error
}

func TopicFindBySlug(db *WadlDB, slug string) (Topic, error) {
	var topic Topic

	tx := db.
		Where("slug = ?", slug).
		First(&topic)

	return topic, tx.Error
}
