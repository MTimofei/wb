package cachetest

import "strings"

type CacheTest struct {
	cache map[string]string
}

func NewCacheTest() *CacheTest {
	return &CacheTest{
		cache: make(map[string]string),
	}
}

func (c *CacheTest) Set(json string) error {
	c.cache[kay(json)] = json

	return nil
}

func (c *CacheTest) SetList(jsonList []string) error {
	for _, json := range jsonList {
		c.cache[kay(json)] = json
	}
	return nil
}

func (c *CacheTest) Single(kay string) (string, error) {
	return c.cache[kay], nil
}

func kay(json string) string {
	return strings.Trim(strings.SplitAfterN(json, "\"", 5)[3], "\"")
}
