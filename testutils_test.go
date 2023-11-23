package yago_test

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orus-io/yago"
	"github.com/orus-io/qb"
	_ "github.com/orus-io/qb/dialects/postgres"
	_ "github.com/orus-io/qb/dialects/sqlite"

	"github.com/stretchr/testify/assert"
)

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

func initModel(t *testing.T) (db *yago.DB, model FixtureModel, cleanup func()) {
	return initModelWithDriver(t, "sqlite3")
}

func initModelWithDriver(t *testing.T, driver string) (db *yago.DB, model FixtureModel, cleanup func()) {
	var dsn string
	switch driver {
	case "postgres":
		dsn = os.Getenv("YAGO_TEST_POSTGRES")
	case "sqlite3":
		dsn = os.Getenv("YAGO_TEST_SQLITE")
		if dsn == "" {
			dsn = ":memory:"
		}
	default:
		panic("Unknown driver for tests: " + driver)
	}
	engine, err := qb.New(driver, dsn)
	assert.Nil(t, err)

	meta := yago.NewMetadata()
	model = NewFixtureModel(meta)
	db = yago.New(meta, engine)
	CleanupDB(t, db, false)

	assert.Nil(t, meta.GetQbMetadata().CreateAll(engine))
	cleanup = func() { CleanupFunc(t, db, true) }
	return
}

func CleanupFunc(t *testing.T, db *yago.DB, reportErrors bool) {
	CleanupDB(t, db, reportErrors)
	db.Close()
}

func CleanupDB(t *testing.T, db *yago.DB, reportErrors bool) {
	for _, table := range db.Metadata.GetQbMetadata().Tables() {
		_, err := db.Engine.DB().Exec("DROP TABLE " + table.Name)
		if err != nil && reportErrors {
			t.Errorf("Could not drop table '%s': %s", table.Name, err)
		}
	}
}
