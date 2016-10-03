# Yet Another Go ORM

[![Go Report Card](https://goreportcard.com/badge/github.com/orus-io/yago)](https://goreportcard.com/report/github.com/orus-io/yago)

Yago is an attempt to create an ORM for Go.

## Features

- SQLAlchemy inspired
- based on the 'qb' database toolkit (https://github.com/aacanakin/qb),
- based on non-empty interface and code generation as 'reform' does
  (https://github.com/go-reform/reform/)

## Goals/Key concepts

First goals are:

- Do not hide the SQL, only wrap it with a thin layer allowing multiple
  dialects. This layer is 'qb', to which we contribute back.
- Provide a low-cost abstraction on the schema, removing the need to
  know any actual column or table name when writting queries.
- Have a 'record-level' API, which is very efficient at CRUD operations
  with plain structs, and is developer friendly.
- Be backend agnotistic (as much as possible).

Long term goal is to also provide higher level ORM concepts like
instances, session, unit-of-work... It will be based on the record-level
API, which needs to be ironned first.

## Try it out

```bash

go get -u github.com/orus-io/yago
go install github.com/orus-io/cmd/yago
```

Define your model (for example in model.go):

```go
package main

import (
	"github.com/orus-io/yago"
)

//go:generate yago --fmt

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

//yago:autoattrs
type Person struct {
	Base
	Name  string  `yago:"index"`
	Email *string `yago:"email_address,unique_index"`
}

// PhoneNumber is a phone number
//yago:autoattrs
type PhoneNumber struct {
	Base
	PersonID int64 `yago:"fk=Person ondelete cascade onupdate cascade"`
	Name     string
	Number   string
}
```

Let yago generate the code:

```bash
go generate
```

If you named the file "model.go", this command generated a new file "model_yago.go".

Now we can use the model to read/write things in the database:

```go
package main

import (
	"fmt"

	"github.com/orus-io/qb"
	_ "github.com/mattn/go-sqlite3"

	"github.com/orus-io/yago"
)

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

	q = db.Query(model.Person)
	q = q.Where(model.Person.Name.Eq("Toto"))

	p = &Person{}
	if err := q.One(p); err != nil {
		panic(err)
	}

	var all []Person
	q = db.Query(model.Person)
	if err := q.All(&all); err != nil {
		panic(err)
	}
	fmt.Println("Loaded all persons:", all)
}
```

## Current state:

- Very experimental, APIs may brake without warning.
- Current focus is on finding the right API.

Using this code in production is not recommended at this stage (although
we are about to do it).

Feel free to come and discuss the design, propose patches...
