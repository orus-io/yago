package yagorm

import (
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

// Save insert a struct in the database
func (db *DB) Save(s MappedStruct) {
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
