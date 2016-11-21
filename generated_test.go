package yago_test

import (
	"testing"

	"github.com/m4rw3r/uuid"
	"github.com/orus-io/yago"
	"github.com/stretchr/testify/assert"
)

func getUUID() uuid.UUID {
	id, err := uuid.V4()
	if err != nil {
		panic(err)
	}
	return id
}

var (
	uuid1          = getUUID()
	SQLValuesTests = []struct {
		m      yago.Mapper
		s      yago.MappedStruct
		fields []string
		expect map[string]interface{}
	}{
		{
			NewPersonStructMapper(),
			&PersonStruct{BaseStruct: BaseStruct{ID: uuid1}, FirstName: "John", LastName: "Reece"},
			[]string{PersonStructFirstName},
			map[string]interface{}{
				BaseStructIDColumnName:          uuid1,
				PersonStructFirstNameColumnName: "John",
			},
		},
		{
			NewPersonStructMapper(),
			&PersonStruct{FirstName: "John", LastName: "Reece"},
			[]string{PersonStructLastName},
			map[string]interface{}{
				PersonStructLastNameColumnName: "Reece",
			},
		},
		{
			NewPersonStructMapper(),
			&PersonStruct{FirstName: "John", LastName: "Reece", Gender: Male},
			[]string{PersonStructLastName, PersonStructGender},
			map[string]interface{}{
				PersonStructLastNameColumnName: "Reece",
				PersonStructGenderColumnName:   []byte("male"),
			},
		},
	}
)

func TestSQLValues(t *testing.T) {
	for _, tt := range SQLValuesTests {
		assert.Equal(t, tt.expect, tt.m.SQLValues(tt.s, tt.fields...))
	}
}

func TestInheritedAutoincPkey(t *testing.T) {
	assert.True(t, NewAutoIncChildMapper().AutoIncrementPKey())
}
