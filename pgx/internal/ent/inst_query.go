// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/it512/loong/pgx/internal/ent/inst"
	"github.com/it512/loong/pgx/internal/ent/predicate"
)

// InstQuery is the builder for querying Inst entities.
type InstQuery struct {
	config
	ctx        *QueryContext
	order      []inst.OrderOption
	inters     []Interceptor
	predicates []predicate.Inst
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InstQuery builder.
func (iq *InstQuery) Where(ps ...predicate.Inst) *InstQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InstQuery) Limit(limit int) *InstQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InstQuery) Offset(offset int) *InstQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InstQuery) Unique(unique bool) *InstQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InstQuery) Order(o ...inst.OrderOption) *InstQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// First returns the first Inst entity from the query.
// Returns a *NotFoundError when no Inst was found.
func (iq *InstQuery) First(ctx context.Context) (*Inst, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{inst.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InstQuery) FirstX(ctx context.Context) *Inst {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Inst ID from the query.
// Returns a *NotFoundError when no Inst ID was found.
func (iq *InstQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{inst.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InstQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Inst entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Inst entity is found.
// Returns a *NotFoundError when no Inst entities are found.
func (iq *InstQuery) Only(ctx context.Context) (*Inst, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{inst.Label}
	default:
		return nil, &NotSingularError{inst.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InstQuery) OnlyX(ctx context.Context) *Inst {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Inst ID in the query.
// Returns a *NotSingularError when more than one Inst ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InstQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{inst.Label}
	default:
		err = &NotSingularError{inst.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InstQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Insts.
func (iq *InstQuery) All(ctx context.Context) ([]*Inst, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Inst, *InstQuery]()
	return withInterceptors[[]*Inst](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InstQuery) AllX(ctx context.Context) []*Inst {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Inst IDs.
func (iq *InstQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(inst.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InstQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InstQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InstQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InstQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InstQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InstQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InstQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InstQuery) Clone() *InstQuery {
	if iq == nil {
		return nil
	}
	return &InstQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]inst.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Inst{}, iq.predicates...),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InstID string `json:"inst_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Inst.Query().
//		GroupBy(inst.FieldInstID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InstQuery) GroupBy(field string, fields ...string) *InstGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InstGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = inst.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		InstID string `json:"inst_id,omitempty"`
//	}
//
//	client.Inst.Query().
//		Select(inst.FieldInstID).
//		Scan(ctx, &v)
func (iq *InstQuery) Select(fields ...string) *InstSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InstSelect{InstQuery: iq}
	sbuild.label = inst.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InstSelect configured with the given aggregations.
func (iq *InstQuery) Aggregate(fns ...AggregateFunc) *InstSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InstQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !inst.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InstQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Inst, error) {
	var (
		nodes = []*Inst{}
		_spec = iq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Inst).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Inst{config: iq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iq *InstQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InstQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(inst.Table, inst.Columns, sqlgraph.NewFieldSpec(inst.FieldID, field.TypeUUID))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inst.FieldID)
		for i := range fields {
			if fields[i] != inst.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InstQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(inst.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = inst.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstGroupBy is the group-by builder for Inst entities.
type InstGroupBy struct {
	selector
	build *InstQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InstGroupBy) Aggregate(fns ...AggregateFunc) *InstGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InstGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstQuery, *InstGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InstGroupBy) sqlScan(ctx context.Context, root *InstQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InstSelect is the builder for selecting fields of Inst entities.
type InstSelect struct {
	*InstQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InstSelect) Aggregate(fns ...AggregateFunc) *InstSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InstSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstQuery, *InstSelect](ctx, is.InstQuery, is, is.inters, v)
}

func (is *InstSelect) sqlScan(ctx context.Context, root *InstQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
