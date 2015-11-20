package util

import (
	"github.com/codegangsta/inject"
	"reflect"
)

const (
	Lifecycle_Singleton = 0
	Lifecycle_Transient = 1
)

var ObjectContainer = objectContainer{i: inject.New()}

type objectContainer struct {
	i inject.Injector
}

func (self *objectContainer) Map(val interface{}) {
	if initializer, ok := val.(Initializer); ok {
		self.i.Invoke(initializer.InitFunc())
	}
	self.i.Map(val)
}

func (self *objectContainer) MapTo(val interface{}, typ interface{}) {
	if initializer, ok := val.(Initializer); ok {
		self.i.Invoke(initializer.InitFunc())
	}
	self.i.MapTo(val, typ)
}

func (self *objectContainer) Invoke(fn interface{}) ([]reflect.Value, error) {
	return self.i.Invoke(fn)
}
