package shi

import (
	"fmt"
	"io"
)

type SuperTypes = map[Type]bool

type Type interface {
	AddSuperTypes(SuperTypes)
	Dump(Value, io.Writer, *VM) error
	Dup(Value, *VM) Value
	Emit(Value, Sloc, *Forms, *VM) error
	SubtypeOf(Type) bool
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
	superTypes SuperTypes
}

func (self *BaseType[T]) Init(name Sym, superTypes...Type) {
	self.name = name
	self.superTypes = make(SuperTypes)

	for _, t := range superTypes {
		self.superTypes[t] = true
		t.AddSuperTypes(self.superTypes)
	}
}

func (self *BaseType[T]) AddSuperTypes(out SuperTypes) {
	for t, _ := range self.superTypes {
		out[t] = true
	}
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

func (self *BaseType[T]) Write(v Value, out io.Writer, vm *VM) error {
	return self.Dump(v, out, vm)
}

func (_ *BaseType[T]) Zero() T {
	var v T
	return v
}
