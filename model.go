package yago

import (
	"github.com/slicebit/qb"
)

// A Model provides handy access to a struct definition.
type Model interface {
	GetMapper() Mapper
}

// ScalarField A simple scalar field
type ScalarField struct {
	Column qb.ColumnElem
}

// NewScalarField returns a new ScalarField
func NewScalarField(column qb.ColumnElem) ScalarField {
	return ScalarField{
		Column: column,
	}
}

// Accept calls the underlying column 'Accept'.
func (f ScalarField) Accept(context *qb.CompilerContext) string {
	return f.Column.Accept(context)
}

// Like returns a LIKE clause
func (f ScalarField) Like(pattern string) qb.Clause {
	return f.Column.Like(pattern)
}

// NotIn returns a NOT IN clause
func (f ScalarField) NotIn(values ...interface{}) qb.Clause {
	return f.Column.NotIn(values...)
}

// In returns a IN clause
func (f ScalarField) In(values ...interface{}) qb.Clause {
	return f.Column.In(values...)
}

// NotEq returns a != clause
func (f ScalarField) NotEq(value interface{}) qb.Clause {
	return f.Column.NotEq(value)
}

// Eq returns a = clause
func (f ScalarField) Eq(value interface{}) qb.Clause {
	return f.Column.Eq(value)
}

// Gt returns a > clause
func (f ScalarField) Gt(value interface{}) qb.Clause {
	return f.Column.Gt(value)
}

// Lt returns a < clause
func (f ScalarField) Lt(value interface{}) qb.Clause {
	return f.Column.Lt(value)
}

// Gte returns a >= clause
func (f ScalarField) Gte(value interface{}) qb.Clause {
	return f.Column.Gte(value)
}

// Lte returns a <= clause
func (f ScalarField) Lte(value interface{}) qb.Clause {
	return f.Column.Lte(value)
}
