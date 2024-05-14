// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/it512/loong"
	"github.com/it512/loong/pgx/internal/ent/predicate"
	"github.com/it512/loong/pgx/internal/ent/usertask"
)

// UserTaskUpdate is the builder for updating UserTask entities.
type UserTaskUpdate struct {
	config
	hooks    []Hook
	mutation *UserTaskMutation
}

// Where appends a list predicates to the UserTaskUpdate builder.
func (utu *UserTaskUpdate) Where(ps ...predicate.UserTask) *UserTaskUpdate {
	utu.mutation.Where(ps...)
	return utu
}

// SetOwner sets the "owner" field.
func (utu *UserTaskUpdate) SetOwner(s string) *UserTaskUpdate {
	utu.mutation.SetOwner(s)
	return utu
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (utu *UserTaskUpdate) SetNillableOwner(s *string) *UserTaskUpdate {
	if s != nil {
		utu.SetOwner(*s)
	}
	return utu
}

// ClearOwner clears the value of the "owner" field.
func (utu *UserTaskUpdate) ClearOwner() *UserTaskUpdate {
	utu.mutation.ClearOwner()
	return utu
}

// SetOperator sets the "operator" field.
func (utu *UserTaskUpdate) SetOperator(s string) *UserTaskUpdate {
	utu.mutation.SetOperator(s)
	return utu
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (utu *UserTaskUpdate) SetNillableOperator(s *string) *UserTaskUpdate {
	if s != nil {
		utu.SetOperator(*s)
	}
	return utu
}

// ClearOperator clears the value of the "operator" field.
func (utu *UserTaskUpdate) ClearOperator() *UserTaskUpdate {
	utu.mutation.ClearOperator()
	return utu
}

// SetInput sets the "input" field.
func (utu *UserTaskUpdate) SetInput(l loong.Var) *UserTaskUpdate {
	utu.mutation.SetInput(l)
	return utu
}

// SetEndTime sets the "end_time" field.
func (utu *UserTaskUpdate) SetEndTime(t time.Time) *UserTaskUpdate {
	utu.mutation.SetEndTime(t)
	return utu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (utu *UserTaskUpdate) SetNillableEndTime(t *time.Time) *UserTaskUpdate {
	if t != nil {
		utu.SetEndTime(*t)
	}
	return utu
}

// ClearEndTime clears the value of the "end_time" field.
func (utu *UserTaskUpdate) ClearEndTime() *UserTaskUpdate {
	utu.mutation.ClearEndTime()
	return utu
}

// SetStatus sets the "status" field.
func (utu *UserTaskUpdate) SetStatus(i int) *UserTaskUpdate {
	utu.mutation.ResetStatus()
	utu.mutation.SetStatus(i)
	return utu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (utu *UserTaskUpdate) SetNillableStatus(i *int) *UserTaskUpdate {
	if i != nil {
		utu.SetStatus(*i)
	}
	return utu
}

// AddStatus adds i to the "status" field.
func (utu *UserTaskUpdate) AddStatus(i int) *UserTaskUpdate {
	utu.mutation.AddStatus(i)
	return utu
}

// SetVersion sets the "version" field.
func (utu *UserTaskUpdate) SetVersion(i int) *UserTaskUpdate {
	utu.mutation.ResetVersion()
	utu.mutation.SetVersion(i)
	return utu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (utu *UserTaskUpdate) SetNillableVersion(i *int) *UserTaskUpdate {
	if i != nil {
		utu.SetVersion(*i)
	}
	return utu
}

// AddVersion adds i to the "version" field.
func (utu *UserTaskUpdate) AddVersion(i int) *UserTaskUpdate {
	utu.mutation.AddVersion(i)
	return utu
}

// Mutation returns the UserTaskMutation object of the builder.
func (utu *UserTaskUpdate) Mutation() *UserTaskMutation {
	return utu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (utu *UserTaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, utu.sqlSave, utu.mutation, utu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utu *UserTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := utu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (utu *UserTaskUpdate) Exec(ctx context.Context) error {
	_, err := utu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utu *UserTaskUpdate) ExecX(ctx context.Context) {
	if err := utu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utu *UserTaskUpdate) check() error {
	if v, ok := utu.mutation.Owner(); ok {
		if err := usertask.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "UserTask.owner": %w`, err)}
		}
	}
	if v, ok := utu.mutation.Operator(); ok {
		if err := usertask.OperatorValidator(v); err != nil {
			return &ValidationError{Name: "operator", err: fmt.Errorf(`ent: validator failed for field "UserTask.operator": %w`, err)}
		}
	}
	return nil
}

func (utu *UserTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := utu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usertask.Table, usertask.Columns, sqlgraph.NewFieldSpec(usertask.FieldID, field.TypeUUID))
	if ps := utu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if utu.mutation.ExecIDCleared() {
		_spec.ClearField(usertask.FieldExecID, field.TypeString)
	}
	if utu.mutation.FormKeyCleared() {
		_spec.ClearField(usertask.FieldFormKey, field.TypeString)
	}
	if utu.mutation.ActNameCleared() {
		_spec.ClearField(usertask.FieldActName, field.TypeString)
	}
	if value, ok := utu.mutation.Owner(); ok {
		_spec.SetField(usertask.FieldOwner, field.TypeString, value)
	}
	if utu.mutation.OwnerCleared() {
		_spec.ClearField(usertask.FieldOwner, field.TypeString)
	}
	if utu.mutation.AssigneeCleared() {
		_spec.ClearField(usertask.FieldAssignee, field.TypeString)
	}
	if utu.mutation.CandidateUsersCleared() {
		_spec.ClearField(usertask.FieldCandidateUsers, field.TypeString)
	}
	if utu.mutation.CandidateGroupsCleared() {
		_spec.ClearField(usertask.FieldCandidateGroups, field.TypeString)
	}
	if value, ok := utu.mutation.Operator(); ok {
		_spec.SetField(usertask.FieldOperator, field.TypeString, value)
	}
	if utu.mutation.OperatorCleared() {
		_spec.ClearField(usertask.FieldOperator, field.TypeString)
	}
	if value, ok := utu.mutation.Input(); ok {
		_spec.SetField(usertask.FieldInput, field.TypeJSON, value)
	}
	if utu.mutation.ResultCleared() {
		_spec.ClearField(usertask.FieldResult, field.TypeInt)
	}
	if value, ok := utu.mutation.EndTime(); ok {
		_spec.SetField(usertask.FieldEndTime, field.TypeTime, value)
	}
	if utu.mutation.EndTimeCleared() {
		_spec.ClearField(usertask.FieldEndTime, field.TypeTime)
	}
	if value, ok := utu.mutation.Status(); ok {
		_spec.SetField(usertask.FieldStatus, field.TypeInt, value)
	}
	if value, ok := utu.mutation.AddedStatus(); ok {
		_spec.AddField(usertask.FieldStatus, field.TypeInt, value)
	}
	if value, ok := utu.mutation.Version(); ok {
		_spec.SetField(usertask.FieldVersion, field.TypeInt, value)
	}
	if value, ok := utu.mutation.AddedVersion(); ok {
		_spec.AddField(usertask.FieldVersion, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, utu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	utu.mutation.done = true
	return n, nil
}

// UserTaskUpdateOne is the builder for updating a single UserTask entity.
type UserTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserTaskMutation
}

// SetOwner sets the "owner" field.
func (utuo *UserTaskUpdateOne) SetOwner(s string) *UserTaskUpdateOne {
	utuo.mutation.SetOwner(s)
	return utuo
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (utuo *UserTaskUpdateOne) SetNillableOwner(s *string) *UserTaskUpdateOne {
	if s != nil {
		utuo.SetOwner(*s)
	}
	return utuo
}

// ClearOwner clears the value of the "owner" field.
func (utuo *UserTaskUpdateOne) ClearOwner() *UserTaskUpdateOne {
	utuo.mutation.ClearOwner()
	return utuo
}

// SetOperator sets the "operator" field.
func (utuo *UserTaskUpdateOne) SetOperator(s string) *UserTaskUpdateOne {
	utuo.mutation.SetOperator(s)
	return utuo
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (utuo *UserTaskUpdateOne) SetNillableOperator(s *string) *UserTaskUpdateOne {
	if s != nil {
		utuo.SetOperator(*s)
	}
	return utuo
}

// ClearOperator clears the value of the "operator" field.
func (utuo *UserTaskUpdateOne) ClearOperator() *UserTaskUpdateOne {
	utuo.mutation.ClearOperator()
	return utuo
}

// SetInput sets the "input" field.
func (utuo *UserTaskUpdateOne) SetInput(l loong.Var) *UserTaskUpdateOne {
	utuo.mutation.SetInput(l)
	return utuo
}

// SetEndTime sets the "end_time" field.
func (utuo *UserTaskUpdateOne) SetEndTime(t time.Time) *UserTaskUpdateOne {
	utuo.mutation.SetEndTime(t)
	return utuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (utuo *UserTaskUpdateOne) SetNillableEndTime(t *time.Time) *UserTaskUpdateOne {
	if t != nil {
		utuo.SetEndTime(*t)
	}
	return utuo
}

// ClearEndTime clears the value of the "end_time" field.
func (utuo *UserTaskUpdateOne) ClearEndTime() *UserTaskUpdateOne {
	utuo.mutation.ClearEndTime()
	return utuo
}

// SetStatus sets the "status" field.
func (utuo *UserTaskUpdateOne) SetStatus(i int) *UserTaskUpdateOne {
	utuo.mutation.ResetStatus()
	utuo.mutation.SetStatus(i)
	return utuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (utuo *UserTaskUpdateOne) SetNillableStatus(i *int) *UserTaskUpdateOne {
	if i != nil {
		utuo.SetStatus(*i)
	}
	return utuo
}

// AddStatus adds i to the "status" field.
func (utuo *UserTaskUpdateOne) AddStatus(i int) *UserTaskUpdateOne {
	utuo.mutation.AddStatus(i)
	return utuo
}

// SetVersion sets the "version" field.
func (utuo *UserTaskUpdateOne) SetVersion(i int) *UserTaskUpdateOne {
	utuo.mutation.ResetVersion()
	utuo.mutation.SetVersion(i)
	return utuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (utuo *UserTaskUpdateOne) SetNillableVersion(i *int) *UserTaskUpdateOne {
	if i != nil {
		utuo.SetVersion(*i)
	}
	return utuo
}

// AddVersion adds i to the "version" field.
func (utuo *UserTaskUpdateOne) AddVersion(i int) *UserTaskUpdateOne {
	utuo.mutation.AddVersion(i)
	return utuo
}

// Mutation returns the UserTaskMutation object of the builder.
func (utuo *UserTaskUpdateOne) Mutation() *UserTaskMutation {
	return utuo.mutation
}

// Where appends a list predicates to the UserTaskUpdate builder.
func (utuo *UserTaskUpdateOne) Where(ps ...predicate.UserTask) *UserTaskUpdateOne {
	utuo.mutation.Where(ps...)
	return utuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (utuo *UserTaskUpdateOne) Select(field string, fields ...string) *UserTaskUpdateOne {
	utuo.fields = append([]string{field}, fields...)
	return utuo
}

// Save executes the query and returns the updated UserTask entity.
func (utuo *UserTaskUpdateOne) Save(ctx context.Context) (*UserTask, error) {
	return withHooks(ctx, utuo.sqlSave, utuo.mutation, utuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utuo *UserTaskUpdateOne) SaveX(ctx context.Context) *UserTask {
	node, err := utuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (utuo *UserTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := utuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utuo *UserTaskUpdateOne) ExecX(ctx context.Context) {
	if err := utuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utuo *UserTaskUpdateOne) check() error {
	if v, ok := utuo.mutation.Owner(); ok {
		if err := usertask.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "UserTask.owner": %w`, err)}
		}
	}
	if v, ok := utuo.mutation.Operator(); ok {
		if err := usertask.OperatorValidator(v); err != nil {
			return &ValidationError{Name: "operator", err: fmt.Errorf(`ent: validator failed for field "UserTask.operator": %w`, err)}
		}
	}
	return nil
}

func (utuo *UserTaskUpdateOne) sqlSave(ctx context.Context) (_node *UserTask, err error) {
	if err := utuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usertask.Table, usertask.Columns, sqlgraph.NewFieldSpec(usertask.FieldID, field.TypeUUID))
	id, ok := utuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := utuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertask.FieldID)
		for _, f := range fields {
			if !usertask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usertask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := utuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if utuo.mutation.ExecIDCleared() {
		_spec.ClearField(usertask.FieldExecID, field.TypeString)
	}
	if utuo.mutation.FormKeyCleared() {
		_spec.ClearField(usertask.FieldFormKey, field.TypeString)
	}
	if utuo.mutation.ActNameCleared() {
		_spec.ClearField(usertask.FieldActName, field.TypeString)
	}
	if value, ok := utuo.mutation.Owner(); ok {
		_spec.SetField(usertask.FieldOwner, field.TypeString, value)
	}
	if utuo.mutation.OwnerCleared() {
		_spec.ClearField(usertask.FieldOwner, field.TypeString)
	}
	if utuo.mutation.AssigneeCleared() {
		_spec.ClearField(usertask.FieldAssignee, field.TypeString)
	}
	if utuo.mutation.CandidateUsersCleared() {
		_spec.ClearField(usertask.FieldCandidateUsers, field.TypeString)
	}
	if utuo.mutation.CandidateGroupsCleared() {
		_spec.ClearField(usertask.FieldCandidateGroups, field.TypeString)
	}
	if value, ok := utuo.mutation.Operator(); ok {
		_spec.SetField(usertask.FieldOperator, field.TypeString, value)
	}
	if utuo.mutation.OperatorCleared() {
		_spec.ClearField(usertask.FieldOperator, field.TypeString)
	}
	if value, ok := utuo.mutation.Input(); ok {
		_spec.SetField(usertask.FieldInput, field.TypeJSON, value)
	}
	if utuo.mutation.ResultCleared() {
		_spec.ClearField(usertask.FieldResult, field.TypeInt)
	}
	if value, ok := utuo.mutation.EndTime(); ok {
		_spec.SetField(usertask.FieldEndTime, field.TypeTime, value)
	}
	if utuo.mutation.EndTimeCleared() {
		_spec.ClearField(usertask.FieldEndTime, field.TypeTime)
	}
	if value, ok := utuo.mutation.Status(); ok {
		_spec.SetField(usertask.FieldStatus, field.TypeInt, value)
	}
	if value, ok := utuo.mutation.AddedStatus(); ok {
		_spec.AddField(usertask.FieldStatus, field.TypeInt, value)
	}
	if value, ok := utuo.mutation.Version(); ok {
		_spec.SetField(usertask.FieldVersion, field.TypeInt, value)
	}
	if value, ok := utuo.mutation.AddedVersion(); ok {
		_spec.AddField(usertask.FieldVersion, field.TypeInt, value)
	}
	_node = &UserTask{config: utuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, utuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	utuo.mutation.done = true
	return _node, nil
}
