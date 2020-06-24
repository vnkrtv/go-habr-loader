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

type NewsStorage interface {
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

func (s *Storage) CreateSchema() error {
	res := s.db.MustExec(dbSchema)
	_, err := res.RowsAffected()
	return err
}

func (s *Storage) InsertPost(post HabrPost) error {
	sql := `
		INSERT INTO 
			posts (post_id, title, text, date, views_count, comments_count, bookmarks_count, rating) 
		VALUES 
			(:post_id, :title, :text, :date, :views_count, :comments_count, :bookmarks_count, :rating)`
	_, err := s.db.NamedExec(sql, &post)
	return err
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


