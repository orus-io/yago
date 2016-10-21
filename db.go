package yago

import (
	"database/sql"
	"fmt"

	"github.com/slicebit/qb"
)

// IDB is the common interface of DB and Tx
type IDB interface {
	Insert(MappedStruct) error
	Update(MappedStruct, ...string) error
	Delete(MappedStruct) error
	Query(MapperProvider) Query

	GetEngine() Engine
}

// Engine is the common interface of qb.Engine and qb.Tx
type Engine interface {
	Exec(builder qb.Builder) (sql.Result, error)
	Query(builder qb.Builder) (*sql.Rows, error)
	QueryRow(builder qb.Builder) *sql.Row
}

// New initialise a new DB
func New(metadata *Metadata, engine *qb.Engine) *DB {
	return &DB{
		metadata,
		engine,
		DefaultCallbacks,
	}
}

// DB is a database handle with callbacks that links a Metadata and a qb.Engine
type DB struct {
	Metadata  *Metadata
	Engine    *qb.Engine
	Callbacks Callbacks
}

// GetEngine returns the underlying engine
func (db DB) GetEngine() Engine {
	return db.Engine
}

// Insert a struct in the database
func (db *DB) Insert(s MappedStruct) error {
	return db.doInsert(db.Engine, s)
}

// Update the struct attributes in DB
func (db *DB) Update(s MappedStruct, fields ...string) error {
	return db.doUpdate(db.Engine, s, fields...)
}

// Delete a struct in the database
func (db *DB) Delete(s MappedStruct) error {
	return db.doDelete(db.Engine, s)
}

func (db *DB) doInsertWithReturning(engine Engine, s MappedStruct) error {
	mapper := db.Metadata.GetMapper(s)

	insert := mapper.Table().Insert().
		Values(mapper.SQLValues(s)).
		Returning(mapper.Table().PrimaryCols()...)

	rows, err := engine.Query(insert)
	if err != nil {
		return fmt.Errorf("yago Insert: QueryRow failed with '%s'", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("yago Insert: No row returned by insert")
	}

	if err := mapper.ScanPKey(rows, s); err != nil {
		return fmt.Errorf("yago Insert: Error scanning the returned pkey: %s", err)
	}

	if rows.Next() {
		return ErrMultipleRecords
	}

	return nil
}

func (db *DB) doInsert(engine Engine, s MappedStruct) error {
	db.Callbacks.BeforeInsert.Call(db, s)
	mapper := db.Metadata.GetMapper(s)

	if mapper.AutoIncrementPKey() && db.Engine.Dialect().Driver() == "postgres" {
		return db.doInsertWithReturning(engine, s)
	}

	insert := mapper.Table().Insert().Values(mapper.SQLValues(s))

	res, err := engine.Exec(insert)
	if err != nil {
		return fmt.Errorf("yago Insert: Exec failed with '%s'", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("yago Insert: RowsAffected() failed with '%s'", err)
	}
	if ra != 1 {
		return fmt.Errorf("Update Insert. More than 1 row where affected")
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
	return nil
}

func (db *DB) doUpdate(engine Engine, s MappedStruct, fields ...string) error {
	db.Callbacks.BeforeUpdate.Call(db, s)
	mapper := db.Metadata.GetMapper(s)
	update := mapper.Table().Update().
		Values(mapper.SQLValues(s, fields...)).
		Where(mapper.PKeyClause(mapper.PKey(s)))

	res, err := engine.Exec(update)
	if err != nil {
		return fmt.Errorf("yago Update: Exec failed with '%s'", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("yago Update: RowsAffected() failed with '%s'", err)
	}
	if ra == 0 {
		return ErrRecordNotFound
	} else if ra > 1 {
		return ErrMultipleRecords
	}
	db.Callbacks.AfterUpdate.Call(db, s)
	return nil
}

// Query returns a new Query for the struct
func (db *DB) Query(mp MapperProvider) Query {
	mapper := mp.GetMapper()
	return NewQuery(db, mapper)
}

// Delete a struct from the database
func (db *DB) doDelete(engine Engine, s MappedStruct) error {
	db.Callbacks.BeforeDelete.Call(db, s)
	mapper := db.Metadata.GetMapper(s)
	del := mapper.Table().Delete().Where(mapper.PKeyClause(mapper.PKey(s)))
	res, err := engine.Exec(del)
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

// Begin start a new transaction
func (db *DB) Begin() (*Tx, error) {
	tx, err := db.Engine.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{db: db, tx: tx}, nil
}

// Tx is an on-going database transaction
type Tx struct {
	db *DB
	tx *qb.Tx
}

// GetEngine returns the underlying qb.Tx
func (tx Tx) GetEngine() Engine {
	return tx.tx
}

// Insert a new struct to the database
func (tx Tx) Insert(s MappedStruct) error {
	return tx.db.doInsert(tx.tx, s)
}

// Update write struct values to the database
// If fields is provided, only theses fields are written
func (tx Tx) Update(s MappedStruct, fields ...string) error {
	return tx.db.doUpdate(tx.tx, s, fields...)
}

// Delete drop a struct from the database
func (tx Tx) Delete(s MappedStruct) error {
	return tx.db.doDelete(tx.tx, s)
}

// Query returns a new Query
func (tx Tx) Query(mp MapperProvider) Query {
	return NewQuery(tx, mp.GetMapper())
}

// Commit commits the transaction
func (tx Tx) Commit() error {
	return tx.tx.Commit()
}

// Rollback aborts the transaction
func (tx Tx) Rollback() error {
	return tx.tx.Rollback()
}
