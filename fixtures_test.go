package yago_test

import (
	"fmt"

	"github.com/m4rw3r/uuid"

	"github.com/orus-io/yago"
)

//go:generate yago --package yago_test --output fixtures_yago_test.go

//yago:
type SimpleStruct struct {
	ID   int64  `yago:"primary_key,auto_increment"`
	Name string `yago:"unique_index"`
}

//yago:autoattrs
type PersonStruct struct {
	ID        uuid.UUID `yago:"primary_key"`
	FirstName string
	LastName  string
}

func (p *PersonStruct) BeforeCreate(db *yago.DB) {
	var err error
	p.ID, err = uuid.V4()
	if err != nil {
		panic(fmt.Sprintf("Cannot generate a UUID. Got err %s", err))
	}
}
