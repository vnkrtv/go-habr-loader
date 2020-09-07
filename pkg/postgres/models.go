package postgres

import (
	"time"
)

const dbSchema = `
CREATE TABLE IF NOT EXISTS posts (
	post_id  INTEGER 
			 PRIMARY KEY,
			 
	title    TEXT 
			 NOT NULL,
			 
	text     TEXT
			 NOT NULL,
			 
	date     TIMESTAMP
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
             NOT NULL
);

CREATE TABLE IF NOT EXISTS habs (
	post_id  INTEGER
			 NOT NULL
			 DEFAULT 0,

	hab      TEXT
			 NOT NULL,

	CONSTRAINT fk_post FOREIGN KEY (post_id)
		REFERENCES posts(post_id)
			ON UPDATE CASCADE
			ON DELETE SET DEFAULT,

	CONSTRAINT pk_habs PRIMARY KEY(post_id, hab)
);

CREATE TABLE IF NOT EXISTS tags (
	post_id  INTEGER
			 NOT NULL
			 DEFAULT 0,
	
	tag      TEXT
			 NOT NULL,
	
	CONSTRAINT fk_post FOREIGN KEY (post_id)
		REFERENCES posts(post_id)
			ON UPDATE CASCADE
			ON DELETE SET DEFAULT,

	CONSTRAINT pk_tags PRIMARY KEY(post_id, tag)
);`

type Hab struct {
	PostID         int       `db:"post_id"`
	Hab            string    `db:"hab"`
}

type Tag struct {
	PostID         int       `db:"post_id"`
	Tag            string    `db:"tag"`
}

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
	Habs           []Hab
	Tags           []Tag
}
