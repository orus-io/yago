package main

import (
	"time"

	"github.com/orus-io/yago"
)

//go:generate yago --fmt

// Base struct for all model structs
//yago:notable,autoattrs
type Base struct {
	ID        int64 `yago:"primary_key,auto_increment"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// BeforeInsert callback
func (b *Base) BeforeInsert(db *yago.DB) {
	b.CreatedAt = time.Now()
}

// BeforeUpdate callback
func (b *Base) BeforeUpdate(db *yago.DB) {
	now := time.Now()
	b.UpdatedAt = &now
}
