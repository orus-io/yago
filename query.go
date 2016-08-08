package yagorm

type Query struct {
	m          *Metadata
	rootMapper Mapper
}

func NewQuery(metadata *Metadata, mapper Mapper) (q *Query) {
	return &Query{m: metadata, rootMapper: mapper}
}

func NewQueryFromStruct(metadata *Metadata, s MappedStruct) (q *Query) {
	return NewQuery(metadata, metadata.GetMapper(s))
}
