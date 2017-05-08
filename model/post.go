package model

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id          string
	Title       string
	Content     string
	LastUpdated time.Time
}

type PostRepository struct {
	db *sql.DB
}

func (r *PostRepository) All(page int, amount int) []Post {
	db, err := sql.Open("mysql", "root:@/alkaliskdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from posts")
	var posts []Post
	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Content)
		posts = append(posts, post)
	}
	return posts
}

func (r *PostRepository) Get(id int) *Post {
	db, err := sql.Open("mysql", "root:@/alkaliskdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := r.New()
	err = db.QueryRow("select * from posts where id = ?", id).Scan(&p.Id, &p.Title, &p.Content)
	return p
}

func (r *PostRepository) New() *Post {
	return &Post{}
}

func (r *PostRepository) Update(id int) {

}
