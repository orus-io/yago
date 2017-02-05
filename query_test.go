package yago_test

import (
	"testing"

	"github.com/orus-io/yago"
	"github.com/slicebit/qb"

	"github.com/stretchr/testify/assert"
)

type querySQLTests struct {
	expect string
	q      yago.Query
}

func makeQuerySQLTests(db *yago.DB, model FixtureModel) []querySQLTests {
	return []querySQLTests{
		{
			"SELECT X\nFROM person_struct\nWHERE (Y AND Z)",
			db.Query(model.PersonStruct).Select(qb.SQLText("X")).
				Filter(qb.SQLText("Y")).
				Filter(qb.SQLText("Z")),
		},
		{
			"SELECT X\nFROM person_struct\nWHERE (Y AND Z)",
			db.Query(model.PersonStruct).Select(qb.SQLText("X")).
				Where(qb.SQLText("Y"), qb.SQLText("Z")),
		},
		{
			"SELECT X\nFROM person_struct\nORDER BY last_name ASC",
			db.Query(model.PersonStruct).Select(qb.SQLText("X")).
				OrderBy(model.PersonStruct.LastName),
		},
		{
			"SELECT X\nFROM person_struct\nORDER BY last_name ASC\nLIMIT 5 OFFSET 0",
			db.Query(model.PersonStruct).Select(qb.SQLText("X")).
				OrderBy(model.PersonStruct.LastName).Limit(5, 0),
		},
	}
}

func TestGet(t *testing.T) {
	db, model, cleanup := initModel(t)
	defer cleanup()
	p := PersonStruct{FirstName: "John"}
	assert.Nil(t, db.Insert(&p))
	t.Log(p.ID)

	var p1 PersonStruct

	assert.Nil(t, db.Query(model.PersonStruct).Get(&p1, p.ID))
	assert.Equal(t, p.FirstName, p1.FirstName)
}

func TestQuerySQL(t *testing.T) {
	db, model, cleanup := initModel(t)
	defer cleanup()

	for _, tt := range makeQuerySQLTests(db, model) {
		assert.Equal(t,
			tt.expect,
			asSQL(tt.q),
		)
	}
}
