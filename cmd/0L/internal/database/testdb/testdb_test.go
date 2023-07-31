package testdb_test

import (
	"reflect"
	"testing"

	"github.com/wb/cmd/0L/internal/database"
	"github.com/wb/cmd/0L/internal/database/testdb"
)

var _ database.DataBase = (*testdb.TestDB)(nil)

func Test(t *testing.T) {
	testCases := []struct {
		title   string
		jsonIn  []string
		jsonOut []string
	}{
		{
			"Test",
			[]string{`{"key":"value1"}`, `{"key":"value1"}`},
			[]string{`{"key":"value1"}`, `{"key":"value1"}`},
		},
	}
	db := testdb.NewTestDB()
	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			for _, json := range tC.jsonIn {
				_ = db.Set(json)
			}
			result, _ := db.List()
			if reflect.DeepEqual(tC.jsonOut, result) {
				t.Errorf("got: %v, want: %v", result, tC.jsonOut)
			}
		})
	}
}
