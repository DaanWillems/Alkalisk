package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id            string
	Title         string
	Content       string
	CreatedAt     time.Time
	LastCommentAt time.Time
}

type PostRepository struct {
	db *sql.DB
}

func (r *PostRepository) All(page int, amount int) []Post {
	db, err := sql.Open("mysql", "root:@/alkaliskdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from posts")
	var posts []Post
	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.LastCommentAt)
		posts = append(posts, post)
	}
	return posts
}

func (r *PostRepository) Get(id int) *Post {
	db, err := sql.Open("mysql", "root:@/alkaliskdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	post := r.New()
	err = db.QueryRow("select * from posts where id = ?", id).Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.LastCommentAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TESt: %v", post.Id)
	return post
}

func (r *PostRepository) New() *Post {
	return &Post{}
}

func (r *PostRepository) Update(id int) {

}
