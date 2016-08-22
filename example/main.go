package main

import (
	"fmt"

	"github.com/aacanakin/qb"

	"bitbucket.org/cdevienne/yago"
)

// Model gives easy access to various things
type Model struct {
	Meta *yago.Metadata

	Person *PersonMapper
}

// NewModel initialize a model
func NewModel() *Model {
	model := &Model{
		Meta: yago.NewMetadata(),
	}
	model.Person = NewPersonMapper()
	model.Meta.AddMapper(model.Person)
	return model
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

	p = NewPerson()
	p.Name = "Titi"
	db.Insert(p)

	q := db.Query(&Person{})
	if err := q.One(p); err == nil {
		panic("Should get a TooManyResultsError")
	}

	q = q.Where(model.Person.Fields.Name.Eq("Plouf"))
	if err := q.One(p); err == nil {
		panic("Should get a NoResultError")
	}

	q = db.Query(&Person{})
	q = q.Where(model.Person.Fields.Name.Eq("Titi"))

	p = &Person{}
	if err := q.One(p); err != nil {
		panic(err)
	}
	fmt.Println(p.Name)

	db.Delete(p)
	if err := q.One(p); err == nil {
		panic("Should get a 'NoResultError'")
	}
}
