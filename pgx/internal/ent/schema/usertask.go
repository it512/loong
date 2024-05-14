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

type UserTask struct {
	ent.Schema
}

func (UserTask) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.Nil).
			Default(uid).
			Unique().
			Immutable(),

		field.String("task_id").
			MaxLen(64).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(NewID).
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

		field.String("exec_id").
			MaxLen(64).
			Optional().
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

		field.String("form_key").
			MaxLen(64).
			Optional().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("act_id").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("act_name").
			MaxLen(64).
			Immutable().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("owner").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("assignee").
			MaxLen(64).
			Optional().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("candidate_users").
			MaxLen(64).
			Optional().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("candidate_groups").
			MaxLen(64).
			Optional().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("operator").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.JSON("input", loong.Var{}).Default(loong.NewVar()),

		field.String("batch_no").
			MaxLen(64).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.Int("result").Default(0).Immutable().Optional(),

		field.Time("start_time").Default(time.Now).Immutable(),
		field.Time("end_time").Optional(),

		field.Int("status").Default(0),

		field.Int("version").Default(0),
	}
}

func (UserTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task_id").Unique(),
		index.Fields("assignee"),
		index.Fields("candidate_groups"),
		index.Fields("candidate_users"),
		index.Fields("batch_no"),
		index.Fields("start_time"),
		index.Fields("status"),
		index.Fields("version"),
	}
}

func (UserTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "loong_user_task"},
	}
}
