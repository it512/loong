package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Exec struct {
	ent.Schema
}

func (Exec) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.Nil).
			Default(uid).
			Unique().
			Immutable(),

		field.String("exec_id").
			MaxLen(64).
			NotEmpty().
			Immutable().
			Unique().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("inst_id").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("parent_fork_id").
			MaxLen(64).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("fork_id").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("fork_tag").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("out_tag").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("join_tag").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		/*
			field.String("in_tag").
				MaxLen(64).
				Optional().
				SchemaType(map[string]string{
					dialect.MySQL:    "varchar(64)", // Override MySQL.
					dialect.Postgres: "varchar(64)", // Override Postgres.
					dialect.SQLite:   "varchar(64)", // Override Postgres.
				}),
		*/

		field.Int("gw_type").Default(0),
		field.Int("status").Default(0),

		field.Int("version").Default(0),
	}
}

func (Exec) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("exec_id").Unique(),
		index.Fields("inst_id"),
		index.Fields("status"),
		index.Fields("version"),
	}
}

func (Exec) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "loong_exec"},
	}
}
