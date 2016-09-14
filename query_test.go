package yago_test

import (
	"testing"

	"github.com/aacanakin/qb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orus-io/yago"

	"github.com/stretchr/testify/assert"
)

type querySQLTests struct {
	expect string
	q      yago.Query
}

func makeQuerySQLTests(db *yago.DB, model FixtureModel) []querySQLTests {
	return []querySQLTests{
		querySQLTests{
			"SELECT X\nFROM person_struct\nWHERE (Y AND Z)",
			db.Query(model.PersonStruct).Select(qb.SQLText("X")).
				Filter(qb.SQLText("Y")).
				Filter(qb.SQLText("Z")),
		},
		querySQLTests{
			"SELECT X\nFROM person_struct\nWHERE (Y AND Z)",
			db.Query(model.PersonStruct).Select(qb.SQLText("X")).
				Where(qb.SQLText("Y"), qb.SQLText("Z")),
		},
	}
}

func asSQL(query yago.Query) string {
	sql, _ := asSQLBind(query)
	return sql
}

func asSQLBind(query yago.Query) (string, []interface{}) {
	dialect := qb.NewDialect("default")
	s := query.SelectStmt()
	ctx := qb.NewCompilerContext(dialect)
	return s.Accept(ctx), ctx.Binds
}

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

func TestQuerySQL(t *testing.T) {
	db, model := initModel(t)

	for _, tt := range makeQuerySQLTests(db, model) {
		assert.Equal(t,
			tt.expect,
			asSQL(tt.q),
		)
	}
}
