package yagorm

import (
	"reflect"

	"github.com/aacanakin/qb"
)

// Metadata holds the table defs & mappers of a db
type Metadata struct {
	qbMeta *qb.MetaDataElem
	// store multiple mappers for structs, with a default one
	mappers map[reflect.Type]Mapper
}

// NewMetadata instanciate a Metadata
func NewMetadata(qbMeta *qb.MetaDataElem) *Metadata {
	return &Metadata{
		qbMeta:  qbMeta,
		mappers: make(map[reflect.Type]Mapper),
	}
}

// AddMapper add a mapper
func (m *Metadata) AddMapper(mapper Mapper) {
	m.qbMeta.AddTable(*mapper.Table())
	m.mappers[mapper.StructType()] = mapper
}

// GetMapper returns the default mapper of a mapped struct
func (m *Metadata) GetMapper(s MappedStruct) Mapper {
	return m.mappers[s.StructType()]
}
