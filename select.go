package yagorm

import (
	"github.com/aacanakin/qb"
)

// Select instanciate a qb.SelectStmt for a given Struct
func Select(meta *Metadata, st MappedStruct) qb.SelectStmt {
	return MapperSelect(meta.GetMapper(st))
}

// MapperSelect instanciate a qb.SelectStmt for a given mapper
func MapperSelect(mapper Mapper) qb.SelectStmt {
	return qb.Select().From(*mapper.Table())
}
