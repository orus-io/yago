package main

import (
	"bitbucket.org/cdevienne/yagorm"

	"github.com/aacanakin/qb"
)

func main() {
	meta := yagorm.NewMetadata()
	meta.AddMapper(&PersonMapper{})

	s := yagorm.Select(meta, &Person{})
	s.GroupBy()

	engine, err := qb.NewEngine("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	engine.SetDialect(qb.NewDialect("sqlite3"))

	meta.GetQbMetadata().CreateAll(engine)
}
