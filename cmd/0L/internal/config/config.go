package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

const ()

type Server struct {
	Port    string        `json:"port"`
	Host    string        `json:"host"`
	Timeout time.Duration `json:"timeout"`
}

type NatsStreams struct {
	Channel string `json:"channel"`
	Cluster string `json:"cluster"`
	NatsURL string `json:"nats_url"`
}

type Postgres struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	Name    string      `json:"name"`
	Version string      `json:"version"`
	Server  Server      `json:"server"`
	NS      NatsStreams `json:"nats-streams"`
	Pg      Postgres    `json:"postgras"`
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
