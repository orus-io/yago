package yago_test

// generated with yago. Better NOT to edit!

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/slicebit/qb"

	"github.com/m4rw3r/uuid"
	"github.com/orus-io/yago"
)

const (
	// SimpleStructID is the ID field name
	SimpleStructID = "ID"
	// SimpleStructIDColumnName is the ID field associated column name
	SimpleStructIDColumnName = "id"
	// SimpleStructName is the Name field name
	SimpleStructName = "Name"
	// SimpleStructNameColumnName is the Name field associated column name
	SimpleStructNameColumnName = "name"
)

const (
	// SimpleStructTableName is the SimpleStruct associated table name
	SimpleStructTableName = "simple_struct"
)

var simpleStructTable = qb.Table(
	SimpleStructTableName,
	qb.Column(SimpleStructIDColumnName, qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
	qb.Column(SimpleStructNameColumnName, qb.Varchar()).NotNull(),
	qb.UniqueKey(
		SimpleStructNameColumnName,
	),
)

var simpleStructType = reflect.TypeOf(SimpleStruct{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (SimpleStruct) StructType() reflect.Type {
	return simpleStructType
}

// SimpleStructModel provides direct access to helpers for SimpleStruct
// queries
type SimpleStructModel struct {
	mapper *SimpleStructMapper
	ID     yago.ScalarField
	Name   yago.ScalarField
}

// NewSimpleStructModel returns a new SimpleStructModel
func NewSimpleStructModel(meta *yago.Metadata) SimpleStructModel {
	mapper := NewSimpleStructMapper()
	meta.AddMapper(mapper)
	return SimpleStructModel{
		mapper: mapper,
		ID:     yago.NewScalarField(mapper.Table().C(SimpleStructIDColumnName)),
		Name:   yago.NewScalarField(mapper.Table().C(SimpleStructNameColumnName)),
	}
}

// GetMapper returns the associated SimpleStructMapper instance
func (m SimpleStructModel) GetMapper() yago.Mapper {
	return m.mapper
}

// NewSimpleStructMapper initialize a NewSimpleStructMapper
func NewSimpleStructMapper() *SimpleStructMapper {
	m := &SimpleStructMapper{}
	return m
}

// SimpleStructMapper is the SimpleStruct mapper
type SimpleStructMapper struct{}

// GetMapper returns itself
func (mapper *SimpleStructMapper) GetMapper() yago.Mapper {
	return mapper
}

// Name returns the mapper name
func (*SimpleStructMapper) Name() string {
	return "yago_test/SimpleStruct"
}

// Table returns the mapper table
func (*SimpleStructMapper) Table() *qb.TableElem {
	return &simpleStructTable
}

// StructType returns the reflect.Type of the mapped structure
func (SimpleStructMapper) StructType() reflect.Type {
	return simpleStructType
}

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper SimpleStructMapper) SQLValues(instance yago.MappedStruct, fields ...string) map[string]interface{} {
	s, ok := instance.(*SimpleStruct)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &SimpleStruct{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	allValues := len(fields) == 0
	m := make(map[string]interface{})
	if s.ID != 0 {
		m[SimpleStructIDColumnName] = s.ID
	}
	if allValues || yago.StringListContains(fields, SimpleStructName) {
		m[SimpleStructNameColumnName] = s.Name
	}
	return m
}

// FieldList returns the list of fields for a select
func (mapper SimpleStructMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		simpleStructTable.C(SimpleStructIDColumnName),
		simpleStructTable.C(SimpleStructNameColumnName),
	}
}

// ScanPKey scans the primary key only
func (mapper SimpleStructMapper) ScanPKey(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*SimpleStruct)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &SimpleStruct{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	return rows.Scan(
		&s.ID,
	)
}

// Scan a struct
func (mapper SimpleStructMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*SimpleStruct)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &SimpleStruct{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	return rows.Scan(
		&s.ID,
		&s.Name,
	)
}

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (SimpleStructMapper) AutoIncrementPKey() bool {
	return true
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (SimpleStructMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	s := instance.(*SimpleStruct)
	s.ID = value
}

// PKey returns the instance primary key values
func (mapper SimpleStructMapper) PKey(instance yago.MappedStruct) (values []interface{}) {
	str := instance.(*SimpleStruct)
	values = append(values, str.ID)

	return
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper SimpleStructMapper) PKeyClause(values []interface{}) qb.Clause {
	return simpleStructTable.C(SimpleStructIDColumnName).Eq(values[0])
}

const (
	// BaseStructID is the ID field name
	BaseStructID = "ID"
	// BaseStructIDColumnName is the ID field associated column name
	BaseStructIDColumnName = "id"
	// BaseStructCreatedAt is the CreatedAt field name
	BaseStructCreatedAt = "CreatedAt"
	// BaseStructCreatedAtColumnName is the CreatedAt field associated column name
	BaseStructCreatedAtColumnName = "created_at"
	// BaseStructUpdatedAt is the UpdatedAt field name
	BaseStructUpdatedAt = "UpdatedAt"
	// BaseStructUpdatedAtColumnName is the UpdatedAt field associated column name
	BaseStructUpdatedAtColumnName = "updated_at"
)

const (
	// PersonStructActive is the Active field name
	PersonStructActive = "Active"
	// PersonStructActiveColumnName is the Active field associated column name
	PersonStructActiveColumnName = "active"
	// PersonStructFirstName is the FirstName field name
	PersonStructFirstName = "FirstName"
	// PersonStructFirstNameColumnName is the FirstName field associated column name
	PersonStructFirstNameColumnName = "first_name"
	// PersonStructLastName is the LastName field name
	PersonStructLastName = "LastName"
	// PersonStructLastNameColumnName is the LastName field associated column name
	PersonStructLastNameColumnName = "last_name"
)

const (
	// PersonStructTableName is the PersonStruct associated table name
	PersonStructTableName = "person_struct"
)

var personStructTable = qb.Table(
	PersonStructTableName,
	qb.Column(PersonStructActiveColumnName, qb.Boolean()).NotNull(),
	qb.Column(PersonStructFirstNameColumnName, qb.Varchar()).NotNull(),
	qb.Column(PersonStructLastNameColumnName, qb.Varchar()).Null(),
	qb.Column(BaseStructIDColumnName, qb.UUID()).PrimaryKey().NotNull(),
	qb.Column(BaseStructCreatedAtColumnName, qb.Timestamp()).NotNull(),
	qb.Column(BaseStructUpdatedAtColumnName, qb.Timestamp()).NotNull(),
	qb.UniqueKey(
		PersonStructFirstNameColumnName,
	),
)

var personStructType = reflect.TypeOf(PersonStruct{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (PersonStruct) StructType() reflect.Type {
	return personStructType
}

// PersonStructModel provides direct access to helpers for PersonStruct
// queries
type PersonStructModel struct {
	mapper    *PersonStructMapper
	Active    yago.ScalarField
	FirstName yago.ScalarField
	LastName  yago.ScalarField
	ID        yago.ScalarField
	CreatedAt yago.ScalarField
	UpdatedAt yago.ScalarField
}

// NewPersonStructModel returns a new PersonStructModel
func NewPersonStructModel(meta *yago.Metadata) PersonStructModel {
	mapper := NewPersonStructMapper()
	meta.AddMapper(mapper)
	return PersonStructModel{
		mapper:    mapper,
		Active:    yago.NewScalarField(mapper.Table().C(PersonStructActiveColumnName)),
		FirstName: yago.NewScalarField(mapper.Table().C(PersonStructFirstNameColumnName)),
		LastName:  yago.NewScalarField(mapper.Table().C(PersonStructLastNameColumnName)),
		ID:        yago.NewScalarField(mapper.Table().C(BaseStructIDColumnName)),
		CreatedAt: yago.NewScalarField(mapper.Table().C(BaseStructCreatedAtColumnName)),
		UpdatedAt: yago.NewScalarField(mapper.Table().C(BaseStructUpdatedAtColumnName)),
	}
}

// GetMapper returns the associated PersonStructMapper instance
func (m PersonStructModel) GetMapper() yago.Mapper {
	return m.mapper
}

// NewPersonStructMapper initialize a NewPersonStructMapper
func NewPersonStructMapper() *PersonStructMapper {
	m := &PersonStructMapper{}
	return m
}

// PersonStructMapper is the PersonStruct mapper
type PersonStructMapper struct{}

// GetMapper returns itself
func (mapper *PersonStructMapper) GetMapper() yago.Mapper {
	return mapper
}

// Name returns the mapper name
func (*PersonStructMapper) Name() string {
	return "yago_test/PersonStruct"
}

// Table returns the mapper table
func (*PersonStructMapper) Table() *qb.TableElem {
	return &personStructTable
}

// StructType returns the reflect.Type of the mapped structure
func (PersonStructMapper) StructType() reflect.Type {
	return personStructType
}

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper PersonStructMapper) SQLValues(instance yago.MappedStruct, fields ...string) map[string]interface{} {
	s, ok := instance.(*PersonStruct)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &PersonStruct{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	allValues := len(fields) == 0
	m := make(map[string]interface{})
	if s.ID != (uuid.UUID{}) {
		m[BaseStructIDColumnName] = s.ID
	}
	if allValues || yago.StringListContains(fields, PersonStructActive) {
		m[PersonStructActiveColumnName] = s.Active
	}
	if allValues || yago.StringListContains(fields, PersonStructFirstName) {
		m[PersonStructFirstNameColumnName] = s.FirstName
	}
	if allValues || yago.StringListContains(fields, PersonStructLastName) {
		m[PersonStructLastNameColumnName] = s.LastName
	}
	if allValues || yago.StringListContains(fields, BaseStructCreatedAt) {
		m[BaseStructCreatedAtColumnName] = s.CreatedAt
	}
	if allValues || yago.StringListContains(fields, BaseStructUpdatedAt) {
		m[BaseStructUpdatedAtColumnName] = s.UpdatedAt
	}
	return m
}

// FieldList returns the list of fields for a select
func (mapper PersonStructMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		personStructTable.C(PersonStructActiveColumnName),
		personStructTable.C(PersonStructFirstNameColumnName),
		personStructTable.C(PersonStructLastNameColumnName),
		personStructTable.C(BaseStructIDColumnName),
		personStructTable.C(BaseStructCreatedAtColumnName),
		personStructTable.C(BaseStructUpdatedAtColumnName),
	}
}

// ScanPKey scans the primary key only
func (mapper PersonStructMapper) ScanPKey(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*PersonStruct)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &PersonStruct{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	return rows.Scan(
		&s.ID,
	)
}

// Scan a struct
func (mapper PersonStructMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*PersonStruct)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &PersonStruct{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	return rows.Scan(
		&s.Active,
		&s.FirstName,
		&s.LastName,
		&s.ID,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
}

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (PersonStructMapper) AutoIncrementPKey() bool {
	return false
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (PersonStructMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	panic("PersonStruct has no auto increment column in its pkey")
}

// PKey returns the instance primary key values
func (mapper PersonStructMapper) PKey(instance yago.MappedStruct) (values []interface{}) {
	str := instance.(*PersonStruct)
	values = append(values, str.ID)

	return
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PersonStructMapper) PKeyClause(values []interface{}) qb.Clause {
	return personStructTable.C(BaseStructIDColumnName).Eq(values[0])
}

const (
	// AutoIncBaseID is the ID field name
	AutoIncBaseID = "ID"
	// AutoIncBaseIDColumnName is the ID field associated column name
	AutoIncBaseIDColumnName = "id"
)

const (
	// AutoIncChildName is the Name field name
	AutoIncChildName = "Name"
	// AutoIncChildNameColumnName is the Name field associated column name
	AutoIncChildNameColumnName = "name"
)

const (
	// AutoIncChildTableName is the AutoIncChild associated table name
	AutoIncChildTableName = "auto_inc_child"
)

var autoIncChildTable = qb.Table(
	AutoIncChildTableName,
	qb.Column(AutoIncChildNameColumnName, qb.Varchar()).NotNull(),
	qb.Column(AutoIncBaseIDColumnName, qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
)

var autoIncChildType = reflect.TypeOf(AutoIncChild{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (AutoIncChild) StructType() reflect.Type {
	return autoIncChildType
}

// AutoIncChildModel provides direct access to helpers for AutoIncChild
// queries
type AutoIncChildModel struct {
	mapper *AutoIncChildMapper
	Name   yago.ScalarField
	ID     yago.ScalarField
}

// NewAutoIncChildModel returns a new AutoIncChildModel
func NewAutoIncChildModel(meta *yago.Metadata) AutoIncChildModel {
	mapper := NewAutoIncChildMapper()
	meta.AddMapper(mapper)
	return AutoIncChildModel{
		mapper: mapper,
		Name:   yago.NewScalarField(mapper.Table().C(AutoIncChildNameColumnName)),
		ID:     yago.NewScalarField(mapper.Table().C(AutoIncBaseIDColumnName)),
	}
}

// GetMapper returns the associated AutoIncChildMapper instance
func (m AutoIncChildModel) GetMapper() yago.Mapper {
	return m.mapper
}

// NewAutoIncChildMapper initialize a NewAutoIncChildMapper
func NewAutoIncChildMapper() *AutoIncChildMapper {
	m := &AutoIncChildMapper{}
	return m
}

// AutoIncChildMapper is the AutoIncChild mapper
type AutoIncChildMapper struct{}

// GetMapper returns itself
func (mapper *AutoIncChildMapper) GetMapper() yago.Mapper {
	return mapper
}

// Name returns the mapper name
func (*AutoIncChildMapper) Name() string {
	return "yago_test/AutoIncChild"
}

// Table returns the mapper table
func (*AutoIncChildMapper) Table() *qb.TableElem {
	return &autoIncChildTable
}

// StructType returns the reflect.Type of the mapped structure
func (AutoIncChildMapper) StructType() reflect.Type {
	return autoIncChildType
}

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper AutoIncChildMapper) SQLValues(instance yago.MappedStruct, fields ...string) map[string]interface{} {
	s, ok := instance.(*AutoIncChild)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &AutoIncChild{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	allValues := len(fields) == 0
	m := make(map[string]interface{})
	if s.ID != 0 {
		m[AutoIncBaseIDColumnName] = s.ID
	}
	if allValues || yago.StringListContains(fields, AutoIncChildName) {
		m[AutoIncChildNameColumnName] = s.Name
	}
	return m
}

// FieldList returns the list of fields for a select
func (mapper AutoIncChildMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		autoIncChildTable.C(AutoIncChildNameColumnName),
		autoIncChildTable.C(AutoIncBaseIDColumnName),
	}
}

// ScanPKey scans the primary key only
func (mapper AutoIncChildMapper) ScanPKey(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*AutoIncChild)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &AutoIncChild{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	return rows.Scan(
		&s.ID,
	)
}

// Scan a struct
func (mapper AutoIncChildMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*AutoIncChild)
	if !ok {
		panic(fmt.Sprintf(
			"Wrong struct type passed to the mapper. Expected &AutoIncChild{}, got %s",
			reflect.TypeOf(instance).Name(),
		))
	}
	return rows.Scan(
		&s.Name,
		&s.ID,
	)
}

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (AutoIncChildMapper) AutoIncrementPKey() bool {
	return true
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (AutoIncChildMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	s := instance.(*AutoIncChild)
	s.ID = value
}

// PKey returns the instance primary key values
func (mapper AutoIncChildMapper) PKey(instance yago.MappedStruct) (values []interface{}) {
	str := instance.(*AutoIncChild)
	values = append(values, str.ID)

	return
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper AutoIncChildMapper) PKeyClause(values []interface{}) qb.Clause {
	return autoIncChildTable.C(AutoIncBaseIDColumnName).Eq(values[0])
}
