package config

import (
	"encoding/json"
	"log"
	"os"
)

const ()

type Config struct {
	Channel string `json:"channel"`
	Cluster string `json:"cluster"`
	NatsURL string `json:"nats_url"`
}

var App Config

func init() {
	body, err := os.ReadFile("config/0L/config.json")
	if err != nil {
		log.Fatal("cen't read config.json: ", err)
	}

	err = json.Unmarshal(body, &App)
	if err != nil {
		log.Fatal("can't unmarshal config.json: ", err)
	}
}
