// Service api is the core API service.
package api

import (
	_ "encore.app/api/ent/runtime"

	"encore.app/api/ent"
	"encore.dev/storage/sqldb"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

//encore:service
type Service struct {
	ent *ent.Client
}

func initService() (*Service, error) {
	driver := entsql.OpenDB(dialect.Postgres, db.Stdlib())
	entClient := ent.NewClient(ent.Driver(driver))
	return &Service{
		ent: entClient,
	}, nil
}

// Define a database named 'api', using the database migrations in the
// "./migrations" folder. Encore provisions, migrates, and connects to
// the database automatically on startup.
// Learn more: https://encore.dev/docs/primitives/databases
var db = sqldb.NewDatabase("api", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
