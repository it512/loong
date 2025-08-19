package orm

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/it512/loong/orm/ent"
)

func OpenClient(db *sql.DB, driverName string, opts ...ent.Option) (*ent.Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv := entsql.OpenDB(driverName, db)
		return ent.NewClient(append(opts, ent.Driver(drv))...), nil
	case "pgx":
		drv := entsql.OpenDB(dialect.Postgres, db)
		return ent.NewClient(append(opts, ent.Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}
