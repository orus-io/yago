package yago

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeStruct struct{ lastCall string }

// Implements the MappedStruct interface
func (FakeStruct) StructType() reflect.Type { return reflect.TypeOf(FakeStruct{}) }

func (s *FakeStruct) BeforeInsert(db *DB) { s.lastCall = "BeforeInsert" }
func (s *FakeStruct) AfterInsert(db *DB)  { s.lastCall = "AfterInsert" }
func (s *FakeStruct) BeforeUpdate(db *DB) { s.lastCall = "BeforeUpdate" }
func (s *FakeStruct) AfterUpdate(db *DB)  { s.lastCall = "AfterUpdate" }
func (s *FakeStruct) BeforeDelete(db *DB) { s.lastCall = "BeforeDelete" }
func (s *FakeStruct) AfterDelete(db *DB)  { s.lastCall = "AfterDelete" }

func TestCallbackDefInit(t *testing.T) {
	var db DB
	var called bool
	var str FakeStruct
	var cb CallbackFunc = func(db *DB, s MappedStruct) { called = true }

	assert.Equal(t, "test", Callback("test", cb).name)
	assert.Equal(t, "other", Callback("test", cb).After("other").after)
	assert.Equal(t, "other", Callback("test", cb).Before("other").before)
	Callback("test", cb).callback(&db, &str)
	assert.True(t, called)
}

func TestCallbackList(t *testing.T) {
	// var db DB
	var callCount int
	var cb CallbackFunc = func(db *DB, s MappedStruct) { callCount++ }

	var list CallbackList
	list.Add(Callback("test1", cb))
	list.Add(Callback("test2", cb))
	assert.Nil(t, list.Get("invalid"))
	assert.Equal(t, "test1", list.Get("test1").name)
	assert.Equal(t, "test2", list.Get("test2").name)
	assert.Equal(t, "test1", list.Get("test2").after)

	list.Call(nil, nil)
	assert.Equal(t, 2, callCount)

	list.Remove("test1")
	assert.Nil(t, list.Get("test1"))
}

func TestDefaultCallbackInterface(t *testing.T) {
	var s FakeStruct
	DefaultCallbacks.BeforeInsert.Call(nil, &s)
	assert.Equal(t, "BeforeInsert", s.lastCall)
	DefaultCallbacks.AfterInsert.Call(nil, &s)
	assert.Equal(t, "AfterInsert", s.lastCall)
	DefaultCallbacks.BeforeUpdate.Call(nil, &s)
	assert.Equal(t, "BeforeUpdate", s.lastCall)
	DefaultCallbacks.AfterUpdate.Call(nil, &s)
	assert.Equal(t, "AfterUpdate", s.lastCall)
	DefaultCallbacks.BeforeDelete.Call(nil, &s)
	assert.Equal(t, "BeforeDelete", s.lastCall)
	DefaultCallbacks.AfterDelete.Call(nil, &s)
	assert.Equal(t, "AfterDelete", s.lastCall)
}

func TestCallbackListReorder(t *testing.T) {
	var calls []string
	getCb := func(name string) CallbackFunc {
		return func(db *DB, s MappedStruct) { calls = append(calls, name) }
	}

	var list CallbackList
	list.Add(Callback("test2", getCb("test2")))
	list.Add(Callback("test1", getCb("test1")).Before("test2"))
	list.Add(Callback("test3", getCb("test3")).After("test2"))

	list.Call(nil, nil)

	assert.Equal(t, []string{"test1", "test2", "test3"}, calls)

	list = CallbackList{}
	list.Add(Callback("test1", nil))
	assert.Panics(t, func() {
		list.Add(Callback("test2", nil).Before("none"))
	})
	list.Add(Callback("test2", nil).Before("test1"))
	assert.Panics(t, func() {
		list.Add(Callback("test3", nil).After("test1").Before("test2"))
	})
	assert.Panics(t, func() {
		list.Add(Callback("test3", nil).Before("test1").After("test1"))
	})

}
