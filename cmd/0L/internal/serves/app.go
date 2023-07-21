package serves

import (
	"log"
	"time"

	"github.com/wb/cmd/0L/internal/broker"
	"github.com/wb/cmd/0L/internal/broker/natstreams"
	"github.com/wb/cmd/0L/internal/cache"
	"github.com/wb/cmd/0L/internal/cache/cachetest"
	"github.com/wb/cmd/0L/internal/database"
	"github.com/wb/cmd/0L/internal/database/testdb"
	"github.com/wb/pkg/erro"
)

const (
	ErrStarted    = "cen't started"
	ErrDistribute = "can't distribute"
)

type app struct {
	broker    broker.Broker
	cache     cache.Cache
	db        database.DataBase
	transport chan string
	ack       chan bool
}

var App app

func (app *app) Start() (err error) {
	defer func() { err = erro.IsError(ErrStarted, err) }()

	list, err := app.db.List()
	if err != nil {
		return err
	}
	err = app.cache.SetList(list)
	if err != nil {
		return err
	}

	return nil
}

func (app *app) Work() {
	go func() {
		err := app.broker.Started()
		if err != nil {
			panic(err)
		}
	}()
	for {
		select {
		case json := <-app.transport:
			var ok bool = true

			err := app.db.Set(json)
			if err != nil {
				log.Println(err)
				ok = false
				app.ack <- ok
				continue
			}

			err = app.cache.Set(json)
			if err != nil {
				log.Println(err)
				ok = false
				app.ack <- ok
				continue
			}

			app.ack <- ok
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// distribute
func (app *app) Distribute(kay string) (json string, err error) {
	defer func() { err = erro.IsError(ErrStarted, err) }()

	json, err = app.cache.Single(kay)
	if err != nil {
		return "", err
	}
	return json, nil
}

func init() {
	var transport chan string = make(chan string)
	var ack chan bool = make(chan bool)
	var db database.DataBase
	var cache cache.Cache
	var broker broker.Broker
	var err error

	db = testdb.NewTestDB()
	// if err != nil {
	// 	panic(err)
	// }

	cache = cachetest.NewCacheTest()
	// if err != nil {
	// 	panic(err)
	// }

	broker, err = natstreams.NewNatsStreams(transport, ack)
	if err != nil {
		panic(err)
	}

	App = app{
		broker:    broker,
		cache:     cache,
		db:        db,
		transport: transport,
		ack:       ack,
	}
}
