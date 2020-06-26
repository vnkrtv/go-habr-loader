package main

import (
	"log"
	"runtime"
	"time"

	"../pkg/service"
)

const cfgPath = "config/config.json"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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
	if err := postsLoader.CloseDBConn(); err != nil {
		log.Fatal(err)
	}

	postsPerGoroutine := cfg.PostsCount / 100
	finChan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go func(i int, finChan chan int) {
			postsLoader, err := service.NewPostsLoaderService(
				cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort, cfg.PGName)
			if err != nil {
				log.Println(err)
			}
			for postID := 1 + postsPerGoroutine * i; postID <= postsPerGoroutine * (i + 1); postID++ {
				if err := postsLoader.DownloadPost(postID); err != nil {
					log.Printf("error: %s\n", err)
				}
				time.Sleep(time.Millisecond * 355 + time.Millisecond * time.Duration(i) * 10)
			}
			if err := postsLoader.CloseDBConn(); err != nil {
				log.Println(err)
			}
			finChan <- i
		}(i, finChan)
	}
	for i := 0; i < 100; i++ {
		log.Printf("%d goroutine finished\n", <-finChan)
	}
}