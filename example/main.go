package main

import (
	"bitbucket.org/cdevienne/yagorm"

	"github.com/aacanakin/qb"
)

func main() {
	db, err := qb.New("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	session := yagorm.New(db)
	defer session.Close()
	meta := yagorm.NewMetadata(db.Metadata())

	meta.AddMapper(&PersonMapper{})

	if err := db.CreateAll(); err != nil {
		panic(err)
	}

	s := yagorm.Select(meta, &Person{})
	s.GroupBy()
}
