// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dhruv-vavliya/BookStore/ent/author"
	"github.com/dhruv-vavliya/BookStore/ent/book"
	"github.com/dhruv-vavliya/BookStore/ent/predicate"
)

// BookQuery is the builder for querying Book entities.
type BookQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Book
	withAuthor *AuthorQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookQuery builder.
func (bq *BookQuery) Where(ps ...predicate.Book) *BookQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BookQuery) Limit(limit int) *BookQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BookQuery) Offset(offset int) *BookQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BookQuery) Unique(unique bool) *BookQuery {
	bq.unique = &unique
	return bq
}

// Order adds an order step to the query.
func (bq *BookQuery) Order(o ...OrderFunc) *BookQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryAuthor chains the current query on the "author" edge.
func (bq *BookQuery) QueryAuthor() *AuthorQuery {
	query := &AuthorQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, selector),
			sqlgraph.To(author.Table, author.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.AuthorTable, book.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Book entity from the query.
// Returns a *NotFoundError when no Book was found.
func (bq *BookQuery) First(ctx context.Context) (*Book, error) {
	nodes, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{book.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BookQuery) FirstX(ctx context.Context) *Book {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Book ID from the query.
// Returns a *NotFoundError when no Book ID was found.
func (bq *BookQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{book.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BookQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Book entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Book entity is found.
// Returns a *NotFoundError when no Book entities are found.
func (bq *BookQuery) Only(ctx context.Context) (*Book, error) {
	nodes, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{book.Label}
	default:
		return nil, &NotSingularError{book.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BookQuery) OnlyX(ctx context.Context) *Book {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Book ID in the query.
// Returns a *NotSingularError when more than one Book ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BookQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{book.Label}
	default:
		err = &NotSingularError{book.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BookQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Books.
func (bq *BookQuery) All(ctx context.Context) ([]*Book, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BookQuery) AllX(ctx context.Context) []*Book {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Book IDs.
func (bq *BookQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(book.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BookQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BookQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BookQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BookQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BookQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BookQuery) Clone() *BookQuery {
	if bq == nil {
		return nil
	}
	return &BookQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]OrderFunc{}, bq.order...),
		predicates: append([]predicate.Book{}, bq.predicates...),
		withAuthor: bq.withAuthor.Clone(),
		// clone intermediate query.
		sql:    bq.sql.Clone(),
		path:   bq.path,
		unique: bq.unique,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookQuery) WithAuthor(opts ...func(*AuthorQuery)) *BookQuery {
	query := &AuthorQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withAuthor = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Book.Query().
//		GroupBy(book.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BookQuery) GroupBy(field string, fields ...string) *BookGroupBy {
	grbuild := &BookGroupBy{config: bq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(ctx), nil
	}
	grbuild.label = book.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Book.Query().
//		Select(book.FieldName).
//		Scan(ctx, &v)
func (bq *BookQuery) Select(fields ...string) *BookSelect {
	bq.fields = append(bq.fields, fields...)
	selbuild := &BookSelect{BookQuery: bq}
	selbuild.label = book.Label
	selbuild.flds, selbuild.scan = &bq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a BookSelect configured with the given aggregations.
func (bq *BookQuery) Aggregate(fns ...AggregateFunc) *BookSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BookQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bq.fields {
		if !book.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BookQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Book, error) {
	var (
		nodes       = []*Book{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [1]bool{
			bq.withAuthor != nil,
		}
	)
	if bq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, book.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Book).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Book{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withAuthor; query != nil {
		if err := bq.loadAuthor(ctx, query, nodes, nil,
			func(n *Book, e *Author) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BookQuery) loadAuthor(ctx context.Context, query *AuthorQuery, nodes []*Book, init func(*Book), assign func(*Book, *Author)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Book)
	for i := range nodes {
		if nodes[i].author_books == nil {
			continue
		}
		fk := *nodes[i].author_books
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(author.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "author_books" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bq *BookQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.fields
	if len(bq.fields) > 0 {
		_spec.Unique = bq.unique != nil && *bq.unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BookQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (bq *BookQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for i := range fields {
			if fields[i] != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BookQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(book.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = book.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.unique != nil && *bq.unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookGroupBy is the group-by builder for Book entities.
type BookGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BookGroupBy) Aggregate(fns ...AggregateFunc) *BookGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bgb *BookGroupBy) Scan(ctx context.Context, v any) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

func (bgb *BookGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range bgb.fields {
		if !book.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BookGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql.Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
		for _, f := range bgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bgb.fields...)...)
}

// BookSelect is the builder for selecting fields of Book entities.
type BookSelect struct {
	*BookQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BookSelect) Aggregate(fns ...AggregateFunc) *BookSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BookSelect) Scan(ctx context.Context, v any) error {
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	bs.sql = bs.BookQuery.sqlQuery(ctx)
	return bs.sqlScan(ctx, v)
}

func (bs *BookSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(bs.sql))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		bs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		bs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := bs.sql.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
