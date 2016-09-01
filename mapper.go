package yago

import (
	"database/sql"
	"reflect"

	"github.com/aacanakin/qb"
)

// Mapper links a mapped struct and table definition
type Mapper interface {
	Name() string
	Table() *qb.TableElem
	StructType() reflect.Type
	FieldList() []qb.Clause

	AutoIncrementPKey() bool
	LoadAutoIncrementPKeyValue(instance MappedStruct, value int64)
	PKeyClause(instance MappedStruct) qb.Clause
	SQLValues(instance MappedStruct) map[string]interface{}

	Scan(rows *sql.Rows, instance MappedStruct) error
}

// MappedStruct is implemented by all mapped structures
type MappedStruct interface {
	StructType() reflect.Type
}
