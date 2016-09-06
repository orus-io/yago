package yago

// BeforeInsert can be implemented by structs that need a before insert callback
type BeforeInsert interface {
	BeforeInsert(db *DB)
}

// AfterInsert can be implemented by structs that need a after insert callback
type AfterInsert interface {
	AfterInsert(db *DB)
}

// BeforeUpdate can be implemented by structs that need a before update callback
type BeforeUpdate interface {
	BeforeUpdate(db *DB)
}

// AfterUpdate can be implemented by structs that need a after update callback
type AfterUpdate interface {
	AfterUpdate(db *DB)
}

// BeforeDelete can be implemented by structs that need a before delete callback
type BeforeDelete interface {
	BeforeDelete(db *DB)
}

// AfterDelete can be implemented by structs that need a after delete callback
type AfterDelete interface {
	AfterDelete(db *DB)
}

func beforeInsertInterfaceCallback(db *DB, s MappedStruct) {
	if i, ok := s.(BeforeInsert); ok {
		i.BeforeInsert(db)
	}
}

func afterInsertInterfaceCallback(db *DB, s MappedStruct) {
	if i, ok := s.(AfterInsert); ok {
		i.AfterInsert(db)
	}
}

func beforeUpdateInterfaceCallback(db *DB, s MappedStruct) {
	if i, ok := s.(BeforeUpdate); ok {
		i.BeforeUpdate(db)
	}
}

func afterUpdateInterfaceCallback(db *DB, s MappedStruct) {
	if i, ok := s.(AfterUpdate); ok {
		i.AfterUpdate(db)
	}
}

func beforeDeleteInterfaceCallback(db *DB, s MappedStruct) {
	if i, ok := s.(BeforeDelete); ok {
		i.BeforeDelete(db)
	}
}

func afterDeleteInterfaceCallback(db *DB, s MappedStruct) {
	if i, ok := s.(AfterDelete); ok {
		i.AfterDelete(db)
	}
}

func init() {
	DefaultCallbacks.AfterInsert.Add(
		Callback("yago:interface", afterInsertInterfaceCallback))
	DefaultCallbacks.BeforeInsert.Add(
		Callback("yago:interface", beforeInsertInterfaceCallback))
	DefaultCallbacks.AfterUpdate.Add(
		Callback("yago:interface", afterUpdateInterfaceCallback))
	DefaultCallbacks.BeforeUpdate.Add(
		Callback("yago:interface", beforeUpdateInterfaceCallback))
	DefaultCallbacks.AfterDelete.Add(
		Callback("yago:interface", afterDeleteInterfaceCallback))
	DefaultCallbacks.BeforeDelete.Add(
		Callback("yago:interface", beforeDeleteInterfaceCallback))
}
