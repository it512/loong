// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/it512/loong/pgx/internal/ent/exec"
)

// Exec is the model entity for the Exec schema.
type Exec struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ExecID holds the value of the "exec_id" field.
	ExecID string `json:"exec_id,omitempty"`
	// InstID holds the value of the "inst_id" field.
	InstID string `json:"inst_id,omitempty"`
	// ParentForkID holds the value of the "parent_fork_id" field.
	ParentForkID string `json:"parent_fork_id,omitempty"`
	// ForkID holds the value of the "fork_id" field.
	ForkID string `json:"fork_id,omitempty"`
	// ForkTag holds the value of the "fork_tag" field.
	ForkTag string `json:"fork_tag,omitempty"`
	// OutTag holds the value of the "out_tag" field.
	OutTag string `json:"out_tag,omitempty"`
	// JoinTag holds the value of the "join_tag" field.
	JoinTag string `json:"join_tag,omitempty"`
	// GwType holds the value of the "gw_type" field.
	GwType int `json:"gw_type,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Version holds the value of the "version" field.
	Version      int `json:"version,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exec) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exec.FieldGwType, exec.FieldStatus, exec.FieldVersion:
			values[i] = new(sql.NullInt64)
		case exec.FieldExecID, exec.FieldInstID, exec.FieldParentForkID, exec.FieldForkID, exec.FieldForkTag, exec.FieldOutTag, exec.FieldJoinTag:
			values[i] = new(sql.NullString)
		case exec.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exec fields.
func (e *Exec) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exec.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case exec.FieldExecID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exec_id", values[i])
			} else if value.Valid {
				e.ExecID = value.String
			}
		case exec.FieldInstID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inst_id", values[i])
			} else if value.Valid {
				e.InstID = value.String
			}
		case exec.FieldParentForkID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_fork_id", values[i])
			} else if value.Valid {
				e.ParentForkID = value.String
			}
		case exec.FieldForkID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fork_id", values[i])
			} else if value.Valid {
				e.ForkID = value.String
			}
		case exec.FieldForkTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fork_tag", values[i])
			} else if value.Valid {
				e.ForkTag = value.String
			}
		case exec.FieldOutTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field out_tag", values[i])
			} else if value.Valid {
				e.OutTag = value.String
			}
		case exec.FieldJoinTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field join_tag", values[i])
			} else if value.Valid {
				e.JoinTag = value.String
			}
		case exec.FieldGwType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gw_type", values[i])
			} else if value.Valid {
				e.GwType = int(value.Int64)
			}
		case exec.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				e.Status = int(value.Int64)
			}
		case exec.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				e.Version = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Exec.
// This includes values selected through modifiers, order, etc.
func (e *Exec) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// Update returns a builder for updating this Exec.
// Note that you need to call Exec.Unwrap() before calling this method if this Exec
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exec) Update() *ExecUpdateOne {
	return NewExecClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Exec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Exec) Unwrap() *Exec {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exec is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exec) String() string {
	var builder strings.Builder
	builder.WriteString("Exec(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("exec_id=")
	builder.WriteString(e.ExecID)
	builder.WriteString(", ")
	builder.WriteString("inst_id=")
	builder.WriteString(e.InstID)
	builder.WriteString(", ")
	builder.WriteString("parent_fork_id=")
	builder.WriteString(e.ParentForkID)
	builder.WriteString(", ")
	builder.WriteString("fork_id=")
	builder.WriteString(e.ForkID)
	builder.WriteString(", ")
	builder.WriteString("fork_tag=")
	builder.WriteString(e.ForkTag)
	builder.WriteString(", ")
	builder.WriteString("out_tag=")
	builder.WriteString(e.OutTag)
	builder.WriteString(", ")
	builder.WriteString("join_tag=")
	builder.WriteString(e.JoinTag)
	builder.WriteString(", ")
	builder.WriteString("gw_type=")
	builder.WriteString(fmt.Sprintf("%v", e.GwType))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", e.Status))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", e.Version))
	builder.WriteByte(')')
	return builder.String()
}

// Execs is a parsable slice of Exec.
type Execs []*Exec
