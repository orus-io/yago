package yago

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/orus-io/qb"
)

// IDB is the common interface of DB and Tx
type IDB interface {
	Insert(MappedStruct) error
	Update(MappedStruct, ...string) error
	Delete(MappedStruct) error
	Query(MapperProvider) Query

	InsertContext(context.Context, MappedStruct) error
	UpdateContext(context.Context, MappedStruct, ...string) error
	DeleteContext(context.Context, MappedStruct) error

	GetEngine() Engine
}

// Engine is the common interface of qb.Engine and qb.Tx
type Engine interface {
	Exec(builder qb.Builder) (sql.Result, error)
	Query(builder qb.Builder) (*sql.Rows, error)
	QueryRow(builder qb.Builder) qb.Row

	ExecContext(ctx context.Context, builder qb.Builder) (sql.Result, error)
	QueryContext(ctx context.Context, builder qb.Builder) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, builder qb.Builder) qb.Row
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
	return db.doInsert(context.Background(), db.Engine, s)
}

// InsertContext a struct in the database
func (db *DB) InsertContext(ctx context.Context, s MappedStruct) error {
	return db.doInsert(ctx, db.Engine, s)
}

// Update the struct attributes in DB
func (db *DB) Update(s MappedStruct, fields ...string) error {
	return db.doUpdate(context.Background(), db.Engine, s, fields...)
}

// UpdateContext the struct attributes in DB
func (db *DB) UpdateContext(ctx context.Context, s MappedStruct, fields ...string) error {
	return db.doUpdate(ctx, db.Engine, s, fields...)
}

// Delete a struct in the database
func (db *DB) Delete(s MappedStruct) error {
	return db.doDelete(context.Background(), db.Engine, s)
}

// Delete a struct in the database
func (db *DB) DeleteContext(ctx context.Context, s MappedStruct) error {
	return db.doDelete(ctx, db.Engine, s)
}

func (db *DB) doInsertWithReturning(ctx context.Context, engine Engine, s MappedStruct) error {
	mapper := db.Metadata.GetMapper(s)

	insert := mapper.Table().Insert().
		Values(mapper.SQLValues(s)).
		Returning(mapper.Table().PrimaryCols()...)

	rows, err := engine.QueryContext(ctx, insert)
	if err != nil {
		return err
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

// Close closes the underlying db connection
func (db *DB) Close() error {
	return db.Engine.Close()
}

func (db *DB) doInsert(ctx context.Context, engine Engine, s MappedStruct) error {
	db.Callbacks.BeforeInsert.Call(db, s)
	mapper := db.Metadata.GetMapper(s)

	if mapper.AutoIncrementPKey() && db.Engine.Dialect().Driver() == "postgres" {
		return db.doInsertWithReturning(ctx, engine, s)
	}

	insert := mapper.Table().Insert().Values(mapper.SQLValues(s))

	res, err := engine.ExecContext(ctx, insert)
	if err != nil {
		return err
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

func (db *DB) doUpdate(ctx context.Context, engine Engine, s MappedStruct, fields ...string) error {
	db.Callbacks.BeforeUpdate.Call(db, s)
	mapper := db.Metadata.GetMapper(s)
	update := mapper.Table().Update().
		Values(mapper.SQLValues(s, fields...)).
		Where(mapper.PKeyClause(mapper.PKey(s)))

	res, err := engine.ExecContext(ctx, update)
	if err != nil {
		return err
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
func (db *DB) doDelete(ctx context.Context, engine Engine, s MappedStruct) error {
	db.Callbacks.BeforeDelete.Call(db, s)
	mapper := db.Metadata.GetMapper(s)
	del := mapper.Table().Delete().Where(mapper.PKeyClause(mapper.PKey(s)))
	res, err := engine.ExecContext(ctx, del)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	} else if rowsAffected != 1 {
		return fmt.Errorf("Wrong number of rows affected. Expected 1, got %v", rowsAffected)
	}
	db.Callbacks.AfterDelete.Call(db, s)
	return nil
}

// Begin start a new transaction
func (db *DB) Begin() (*Tx, error) {
	return db.BeginContext(context.Background())
}

// BeginContext start a new transaction
func (db *DB) BeginContext(ctx context.Context) (*Tx, error) {
	tx, err := db.Engine.BeginContext(ctx)
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
	return tx.db.doInsert(context.Background(), tx.tx, s)
}

// InsertContext a new struct to the database
func (tx Tx) InsertContext(ctx context.Context, s MappedStruct) error {
	return tx.db.doInsert(ctx, tx.tx, s)
}

// Update write struct values to the database
// If fields is provided, only theses fields are written
func (tx Tx) Update(s MappedStruct, fields ...string) error {
	return tx.db.doUpdate(context.Background(), tx.tx, s, fields...)
}

// UpdateContext write struct values to the database
// If fields is provided, only theses fields are written
func (tx Tx) UpdateContext(ctx context.Context, s MappedStruct, fields ...string) error {
	return tx.db.doUpdate(ctx, tx.tx, s, fields...)
}

// Delete drop a struct from the database
func (tx Tx) Delete(s MappedStruct) error {
	return tx.db.doDelete(context.Background(), tx.tx, s)
}

// DeleteContext drop a struct from the database
func (tx Tx) DeleteContext(ctx context.Context, s MappedStruct) error {
	return tx.db.doDelete(ctx, tx.tx, s)
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
