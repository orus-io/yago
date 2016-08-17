package yagorm

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
	Values(instance MappedStruct) map[string]interface{}
	FieldList() []qb.Clause

	Scan(rows *sql.Rows, instance MappedStruct) error
}

// MappedStruct is implemented by all mapped structures
type MappedStruct interface {
	StructType() reflect.Type
}
