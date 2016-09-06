package yago

import "fmt"

// CallbackFunc is the type of the callback functions
type CallbackFunc func(db *DB, s MappedStruct)

// DefaultCallbacks contains the default callbacks every database
// will be initialized with.
var DefaultCallbacks Callbacks

// Callback returns a new CallbackDef
func Callback(name string, callback CallbackFunc) CallbackDef {
	return CallbackDef{
		name:     name,
		callback: callback,
	}
}

// CallbackDef is a single callback definition
type CallbackDef struct {
	name     string
	callback CallbackFunc

	before string
	after  string
}

// CallbackList containts a list of same-type callbacks
type CallbackList struct {
	defs      []CallbackDef
	callbacks []CallbackFunc
}

// Callbacks containts a set of callbacks
type Callbacks struct {
	BeforeInsert CallbackList
	AfterInsert  CallbackList
	BeforeUpdate CallbackList
	AfterUpdate  CallbackList
	BeforeDelete CallbackList
	AfterDelete  CallbackList
}

// After which callback to be called
func (c CallbackDef) After(name string) CallbackDef {
	c.after = name
	return c
}

// Before which callback to be called
func (c CallbackDef) Before(name string) CallbackDef {
	c.before = name
	return c
}

// Call runs all the callbacks with the passed arguments
func (c CallbackList) Call(db *DB, s MappedStruct) {
	for _, f := range c.callbacks {
		f(db, s)
	}
}

// Add a new callback
func (c *CallbackList) Add(def CallbackDef) {
	if def.after == "" && def.before == "" && len(c.defs) != 0 {
		def.after = c.defs[len(c.defs)-1].name
	}
	c.defs = AddCallbackSorted(c.defs, def)
	c.callbacks = []CallbackFunc{}
	for _, d := range c.defs {
		c.callbacks = append(c.callbacks, d.callback)
	}
}

// Get a callback by name
func (c *CallbackList) Get(name string) *CallbackDef {
	for i := range c.defs {
		if c.defs[i].name == name {
			return &c.defs[i]
		}
	}
	return nil
}

// Remove a callback by name
func (c *CallbackList) Remove(name string) {
	for i := range c.defs {
		if c.defs[i].name == name {
			c.defs = append(c.defs[0:i], c.defs[i+1:len(c.defs)]...)
			break
		}
	}
}

// AddCallbackSorted insert a CallbackDef respecting the before/after args
func AddCallbackSorted(defs []CallbackDef, def CallbackDef) []CallbackDef {
	if len(defs) == 0 {
		return []CallbackDef{def}
	}

	// find a spot
	index := -1
	for i, d := range defs {
		if def.after == d.name {
			index = i + 1
			break
		} else if def.before == d.name {
			index = i
			break
		}
	}
	if index == -1 {
		panic("yago.AddCallbackSorted: No spot candidate for callbacks")
	}

	// Make sure the spot is consistent with all constraints
	for i, d := range defs {
		if i < index {
			if d.after == def.name || def.before == d.name {
				panic(fmt.Sprintf(
					"yago.AddCallbackSorted: No consistent spot for callback %s",
					def.name,
				))
			}
		} else {
			if d.before == def.name || def.after == d.name {
				panic(fmt.Sprintf(
					"yago.AddCallbackSorted: No consistent spot for callback %s",
					def.name,
				))
			}
		}
	}

	if index == len(defs) {
		defs = append(defs, def)
	} else {
		defs = append(
			defs[0:index],
			append([]CallbackDef{def}, defs[index:]...)...,
		)
	}
	return defs
}
