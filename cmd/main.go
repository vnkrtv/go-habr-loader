package main

import (
	"fmt"
	"log"
	"time"

	"../pkg/loader"
)

const cfgPath = "config/config.json"

func main() {
	cfg, err := service.GetConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	newsService, err := service.NewNewsService(
		cfg.VKToken, cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort, cfg.PGName)
	if err != nil {
		log.Fatal(err)
	}

	groupsScreenNames,err := service.GetGroupsScreenNames(groupsPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := newsService.InitDB(); err != nil {
		log.Fatal(err)
	}
	if err := newsService.AddNewsSources(groupsScreenNames); err != nil {
		log.Fatal(err)
	}
	for {
		if err := newsService.LoadNews(100); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("\n")
		}
		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}
}