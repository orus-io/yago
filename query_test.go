package yago_test

import (
	"testing"

	"github.com/aacanakin/qb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orus-io/yago"

	"github.com/stretchr/testify/assert"
)

func initModel(t *testing.T) (db *yago.DB, model FixtureModel) {
	engine, err := qb.New("sqlite3", ":memory:")
	meta := yago.NewMetadata()
	model = NewFixtureModel(meta)
	db = yago.New(meta, engine)
	meta.GetQbMetadata().CreateAll(engine)
	assert.Nil(t, err)
	return
}

func TestGet(t *testing.T) {
	db, model := initModel(t)
	p := PersonStruct{FirstName: "John"}
	assert.Nil(t, db.Insert(&p))
	t.Log(p.ID)

	var p1 PersonStruct

	assert.Nil(t, db.Query(model.PersonStruct).Get(&p1, p.ID))
	assert.Equal(t, p.FirstName, p1.FirstName)
}
