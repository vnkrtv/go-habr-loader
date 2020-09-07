package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostsStorage interface {
	InsertPost(post HabrPost) error
	InsertPosts(post []HabrPost) error
	UpdatePost(post HabrPost) error
	UpdatePosts(post []HabrPost) error
}

type HabrStorage interface {
	PostsStorage
	CreateSchema() error
}

type Storage struct {
	db *sqlx.DB
}

func OpenConnection(user, password, host, port, dbName string) (*Storage, error) {
	conStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName)
	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return &Storage{db: db}, err
}

func (s *Storage) CloseConnection() error {
	return s.db.Close()
}

func (s *Storage) CreateSchema() error {
	res := s.db.MustExec(dbSchema)
	_, err := res.RowsAffected()
	return err
}

func (s *Storage) InsertPost(post HabrPost) error {
	sql := `
		INSERT INTO 
			posts (post_id, title, text, date, views_count, comments_count, 
			       bookmarks_count, rating, author_nickname) 
		VALUES 
			(:post_id, :title, :text, :date, :views_count, :comments_count, 
			 :bookmarks_count, :rating, :author_nickname)
		ON CONFLICT (post_id)
    		DO UPDATE SET
    			title = :title,
    			text = :text,
    			views_count = :views_count,
    			comments_count = :comments_count,
    			bookmarks_count = :bookmarks_count,
    			rating = :rating`
	if _, err := s.db.NamedExec(sql, &post); err != nil {
		return err
	}
	sql = `
		INSERT INTO
			habs (post_id, hab)
		VALUES
			(:post_id, :hab)`
	for _, hab := range post.Habs {
		if _, err := s.db.NamedExec(sql, &hab); err != nil {
			return err
		}
	}
	sql = `
		INSERT INTO
			tags (post_id, tag)
		VALUES
			(:post_id, :tag)`
	for _, tag := range post.Tags {
		if _, err := s.db.NamedExec(sql, &tag); err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) InsertPosts(posts []HabrPost) error {
	for _, post := range posts {
		if err := s.InsertPost(post); err != nil {
			fmt.Printf("error: %#v", post)
			return err
		}
	}
	return nil
}

func (s *Storage) UpdatePost(post HabrPost) error {
	sql := `
		UPDATE posts SET 
			title = :title, text = :text, 
			views_count = :views_count, comments_count = :comments_count,  
			bookmarks_count = :bookmarks_count, rating = :rating
		WHERE 
			post_id = :post_id AND date = :date`
	_, err := s.db.NamedExec(sql, &post)
	return err
}

func (s *Storage) UpdatePosts(posts []HabrPost) error {
	for _, post := range posts {
		if err := s.UpdatePost(post); err != nil {
			return err
		}
	}
	return nil
}


