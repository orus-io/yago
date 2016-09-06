package yago_test

import (
	"fmt"
	"time"

	"github.com/m4rw3r/uuid"

	"github.com/orus-io/yago"
)

//go:generate yago --package yago_test --output fixtures_yago_test.go

//yago:
type SimpleStruct struct {
	ID   int64  `yago:"primary_key,auto_increment"`
	Name string `yago:"unique_index"`
}

//yago:notable,autoattrs
type BaseStruct struct {
	ID        uuid.UUID `yago:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//yago:autoattrs
type PersonStruct struct {
	BaseStruct

	Active    bool
	FirstName string
	LastName  string
}

func (s *BaseStruct) BeforeCreate(db *yago.DB) {
	var err error
	s.ID, err = uuid.V4()
	if err != nil {
		panic(fmt.Sprintf("Cannot generate a UUID. Got err %s", err))
	}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *BaseStruct) BeforeUpdate(db *yago.DB) {
	s.UpdatedAt = time.Now()
}
