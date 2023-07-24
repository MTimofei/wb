package serves

import (
	"os"

	"github.com/wb/cmd/0L/internal/broker"
	"github.com/wb/cmd/0L/internal/broker/natstreams"
	"github.com/wb/cmd/0L/internal/cache"
	"github.com/wb/cmd/0L/internal/cache/rds"
	"github.com/wb/cmd/0L/internal/config"
	"github.com/wb/cmd/0L/internal/database"
	"github.com/wb/cmd/0L/internal/database/pgsql"
	"github.com/wb/pkg/erro"
	"github.com/wb/pkg/logg"
	"golang.org/x/exp/slog"
)

const (
	ErrStarted    = "cen't started"
	ErrDistribute = "can't distribute"
)

const (
	ErrDuplicate = "err set::ERROR #23505 duplicate key value violates unique constraint \"json_object_pkey\""
)

type app struct {
	broker    broker.Broker
	cache     cache.Cache
	db        database.DataBase
	transport chan string
	ack       chan bool
	Log       *slog.Logger
	ok        bool
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

func (app *app) Close() {
	app.broker.Unsubscribe()
	app.broker.Close()
	app.ok = false
	app.db.Close()
}

func (app *app) Work() {
	//defer app.close()

	go func() {
		err := app.broker.Started()
		if err != nil {
			panic(err)
		}
	}()
	for app.ok {
		select {
		case json := <-app.transport:
			var ok bool = true

			err := app.db.Set(json)
			if err != nil {
				if err.Error() == ErrDuplicate {
					app.ack <- ok
					continue
				}
				app.Log.Error("", logg.Err(err))
				ok = false
				app.ack <- ok
				continue
			}

			err = app.cache.Set(json)
			if err != nil {
				app.Log.Error("", logg.Err(err))
				ok = false
				app.ack <- ok
				continue
			}

			app.ack <- ok
		default:
			// time.Sleep(100 * time.Millisecond)
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

	log := logg.New(config.App.Env)

	db, err = pgsql.New()
	if err != nil {
		log.Error("", logg.Err(err))
		os.Exit(1)
	}

	cache, err = rds.New()
	if err != nil {
		log.Error("", logg.Err(err))
		os.Exit(1)
	}

	broker, err = natstreams.New(transport, ack)
	if err != nil {
		log.Error("", logg.Err(err))
		os.Exit(1)
	}

	App = app{
		broker:    broker,
		cache:     cache,
		db:        db,
		transport: transport,
		ack:       ack,
		Log:       log,
		ok:        true,
	}
}
