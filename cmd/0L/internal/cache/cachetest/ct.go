package cachetest

import "github.com/wb/cmd/0L/internal/cache"

type CacheTest struct {
	cache map[string][]byte
}

func NewCacheTest() *CacheTest {
	return &CacheTest{
		cache: make(map[string][]byte),
	}
}

func (c *CacheTest) Set(p cache.Payload) error {
	c.cache[p.Key()] = p.Value()

	return nil
}

func (c *CacheTest) SetList(p []cache.Payload) error {
	for _, v := range p {
		c.cache[v.Key()] = v.Value()
	}
	return nil
}

func (c *CacheTest) Single(kay string) (cache.Payload, error) {
	return cache.NewPayload(kay, c.cache[kay]), nil
}
