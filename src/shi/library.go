package shi

import (
	"fmt"
	"iter"
)

type Library interface {
	All() iter.Seq2[Symbol, Value]
	Bind(Symbol, Value)
	Find(k Symbol) *Value
	ImportFrom(source Library, keys ...Symbol) error
	Init(name Symbol, parent Library)
	Name() Symbol
}

type BaseLibrary struct {
	name     Symbol
	parent   Library
	bindings map[Symbol]Value
}

func (self *BaseLibrary) Init(name Symbol, parent Library) {
	self.name = name
	self.parent = parent
	self.bindings = make(map[Symbol]Value)
}

func (self *BaseLibrary) All() iter.Seq2[Symbol, Value] {
	return func(yield func(Symbol, Value) bool) {
		for k, v := range self.bindings {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (self *BaseLibrary) Bind(k Symbol, v Value) {
	self.bindings[k] = v
}

func (self *BaseLibrary) Find(k Symbol) *Value {
	v, ok := self.bindings[k]

	if !ok && self.parent != nil {
		return self.parent.Find(k)
	}

	if !ok {
		return nil
	}

	return &v
}

func (self *BaseLibrary) ImportFrom(source Library, keys ...Symbol) error {
	if len(keys) == 0 {
		for k, v := range source.All() {
			self.Bind(k, v)
		}
	} else {
		for _, k := range keys {
			v := source.Find(k)

			if v == nil {
				return fmt.Errorf("Unknown identifier: %v", k)
			}

			self.Bind(k, *v)
		}
	}

	return nil
}

func (self *BaseLibrary) Name() Symbol {
	return self.name
}
