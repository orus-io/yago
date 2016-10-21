package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/slicebit/qb"

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

	engine, err := qb.New("sqlite3", ":memory:")
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

	q := db.Query(model.Person)
	if err := q.One(p); err == nil {
		panic("Should get a TooManyResultsError")
	}

	q = q.Where(model.Person.Name.Eq("Plouf"))
	if err := q.One(p); err == nil {
		panic("Should get a NoResultError")
	}

	q = db.Query(model.Person)
	q = q.Where(model.Person.Name.Eq("Titi"))

	p = &Person{}
	if err := q.One(p); err != nil {
		panic(err)
	}
	fmt.Println(p.Name, "created at", p.CreatedAt)

	var all []Person
	q = db.Query(model.Person)
	if err := q.All(&all); err != nil {
		panic(err)
	}
	fmt.Println("Loaded all persons:", all)

	var allP []*Person
	q = db.Query(model.Person)
	if err := q.All(&allP); err != nil {
		panic(err)
	}
	fmt.Println("Loaded all persons as pointers:", allP)

	var count uint64
	if err := db.Query(model.Person).Count(&count); err != nil || count != 2 {
		panic(fmt.Sprintf("Count returned (%d, %s)", count, err))
	}
	if exists, err := db.Query(model.Person).Exists(); err != nil || exists != true {
		panic(fmt.Sprintf("Exists returned (%d, %s)", exists, err))
	}

	p.Name = "Plouf"

	db.Update(p)

	q = db.Query(model.Person)
	q = q.Where(model.Person.Name.Eq("Plouf"))

	p = &Person{}
	if err := q.One(p); err != nil {
		panic(err)
	}
	fmt.Println(p.Name, "Updated at", p.UpdatedAt)

	n := PhoneNumber{PersonID: p.ID, Name: "mobile", Number: "06"}
	db.Insert(&n)

	q = db.Query(model.Person)
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
