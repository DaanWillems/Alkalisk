package model

import (
	"database/sql"
	"log"
)

type Topic struct {
	Id          string
	Title       string
	Description string
}

type TopicRepository struct {
	db *sql.DB
}

func (r *TopicRepository) All(page int, amount int) []Topic {
	db, err := sql.Open("mysql", "root:@/alkaliskdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from topics")
	var topics []Topic
	for rows.Next() {
		topic := Topic{}
		rows.Scan(&topic.Id, &topic.Title, &topic.Description)
		topics = append(topics, topic)
	}
	return topics
}

func (r *TopicRepository) New() *Topic {
	return &Topic{}
}

func (r *TopicRepository) Update(id int) {

}
