package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/stan.go"
)

// main is the entry point of the program.
// It establishes a connection to a NATS streaming server,
// publishes a message to a channel asynchronously,
// and sleeps for 1 second before publishing the next message.
func main() {
	// // Define the acknowledgement handler function
	// ackHandler := func(guid string, err error) {
	// 	if err != nil {
	// 		log.Fatal("AckHandler:", err, "\n"+guid)
	// 	}
	// }

	// Connect to NATS streaming server
	sc, err := stan.Connect(
		App.NS.Cluster,
		"pub-client",
		stan.NatsURL(App.NS.NatsURL),
		stan.Pings(3, 5),
		stan.MaxPubAcksInflight(10),
	)
	if err != nil {
		log.Fatal("Connect:", err)
	}
	defer sc.Close() // Close connection

	json, err := os.ReadFile("json/0L/model.json")
	if err != nil {
		log.Fatal(
			"ReadFile: ", err,
		)
	}

	// Publish messages asynchronously
	err = sc.Publish(App.NS.Channel, json)
	if err != nil {
		log.Fatal("PublishAsync:", err)
	}
	fmt.Println("pub: ", time.Now())
}

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

type Config struct {
	Name    string      `json:"name"`
	Version string      `json:"version"`
	Server  Server      `json:"server"`
	NS      NatsStreams `json:"nats-streams"`
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
