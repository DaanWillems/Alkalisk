package model

import (
	"database/sql"
	"log"
	"time"
)

type Comment struct {
	Id        string
	Content   string
	PostId    int
	CreatedAt time.Time
}

type CommentRepository struct {
	db *sql.DB
}

func (r *CommentRepository) CreateComment(content string, post_id int) *Comment {
	db, err := sql.Open("mysql", "root:@/alkaliskdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Exec(`INSERT INTO comments (id, content, post_id, created_at) VALUES (NULL, ?, ?, now());`, content, post_id)
	comment := Comment{}
	comment.CreatedAt = time.Now()
	comment.Content = content
	comment.PostId = post_id

	return &comment
}
