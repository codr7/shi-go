package shi

import (
	"fmt"
	"iter"
)

type Lib interface {
	All() iter.Seq2[Sym, Value]
	Bind(Sym, Value)
	Find(k Sym) *Value
	Import(source Lib, keys...Sym) error
	Init(name Sym, parentLib Lib)
	Name() Sym
}

type BaseLib struct {
	name Sym
	parentLib Lib
	bindings map[Sym]Value
}

func (self *BaseLib) Init(name Sym, parentLib Lib) {
	self.name = name
	self.parentLib = parentLib
	self.bindings = make(map[Sym]Value)
}

func (self *BaseLib) All() iter.Seq2[Sym, Value] {
	return func(yield func(Sym, Value) bool) {
		for k, v := range self.bindings {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (self *BaseLib) Bind(k Sym, v Value) {
	self.bindings[k] = v
}

func (self *BaseLib) Find(k Sym) *Value {
	v, ok := self.bindings[k]

	if !ok && self.parentLib != nil {
		return self.parentLib.Find(k)
	}

	if !ok {
		return nil
	}
	
	return &v
}

func (self *BaseLib) Import(source Lib, keys...Sym) error {
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

func (self *BaseLib) Name() Sym {
	return self.name
}
