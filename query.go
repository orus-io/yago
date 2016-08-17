package yagorm

import (
	"database/sql"
	"fmt"

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

// SQLQuery runs the query
func (q Query) SQLQuery() (*sql.Rows, error) {
	return q.db.Engine.Query(q.selectStmt)
}

// SQLQueryRow runs the query and expects at most one row in the result
func (q Query) SQLQueryRow() (*sql.Rows, error) {
	return q.db.Engine.Query(q.selectStmt)
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
