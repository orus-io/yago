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

func TestInsertWithReturning(t *testing.T) {
	db, _, cleanup := initModelWithDriver(t, "postgres")
	defer cleanup()

	p := SimpleStruct{Name: "Test"}
	assert.EqualValues(t, 0, p.ID)
	assert.Nil(t, db.Insert(&p))

	assert.EqualValues(t, 1, p.ID)
}

func TestMarshaledField(t *testing.T) {
	db, model, cleanup := initModel(t)
	defer cleanup()

	p := PersonStruct{FirstName: "Martin", Gender: Male}

	db.Insert(&p)

	np := PersonStruct{}

	err := db.Query(
		model.PersonStruct,
	).Filter(
		model.PersonStruct.Gender.Eq(Male),
	).One(&np)
	assert.Nil(t, err)

	assert.Equal(t, Male, np.Gender)
	assert.Equal(t, "Martin", np.FirstName)
}
