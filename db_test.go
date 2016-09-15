package yago_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTx(t *testing.T) {
	db, model, cleanup := initModel(t)
	defer cleanup()

	p := PersonStruct{FirstName: "Malcom"}

	{
		// Make an insert in a transaction and rollback it
		tx, err := db.Begin()
		assert.Nil(t, err)

		assert.Nil(t, tx.Insert(&p))

		exists, err := tx.Query(model.PersonStruct).Exists()
		assert.Nil(t, err)
		assert.True(t, exists)

		tx.Rollback()
	}

	{
		// Check that the insert was rolled back
		exists, err := db.Query(model.PersonStruct).Exists()
		assert.Nil(t, err)
		assert.False(t, exists)
	}

	{
		// Redo an insert in a transaction and commit it
		tx, err := db.Begin()
		assert.Nil(t, err)

		assert.Nil(t, tx.Insert(&p))

		exists, err := tx.Query(model.PersonStruct).Exists()
		assert.Nil(t, err)
		assert.True(t, exists)

		tx.Commit()
	}

	{
		// Check that the insert was committed
		exists, err := db.Query(model.PersonStruct).Exists()
		assert.Nil(t, err)
		assert.True(t, exists)
	}
}
