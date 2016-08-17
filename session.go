package yagorm

import (
	"github.com/aacanakin/qb"
)

// New initialise a new PseudoSession
func New(session *qb.Session) *PseudoSession {
	return &PseudoSession{
		session,
		NewMetadata(),
	}
}

// PseudoSession is a session-looking thing.
// It provides a SQLA session like API, but has no
// instance cache, change tracking or unit-of-work
type PseudoSession struct {
	qbSession *qb.Session
	Metadata  *Metadata
}

// QbSession returns the underlying qb.session
func (ps *PseudoSession) QbSession() *qb.Session {
	return ps.qbSession
}

// Close things
func (ps *PseudoSession) Close() {
	ps.qbSession.Close()
}
