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

type Hg struct {
	ent.Schema
}

func (Hg) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.Nil).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("human_code").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("human_name").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("group_code").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("group_name").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.Int("status").Default(0),
	}
}

func (Hg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("human_code"),
		index.Fields("group_code"),
		index.Fields("status"),
	}
}

func (Hg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "id_hg"},
	}
}
