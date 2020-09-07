package main

import (
	"log"
	"runtime"
	"time"

	"github.com/LeadNess/go-habr-loader/pkg/service"
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
				time.Sleep(time.Second + time.Millisecond * time.Duration(cfg.Interval + i * 137 % 3313))
				if err := postsLoader.DownloadPost(postID); err != nil {
					log.Printf("error: %s\n", err)
				} else {
					log.Printf("loaded post with post id %d\n", postID)
				}
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