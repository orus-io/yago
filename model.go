package yago

import (
	"encoding"

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

// MarshaledScalarField A text marshaled scalar field
type MarshaledScalarField struct {
	ScalarField
}

// NewScalarField returns a new ScalarField
func NewScalarField(column qb.ColumnElem) ScalarField {
	return ScalarField{
		Column: column,
	}
}

// NewMarshaledScalarField returns a new ScalarField
func NewMarshaledScalarField(column qb.ColumnElem) MarshaledScalarField {
	return MarshaledScalarField{NewScalarField(column)}
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

// marshalValue marshals the value if it implements encoding.TextMarshaler
func (f MarshaledScalarField) marshalValue(value interface{}) interface{} {
	tm, ok := value.(encoding.TextMarshaler)
	if ok {
		b, err := tm.MarshalText()
		if err != nil {
			panic(err)
		}
		return b
	}
	return value
}

func (f MarshaledScalarField) marshalValues(values []interface{}) []interface{} {
	var marshaled = make([]interface{}, len(values))
	for _, value := range values {
		marshaled = append(marshaled, f.marshalValue(value))
	}
	return marshaled
}

// NotIn returns a NOT IN clause
func (f MarshaledScalarField) NotIn(values ...interface{}) qb.Clause {
	return f.Column.NotIn(f.marshalValues(values)...)
}

// In returns a IN clause
func (f MarshaledScalarField) In(values ...interface{}) qb.Clause {
	return f.Column.In(f.marshalValues(values)...)
}

// NotEq returns a != clause
func (f MarshaledScalarField) NotEq(value interface{}) qb.Clause {
	return f.Column.NotEq(f.marshalValue(value))
}

// Eq returns a = clause
func (f MarshaledScalarField) Eq(value interface{}) qb.Clause {
	return f.Column.Eq(f.marshalValue(value))
}

// Gt returns a > clause
func (f MarshaledScalarField) Gt(value interface{}) qb.Clause {
	return f.Column.Gt(f.marshalValue(value))
}

// Lt returns a < clause
func (f MarshaledScalarField) Lt(value interface{}) qb.Clause {
	return f.Column.Lt(f.marshalValue(value))
}

// Gte returns a >= clause
func (f MarshaledScalarField) Gte(value interface{}) qb.Clause {
	return f.Column.Gte(f.marshalValue(value))
}

// Lte returns a <= clause
func (f MarshaledScalarField) Lte(value interface{}) qb.Clause {
	return f.Column.Lte(f.marshalValue(value))
}
