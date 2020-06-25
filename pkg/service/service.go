package service

import (
	"../loader"
	pg "../postgres"
)

type PostsLoaderService struct {
	db  *pg.Storage
}

func NewPostsLoaderService(pgUser, pgPass, pgHost, pgPort, pgDBName string) (*PostsLoaderService, error) {
	db, err := pg.OpenConnection(pgUser, pgPass, pgHost, pgPort, pgDBName)
	return &PostsLoaderService{db: db}, err
}

func (s *PostsLoaderService) InitDB() error {
	return s.db.CreateSchema()
}

func (s *PostsLoaderService) DownloadPost(postID int) error {
	post, err := loader.LoadPost(postID)
	if err != nil {
		return err
	}
	return s.db.InsertPost(post)
}
