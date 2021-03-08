package db

func TopicFindAll(db *WadlDB) []Topic {
	var topics []Topic

	db.
		Find(&topics)

	return topics
}

func TopicFindAllForUser(db *WadlDB, User *User) []Topic {
	var topics []Topic

	db.
		Joins("topic_tag").
		Joins("topic_tags").
		Where("").
		Find(&topics)

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
