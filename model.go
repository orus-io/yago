package yago

import (
	"github.com/aacanakin/qb"
)

type Model interface {
	GetMapper() Mapper
}

type ScalarField struct {
	Column qb.ColumnElem
}

func NewScalarField(column qb.ColumnElem) ScalarField {
	return ScalarField{
		Column: column,
	}
}

func (f ScalarField) Eq(value interface{}) qb.Clause {
	return f.Column.Eq(value)
}
