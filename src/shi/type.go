package shi

import (
	"bufio"
	"fmt"
)

type Type interface {
	Dump(Value, *bufio.Writer, *VM) error
	Dup(Value, *VM) Value
	Name() Sym
	Write(Value, *bufio.Writer, *VM) error
}

type DataType[T any] interface {
	Type
	Zero() T
}

type BaseType[T any] struct {
	name Sym
}

func (self *BaseType[T]) Init(name Sym) {
	self.name = name
}

func (_ BaseType[T]) Dump(v Value, out *bufio.Writer, vm *VM) error {
	_, err := fmt.Fprintf(out, "%v", v.Data)
	return err
}

func (_ BaseType[T]) Dup(v Value, vm *VM) Value {
	return v
}

func (self BaseType[T]) Name() Sym {
	return self.name
}

func (self BaseType[T]) Write(v Value, out *bufio.Writer, vm *VM) error {
	return self.Dump(v, out, vm)
}

func (_ BaseType[T]) Zero() T {
	var v T
	return v
}
