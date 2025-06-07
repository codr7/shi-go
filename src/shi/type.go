package shi

import (
	"fmt"
	"io"
	"iter"
)

type Type interface {
	AsBool(Value) bool
	Dump(Value, io.Writer, *VM) error
	Dup(Value, *VM) Value
	Emit(Value, Sloc, *Forms, *VM) error
	SubtypeOf(Type) bool
	Parents() iter.Seq[Type]
	Name() Symbol
	String() string
	Write(Value, io.Writer, *VM) error
}

type DataType[T any] interface {
	Type
	Zero() T
}

type BaseType[T any] struct {
	name    Symbol
	parents map[Type]bool
}

func (self *BaseType[T]) Init(name Symbol, parents ...Type) {
	self.name = name
	self.parents = make(map[Type]bool)

	for _, pt := range parents {
		self.parents[pt] = true

		for ppt := range pt.Parents() {
			self.parents[ppt] = true
		}
	}
}

func (_ *BaseType[T]) AsBool(_ Value) bool {
	return true
}

func (_ *BaseType[T]) Dump(v Value, out io.Writer, vm *VM) error {
	_, err := fmt.Fprint(out, v.Data)
	return err
}

func (_ *BaseType[T]) Dup(v Value, vm *VM) Value {
	return v
}

func (self *BaseType[T]) Name() Symbol {
	return self.name
}

func (self *BaseType[T]) String() string {
	return self.name.Value()
}

func (self *BaseType[T]) SubtypeOf(other Type) bool {
	_, ok := self.parents[other]
	return ok
}

func (self *BaseType[T]) Parents() iter.Seq[Type] {
	return func(yield func(Type) bool) {
		for t, _ := range self.parents {
			if !yield(t) {
				return
			}
		}
	}
}

func (self *BaseType[T]) Write(v Value, out io.Writer, vm *VM) error {
	return self.Dump(v, out, vm)
}

func (_ *BaseType[T]) Zero() T {
	var v T
	return v
}
