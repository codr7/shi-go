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
	SuperTypes() iter.Seq[Type]
	Name() Sym
	String() string
	Write(Value, io.Writer, *VM) error
}

type DataType[T any] interface {
	Type
	Zero() T
}

type BaseType[T any] struct {
	name Sym
	superTypes map[Type]bool
}

func (self *BaseType[T]) Init(name Sym, superTypes...Type) {
	self.name = name
	self.superTypes = make(map[Type]bool)
	
	for _, t := range superTypes {
		self.superTypes[t] = true
		
		for st := range t.SuperTypes() {
			self.superTypes[st] = true
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

func (self *BaseType[T]) Name() Sym {
	return self.name
}

func (self *BaseType[T]) String() string {
	return self.name.Value()
}

func (self *BaseType[T]) SubtypeOf(other Type) bool {
	_, ok := self.superTypes[other]
	return ok
}

func (self *BaseType[T]) SuperTypes() iter.Seq[Type] {
	return func(yield func(Type) bool) {
		for t, _ := range self.superTypes {
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
