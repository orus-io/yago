package yago_test

import (
	"fmt"
	"time"

	"github.com/m4rw3r/uuid"

	"github.com/orus-io/yago"
)

//go:generate yago --fmt --package yago_test --output fixtures_yago_test.go

type FixtureModel struct {
	meta *yago.Metadata

	PersonStruct PersonStructModel
	SimpleStruct SimpleStructModel
}

func NewFixtureModel(meta *yago.Metadata) FixtureModel {
	return FixtureModel{
		meta:         meta,
		PersonStruct: NewPersonStructModel(meta),
		SimpleStruct: NewSimpleStructModel(meta),
	}
}

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
	FirstName string `yago:"unique"`
	LastName  string `yago:"null"`
}

//yago:notable
type AutoIncBase struct {
	ID int64 `yago:"primary_key,auto_increment"`
}

//yago:autoattrs
type AutoIncChild struct {
	AutoIncBase

	Name   string
	Person uuid.UUID `yago:"fk=PersonStruct ONDELETE SET NULL ONUPDATE CASCADE"`
}

func (s *BaseStruct) BeforeInsert(db *yago.DB) {
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
