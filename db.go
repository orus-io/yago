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
		DefaultCallbacks,
	}
}

// DB is a session-looking thing.
// It provides a SQLA session like API, but has no
// instance cache, change tracking or unit-of-work
type DB struct {
	Metadata  *Metadata
	Engine    *qb.Engine
	Callbacks Callbacks
}

// Insert a struct in the database
func (db *DB) Insert(s MappedStruct) {
	db.Callbacks.BeforeInsert.Call(db, s)
	mapper := db.Metadata.GetMapper(s)
	insert := mapper.Table().Insert().Values(mapper.SQLValues(s))

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
	if mapper.AutoIncrementPKey() {
		if pkey, err := res.LastInsertId(); err != nil {
			panic(err)
		} else {
			mapper.LoadAutoIncrementPKeyValue(s, pkey)
		}
	}
	// TODO get the generated pkey, if any, and set it on the MappedStruct
	db.Callbacks.AfterInsert.Call(db, s)
}

// Update the struct attributes in DB
func (db *DB) Update(s MappedStruct, fields ...string) error {
	db.Callbacks.BeforeUpdate.Call(db, s)
	mapper := db.Metadata.GetMapper(s)
	update := mapper.Table().Update().
		Values(mapper.SQLValues(s, fields...)).
		Where(mapper.PKeyClause(s))

	res, err := db.Engine.Exec(update)
	if err != nil {
		return fmt.Errorf("yago Update: Exec failed with '%s'", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("yago Update: RowsAffected() failed with '%s'", err)
	}
	if ra != 1 {
		return fmt.Errorf("Update failed. More than 1 row where affected")
	}
	db.Callbacks.AfterUpdate.Call(db, s)
	return nil
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
	db.Callbacks.BeforeDelete.Call(db, s)
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
	db.Callbacks.AfterDelete.Call(db, s)
	return nil
}
