package rds

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/wb/cmd/0L/internal/config"
	"github.com/wb/pkg/erro"
)

const (
	ErrInit    = "can't init redis:"
	ErrSet     = "redis set:"
	ErrSetList = "redis set list:"
	ErrSingle  = "redis single:"
)

type Redis struct {
	rd *redis.Client
}

func New() (rd *Redis, err error) {
	defer func(error) { err = erro.IsError(ErrInit, err) }(err)

	client := redis.NewClient(&redis.Options{
		Addr:     config.App.Rds.Addr,
		Password: "",
		DB:       0,
	})

	fmt.Println("connect redis:", config.App.Rds.Addr)

	return &Redis{rd: client}, nil
}

func (rd *Redis) Set(json string) (err error) {
	defer func(error) { err = erro.IsError(ErrSet, err) }(err)
	_, err = rd.rd.Set(context.Background(), kay(json), json, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rd *Redis) SetList(jsonList []string) (err error) {
	defer func(error) { err = erro.IsError(ErrSet, err) }(err)
	for _, json := range jsonList {
		_, err = rd.rd.Set(context.Background(), kay(json), json, 0).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (rd *Redis) Single(kay string) (json string, err error) {
	defer func(error) { err = erro.IsError(ErrSingle, err) }(err)
	json, err = rd.rd.Get(context.Background(), kay).Result()
	if err != nil {
		return "", err
	}
	return json, nil
}

func kay(json string) string {
	return strings.Trim(strings.SplitAfterN(json, "\"", 5)[3], "\"")
}
