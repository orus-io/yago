package main

import (
	"fmt"

	"github.com/aacanakin/qb"

	"bitbucket.org/cdevienne/yago"
)

func main() {
	meta := yago.NewMetadata()
	meta.AddMapper(&PersonMapper{})

	s := yago.Select(meta, &Person{})
	s.GroupBy()

	engine, err := qb.NewEngine("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	engine.SetDialect(qb.NewDialect("sqlite3"))

	if err := meta.GetQbMetadata().CreateAll(engine); err != nil {
		panic(err)
	}

	db := yago.New(meta, engine)

	p := NewPerson()
	p.Name = "Toto"
	db.Add(p)

	p = NewPerson()
	p.Name = "Titi"
	db.Add(p)

	q := db.Query(&Person{})
	if err := q.One(p); err == nil {
		panic("Should get a TooManyResultsError")
	}

	q = q.Where(db.Metadata.GetMapper(&Person{}).Table().C("name").Eq("Plouf"))
	if err := q.One(p); err == nil {
		panic("Should get a NoResultError")
	}

	q = db.Query(&Person{})
	q = q.Where(db.Metadata.GetMapper(&Person{}).Table().C("name").Eq("Titi"))

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
