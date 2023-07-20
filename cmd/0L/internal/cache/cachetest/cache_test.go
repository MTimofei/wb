package cachetest_test

import (
	"github.com/wb/cmd/0L/internal/cache"
	"github.com/wb/cmd/0L/internal/cache/cachetest"
)

var _ cache.Cache = (*cachetest.CacheTest)(nil)
