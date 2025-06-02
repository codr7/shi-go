package shi

import (
	"fmt"
	"iter"
)

type Lib interface {
	All() iter.Seq2[Symbol, Value]
	Bind(Symbol, Value)
	Find(k Symbol) *Value
	Import(source Lib, keys...Symbol) error
	Init(name Symbol, parentLib Lib)
	Name() Symbol
}

type BaseLib struct {
	name Symbol
	parentLib Lib
	bindings map[Symbol]Value
}

func (self *BaseLib) Init(name Symbol, parentLib Lib) {
	self.name = name
	self.parentLib = parentLib
	self.bindings = make(map[Symbol]Value)
}

func (self *BaseLib) All() iter.Seq2[Symbol, Value] {
	return func(yield func(Symbol, Value) bool) {
		for k, v := range self.bindings {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (self *BaseLib) Bind(k Symbol, v Value) {
	self.bindings[k] = v
}

func (self *BaseLib) Find(k Symbol) *Value {
	v, ok := self.bindings[k]

	if !ok && self.parentLib != nil {
		return self.parentLib.Find(k)
	}

	if !ok {
		return nil
	}
	
	return &v
}

func (self *BaseLib) Import(source Lib, keys...Symbol) error {
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

func (self *BaseLib) Name() Symbol {
	return self.name
}
