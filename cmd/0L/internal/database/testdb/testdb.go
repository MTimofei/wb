package testdb

import "strings"

type TestDB struct {
	db map[string]string
}

func NewTestDB() *TestDB {
	return &TestDB{
		db: make(map[string]string),
	}
}

func (t *TestDB) Set(json string) error {
	t.db[kay(json)] = json
	return nil
}

func (t *TestDB) List() ([]string, error) {
	var list []string
	for _, v := range t.db {
		list = append(list, v)
	}
	return list, nil
}

func kay(json string) string {
	return strings.Trim(strings.SplitAfterN(json, "\"", 5)[3], "\"")
}

func (t *TestDB) Close() error {
	return nil
}
