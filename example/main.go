package main

import (
	"fmt"

	"github.com/aacanakin/qb"

	"github.com/orus-io/yago"
)

// Model gives easy access to various things
type Model struct {
	Meta *yago.Metadata

	Person      PersonModel
	PhoneNumber PhoneNumberModel
}

// NewModel initialize a model
func NewModel() *Model {
	meta := yago.NewMetadata()
	return &Model{
		Meta:        meta,
		Person:      NewPersonModel(meta),
		PhoneNumber: NewPhoneNumberModel(meta),
	}
}

func main() {
	model := NewModel()

	s := yago.Select(model.Meta, &Person{})
	s.GroupBy()

	engine, err := qb.NewEngine("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	engine.SetDialect(qb.NewDialect("sqlite3"))

	if err := model.Meta.GetQbMetadata().CreateAll(engine); err != nil {
		panic(err)
	}

	db := yago.New(model.Meta, engine)

	p := NewPerson()
	p.Name = "Toto"
	db.Insert(p)
	fmt.Println("Inserted", p.Name, "got ID", p.ID)

	p = NewPerson()
	p.Name = "Titi"
	db.Insert(p)
	fmt.Println("Inserted", p.Name, "got ID", p.ID)

	q := db.Query(&Person{})
	if err := q.One(p); err == nil {
		panic("Should get a TooManyResultsError")
	}

	q = q.Where(model.Person.Name.Eq("Plouf"))
	if err := q.One(p); err == nil {
		panic("Should get a NoResultError")
	}

	q = db.Query(&Person{})
	q = q.Where(model.Person.Name.Eq("Titi"))

	p = &Person{}
	if err := q.One(p); err != nil {
		panic(err)
	}
	fmt.Println(p.Name, "created at", p.CreatedAt)

	p.Name = "Plouf"

	db.Update(p)

	q = db.Query(&Person{})
	q = q.Where(model.Person.Name.Eq("Plouf"))

	p = &Person{}
	if err := q.One(p); err != nil {
		panic(err)
	}
	fmt.Println(p.Name, "Updated at", p.UpdatedAt)

	n := PhoneNumber{PersonID: p.ID, Name: "mobile", Number: "06"}
	db.Insert(&n)

	q = db.Query(&Person{})
	q = q.LeftJoin(model.PhoneNumber, model.Person.ID, model.PhoneNumber.PersonID)
	q = q.Where(model.PhoneNumber.Name.Eq("mobile"))

	if err := q.One(p); err != nil {
		panic(err)
	}

	db.Delete(p)
	if err := q.One(p); err == nil {
		panic("Should get a 'NoResultError'")
	}
}
