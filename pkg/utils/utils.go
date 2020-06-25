package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	PGUser     string `json:"pguser"`
	PGPass     string `json:"pgpass"`
	PGName     string `json:"pgname"`
	PGHost     string `json:"pghost"`
	PGPort     string `json:"pgport"`
	VKToken    string `json:"vktoken"`
	Interval   int    `json:"interval"`
	PostsCount int    `json:"posts_count"`
}

func GetConfig(configPath string) (Config, error) {
	var config Config
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(bytes, &config)
	return config, err
}
