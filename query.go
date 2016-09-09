package yago

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/aacanakin/qb"
)

// Query helps querying structs from the database
type Query struct {
	db         *DB
	mapper     Mapper
	selectStmt qb.SelectStmt
}

// NewQuery creates a new query
func NewQuery(db *DB, mapper Mapper) Query {
	return Query{
		db:         db,
		mapper:     mapper,
		selectStmt: mapper.Table().Select(mapper.FieldList()...),
	}
}

// Select redefines the SELECT clauses
func (q Query) Select(clause ...qb.Clause) Query {
	q.selectStmt = q.selectStmt.Select(clause...)
	return q
}

// Where add filter clauses to the query
func (q Query) Where(clause qb.Clause) Query {
	return Query{
		db:         q.db,
		mapper:     q.mapper,
		selectStmt: q.selectStmt.Where(clause),
	}
}

// InnerJoin joins a table
func (q Query) InnerJoin(mp MapperProvider, clause ...qb.Clause) Query {
	q.selectStmt = q.selectStmt.InnerJoin(mp.GetMapper().Table(), clause...)
	return q
}

// LeftJoin joins a table
func (q Query) LeftJoin(mp MapperProvider, clause ...qb.Clause) Query {
	q.selectStmt = q.selectStmt.LeftJoin(mp.GetMapper().Table(), clause...)
	return q
}

// RightJoin joins a table
func (q Query) RightJoin(mp MapperProvider, clause ...qb.Clause) Query {
	q.selectStmt = q.selectStmt.RightJoin(mp.GetMapper().Table(), clause...)
	return q
}

// SQLQuery runs the query
func (q Query) SQLQuery() (*sql.Rows, error) {
	return q.db.Engine.Query(q.selectStmt)
}

// SQLQueryRow runs the query and expects at most one row in the result
func (q Query) SQLQueryRow() *sql.Row {
	return q.db.Engine.QueryRow(q.selectStmt)
}

// One returns one and only one struct from the query.
// If query has no result or more than one, an error is returned
func (q Query) One(s MappedStruct) error {
	rows, err := q.SQLQuery()
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("NoResultError")
	}

	err = q.mapper.Scan(rows, s)
	if err != nil {
		return err
	}
	if rows.Next() {
		return fmt.Errorf("TooManyResultsError")
	}
	return nil
}

// Get returns a record from its primary key values
func (q Query) Get(s MappedStruct, pkey ...interface{}) error {
	return q.Where(q.mapper.PKeyClause(pkey)).One(s)
}

// All load all the structs matching the query
func (q Query) All(value interface{}) error {
	rows, err := q.SQLQuery()
	if err != nil {
		return err
	}
	defer rows.Close()

	resultType := q.mapper.StructType()

	results := reflect.Indirect(reflect.ValueOf(value))

	var (
		isPtr     bool
		wrongType bool
	)

	if results.Kind() != reflect.Slice {
		wrongType = true
	} else {
		elemType := results.Type().Elem()
		if elemType.Kind() == reflect.Ptr {
			isPtr = true
			wrongType = results.Type().Elem().Elem() != resultType
		} else {
			wrongType = results.Type().Elem() != resultType
		}
	}
	if wrongType {
		return fmt.Errorf("yago Query.All(): Expected a []%s, got %v", resultType, results.Type())
	}

	// Empty the slice
	results.Set(reflect.MakeSlice(results.Type(), 0, 0))

	for rows.Next() {
		elem := reflect.New(resultType).Elem()
		if err := q.mapper.Scan(rows, elem.Addr().Interface().(MappedStruct)); err != nil {
			return fmt.Errorf("yago Query.All(): Error while scanning: %s", err)
		}
		if isPtr {
			results.Set(reflect.Append(results, elem.Addr()))
		} else {
			results.Set(reflect.Append(results, elem))
		}
	}
	if err != nil {
		return err
	}
	if rows.Next() {
		return fmt.Errorf("TooManyResultsError")
	}
	return nil
}

// Scalar execute the query and retrieve a single value from it
func (q Query) Scalar(value interface{}) error {
	rows, err := q.SQLQuery()
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return ErrRecordNotFound
	}
	if columns, err := rows.Columns(); err != nil || len(columns) != 1 {
		return ErrInvalidColumns
	}
	err = rows.Scan(value)
	if err != nil {
		return err
	}
	if rows.Next() {
		return ErrMultipleRecords
	}
	return nil
}

// Count change the columns to COUNT(*), execute the query and returns
// the result
func (q Query) Count(count interface{}) error {

	// XXX mapper should be able to return a list of pkey fields
	// XXX When qb supports COUNT(*), use it
	q.selectStmt = q.selectStmt.Select(qb.Count(
		q.mapper.Table().PrimaryCols()[0]),
	)
	return q.Select(
		qb.Count(qb.SQLText("*")),
	).Scalar(count)
}

// Exists return true if any record matches the current query
func (q Query) Exists() (exists bool, err error) {
	q.selectStmt = q.selectStmt.Select(qb.SQLText("1")).Limit(0, 1)
	err = q.Scalar(&exists)
	return
}
