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

// Where add filter clauses to the query
func (q Query) Where(clause qb.Clause) Query {
	return Query{
		db:         q.db,
		mapper:     q.mapper,
		selectStmt: q.selectStmt.Where(clause),
	}
}

// InnerJoinMapper joins a mapper table
func (q Query) InnerJoinMapper(mapper Mapper, fromCol qb.ColumnElem, col qb.ColumnElem) Query {
	q.selectStmt = q.selectStmt.InnerJoin(*mapper.Table(), fromCol, col)
	return q
}

// LeftJoinMapper joins a mapper table
func (q Query) LeftJoinMapper(mapper Mapper, fromCol qb.ColumnElem, col qb.ColumnElem) Query {
	q.selectStmt = q.selectStmt.LeftJoin(*mapper.Table(), fromCol, col)
	return q
}

// RightJoinMapper joins a mapper table
func (q Query) RightJoinMapper(mapper Mapper, fromCol qb.ColumnElem, col qb.ColumnElem) Query {
	q.selectStmt = q.selectStmt.RightJoin(*mapper.Table(), fromCol, col)
	return q
}

// InnerJoin joins a table
func (q Query) InnerJoin(model Model, fields ...ScalarField) Query {
	// TODO if fields is empty, find the relation based on foreign keys
	return q.InnerJoinMapper(model.GetMapper(), fields[0].Column, fields[1].Column)
}

// LeftJoin joins a table
func (q Query) LeftJoin(model Model, fields ...ScalarField) Query {
	// TODO if fields is empty, find the relation based on foreign keys
	return q.LeftJoinMapper(model.GetMapper(), fields[0].Column, fields[1].Column)
}

// RightJoin joins a table
func (q Query) RightJoin(model Model, fields ...ScalarField) Query {
	// TODO if fields is empty, find the relation based on foreign keys
	return q.RightJoinMapper(model.GetMapper(), fields[0].Column, fields[1].Column)
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

// Count change the columns to COUNT(*), execute the query and returns
// the result
func (q Query) Count() (uint64, error) {
	// XXX mapper should be able to return a list of pkey fields
	// XXX When qb supports COUNT(*), use it
	q.selectStmt = q.selectStmt.Select(qb.Count(
		q.mapper.Table().PrimaryCols()[0]),
	)
	row, err := q.SQLQuery()
	if err != nil {
		return 0, err
	}
	defer row.Close()
	if !row.Next() {
		panic("No result")
	}
	var count uint64
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Exists return true if any record matches the current query
func (q Query) Exists() (bool, error) {
	q.selectStmt = q.selectStmt.Select(qb.SQLText("1")).Limit(0, 1)
	row, err := q.SQLQuery()
	if err != nil {
		return false, err
	}
	defer row.Close()
	if !row.Next() {
		panic("No result")
	}
	var exists bool
	err = row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
