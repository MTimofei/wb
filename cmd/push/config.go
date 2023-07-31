package main

// import (
// 	"encoding/json"
// 	"log"
// 	"os"
// )

// type Config struct {
// 	Channel string `json:"channel"`
// 	Cluster string `json:"cluster"`
// 	NatsURL string `json:"nats_url"`
// }

// var config Config

// // parseConfig parses the configuration file and returns a pointer to the Config struct.
// //
// // It reads the contents of the "config/0L/config.json" file and unmarshals it into the Config struct.
// // Returns a pointer to the Config struct or nil if there is an error.
// func init() {
// 	body, err := os.ReadFile("config/0L/config.json")
// 	if err != nil {
// 		log.Fatal("cen't read config.json: ", err)
// 	}

// 	err = json.Unmarshal(body, &config)
// 	if err != nil {
// 		log.Fatal("can't unmarshal config.json: ", err)
// 	}
// }
