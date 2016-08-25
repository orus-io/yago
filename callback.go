package yago

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
	c.defs = append(c.defs, def)
	c.reorder()
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

// reorder the callbacks
func (c *CallbackList) reorder() {
	// XXX This implementation is obviously broken, as it ignores 'before' and 'after'
	c.callbacks = []CallbackFunc{}
	for _, def := range c.defs {
		c.callbacks = append(c.callbacks, def.callback)
	}
}
