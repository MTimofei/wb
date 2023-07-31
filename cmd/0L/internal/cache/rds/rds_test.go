package rds_test

import (
	"github.com/wb/cmd/0L/internal/cache"
	"github.com/wb/cmd/0L/internal/cache/rds"
)

var _ cache.Cache = (*rds.Redis)(nil)
