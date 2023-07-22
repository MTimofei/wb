package pgsql

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/wb/cmd/0L/internal/config"
	"github.com/wb/pkg/erro"
)

const (
	ErrPostgres = "cen't init postgres:"
	ErrClose    = "err close db:"
	ErrSet      = "err set:"
	ErrList     = "err cen't get list:"
	ErrSingle   = "err cen't get single:"
)

const (
	set     = `INSERT INTO json_object (order_uid, json_data) VALUES (?::json ->'order_uid', ?)`
	list    = `SELECT json_data FROM json_object`
	single  = `SELECT json_data FROM json_object WHERE order_uid = ?`
	migrate = `CREATE TABLE IF NOT EXISTS json_object (
				order_uid VARCHAR(255) NOT NULL PRIMARY KEY,
				json_data JSON NOT NULL
				);`
)

type Postgres struct {
	db *pg.DB
}

func New() (con *Postgres, err error) {
	defer func(error) { err = erro.IsError(ErrPostgres, err) }(err)

	db := pg.Connect(
		&pg.Options{
			Addr:     config.App.Pg.Addr,
			User:     config.App.Pg.User,
			Password: config.App.Pg.Password,
			Database: config.App.Pg.Database,
		},
	)

	_, err = db.Exec(migrate)
	if err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("connect db:", config.App.Pg.Addr)

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) Close() (err error) {
	defer func(error) { err = erro.IsError(ErrClose, err) }(err)
	err = p.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Set(json string) (err error) {
	defer func(error) { err = erro.IsError(ErrSet, err) }(err)

	_, err = p.db.Exec(set, json, json)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) List() (jsonList []string, err error) {
	defer func(error) { err = erro.IsError(ErrList, err) }(err)

	jsonList = make([]string, 0)

	_, err = p.db.Query(&jsonList, list)
	if err != nil {
		return nil, err
	}

	return jsonList, nil
}

func (p *Postgres) Single(kay string) (json string, err error) {
	defer func(error) { err = erro.IsError(ErrList, err) }(err)

	_, err = p.db.QueryOne(&json, single, kay)
	if err != nil {
		return "", err
	}

	return json, nil
}
