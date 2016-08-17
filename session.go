package yagorm

import (
	"github.com/aacanakin/qb"
)

// New initialise a new PseudoSession
func New(metadata *Metadata, engine *qb.Engine) *PseudoSession {
	return &PseudoSession{
		metadata,
		engine,
	}
}

// PseudoSession is a session-looking thing.
// It provides a SQLA session like API, but has no
// instance cache, change tracking or unit-of-work
type PseudoSession struct {
	Metadata *Metadata
	Engine   *qb.Engine
}

// Save insert a struct in the database
func (ps *PseudoSession) Save(s MappedStruct) {
	insert := ps.Metadata.GetMapper(s).Table().Insert().Values(s.Values())

	res, err := ps.Engine.Exec(insert)
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
