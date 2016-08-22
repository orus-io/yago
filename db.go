package yago

import (
	"fmt"

	"github.com/aacanakin/qb"
)

// New initialise a new DB
func New(metadata *Metadata, engine *qb.Engine) *DB {
	return &DB{
		metadata,
		engine,
	}
}

// DB is a session-looking thing.
// It provides a SQLA session like API, but has no
// instance cache, change tracking or unit-of-work
type DB struct {
	Metadata *Metadata
	Engine   *qb.Engine
}

// Add insert a struct in the database
func (db *DB) Add(s MappedStruct) {
	mapper := db.Metadata.GetMapper(s)
	insert := mapper.Table().Insert().Values(mapper.Values(s))

	res, err := db.Engine.Exec(insert)
	if err != nil {
		panic(err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if ra != 1 {
		panic("Insert failed")
	}
}

// Query returns a new Query for the struct
func (db *DB) Query(s MappedStruct) Query {
	mapper := db.Metadata.GetMapper(s)
	return db.QueryFromMapper(mapper)
}

// QueryFromMapper returns a new Query for the mapper
func (db *DB) QueryFromMapper(m Mapper) Query {
	return NewQuery(db, m)
}

// Delete a struct from the database
func (db *DB) Delete(s MappedStruct) error {
	mapper := db.Metadata.GetMapper(s)
	del := mapper.Table().Delete().Where(mapper.PKeyClause(s))
	res, err := db.Engine.Exec(del)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("Wrong number of rows affected. Expected 1, got %v", rowsAffected)
	}
	return err
}
