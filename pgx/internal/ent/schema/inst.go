package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/it512/loong"
)

type Inst struct {
	ent.Schema
}

func (Inst) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.Nil).
			Default(uid).
			Unique().
			Immutable(),

		field.String("inst_id").
			MaxLen(64).
			NotEmpty().
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("proc_id").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("busi_key").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("busi_type").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("starter").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.Time("start_time").Default(time.Now).Immutable(),
		field.Time("end_time").Optional(),

		field.JSON("init", loong.Var{}).Default(loong.NewVar()),

		field.Int("status").Default(0),

		field.Int("version").Default(0),
	}
}

func (Inst) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("inst_id").Unique(),
		index.Fields("busi_key"),
		index.Fields("starter"),
		index.Fields("status"),
		index.Fields("version"),
	}
}

func (Inst) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "loong_inst"},
	}
}
