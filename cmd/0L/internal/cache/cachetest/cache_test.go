package cachetest_test

import (
	"testing"

	"github.com/wb/cmd/0L/internal/cache"
	"github.com/wb/cmd/0L/internal/cache/cachetest"
)

var _ cache.Cache = (*cachetest.CacheTest)(nil)

func Test(t *testing.T) {
	testCases := []struct {
		title     string
		jsonIn    string
		kay       string
		expecting string
	}{
		{
			"Test",
			`{"key":"value1"}`,
			"value1",
			`{"key":"value1"}`,
		},
	}
	db := cachetest.NewCacheTest()

	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			db.Set(tC.jsonIn)
			result, _ := db.Single(tC.kay)
			if tC.expecting != result {
				t.Errorf("got: %v, want: %v", result, tC.expecting)
			}
		})
	}
}
