package main

import (
	"log"

	"../pkg/service"
)

const cfgPath = "config/config.json"

func main() {
	cfg, err := service.GetConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	postsLoader, err := service.NewPostsLoaderService(
		cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort, cfg.PGName)
	if err != nil {
		log.Fatal(err)
	}
	if err := postsLoader.InitDB(); err != nil {
		log.Fatal(err)
	}
	for postID := 1; postID <= 10; postID++ {
		if err := postsLoader.DownloadPost(postID); err != nil {
			log.Printf("error: %s\n", err)
		}
	}
	/*
	for {
		if err := newsService.LoadNews(100); err != nil {
			log.Println(err)
		} else {
			log.Println()
		}
		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}*/
}