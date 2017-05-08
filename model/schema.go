package model

import (
	"database/sql"
	"log"
)

var schema = []string{
	`
	create table if not exists posts (
			id                 bigint            not null auto_increment,
			title              varchar(50)       default null,
			content            LONGTEXT          not null,
			created_at         datetime(6)       not null,
			last_comment_at    datetime(6)       not null,
			primary key (id)
		) default charset = utf8mb4;
	`,
	`
	create table if not exists comments (
			id                 bigint            not null auto_increment,
			content            LONGTEXT          not null,
			post_id   		   bigint    	     not null references posts(id),
			created_at         datetime(6)       not null,
			primary key (id),
			index (post_id)
		) default charset = utf8mb4;
	`,
}

var drop = []string{
	`drop table if exists comments cascade`,
	`drop table if exists posts cascade`,
}

var seed = []string{
	`INSERT INTO posts (id, title, content, created_at, last_comment_at) VALUES (NULL, 'Example post #1', 'Example Content', NOW(), NOW());`,
}

func Migrate() {
	db, err := sql.Open("mysql", "root:@/alkaliskdb")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range drop {
		db.Exec(v)
	}

	for _, v := range schema {
		db.Exec(v)
	}

	for _, v := range seed {
		db.Exec(v)
	}
}
