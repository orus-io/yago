package yago

import (
	"database/sql"
	"reflect"

	"github.com/aacanakin/qb"
)

// MapperProvider is implemented by any struct that can provide a single mapper
type MapperProvider interface {
	GetMapper() Mapper
}

// Mapper links a mapped struct and table definition
type Mapper interface {
	Name() string
	Table() *qb.TableElem
	StructType() reflect.Type
	FieldList() []qb.Clause

	AutoIncrementPKey() bool
	LoadAutoIncrementPKeyValue(instance MappedStruct, value int64)
	SQLValues(instance MappedStruct, fields ...string) map[string]interface{}
	PKey(instance MappedStruct) []interface{}
	PKeyClause(values []interface{}) qb.Clause

	Scan(rows *sql.Rows, instance MappedStruct) error
}

// MappedStruct is implemented by all mapped structures
type MappedStruct interface {
	StructType() reflect.Type
}
