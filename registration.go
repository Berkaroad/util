package util

import (
	"github.com/codegangsta/inject"
	"reflect"
)

var TypeRegistration = typeRegistration{i: inject.New()}

type Initializer interface {
	InitFunc() interface{}
}

type typeRegistration struct {
	i inject.Injector
}

func (self *typeRegistration) Map(val interface{}) {
	if initializer, ok := val.(Initializer); ok {
		self.i.Invoke(initializer.InitFunc())
	}
	self.i.Map(val)
}

func (self *typeRegistration) MapTo(val interface{}, typ interface{}) {
	if initializer, ok := val.(Initializer); ok {
		self.i.Invoke(initializer.InitFunc())
	}
	self.i.MapTo(val, typ)
}

func (self *typeRegistration) Invoke(fn interface{}) ([]reflect.Value, error) {
	return self.i.Invoke(fn)
}
