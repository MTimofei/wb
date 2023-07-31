package pgsql_test

import (
	"github.com/wb/cmd/0L/internal/database"
	"github.com/wb/cmd/0L/internal/database/pgsql"
)

var _ database.DataBase = (*pgsql.Postgres)(nil)

// func Test(t *testing.T) {

// }
