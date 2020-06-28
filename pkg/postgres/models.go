package postgres

import (
	"time"
)

const dbSchema = `
CREATE TABLE posts (
	post_id  INTEGER 
			 PRIMARY KEY,
			 
	title    TEXT 
			 NOT NULL,
			 
	text     TEXT
			 NOT NULL,
			 
	date    TIMESTAMP
			NOT NULL,
			 
	views_count INTEGER
			 NOT NULL
			 CHECK (views_count >= 0),
			 
	comments_count INTEGER
			 NOT NULL
			 CHECK (comments_count >= 0),
			 
	bookmarks_count INTEGER
			 NOT NULL
			 CHECK (bookmarks_count>= 0),
			 
	rating   TEXT
			 NOT NULL,

	author_nickname TEXT
             NOT NULL,

	habs_list TEXT
             NOT NULL,

	tags_list TEXT
             NOT NULL,
);`

type HabrPost struct {
	ID             int       `db:"post_id"`
	Date           time.Time `db:"date"`
	Title          string    `db:"title"`
	Text           string    `db:"text"`
	ViewsCount     int       `db:"views_count"`
	CommentsCount  int       `db:"comments_count"`
	BookmarksCount int       `db:"bookmarks_count"`
	Rating         string    `db:"rating"`
	AuthorNickname string    `db:"author_nickname"`
	HabsList       string    `db:"habs_list"`
	TagsList       string    `db:"tags_list"`
}
