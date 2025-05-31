package shi

import (
	"fmt"
	"io"
)

type Value struct {
	Type Type
	Data any
}

func V[T any](t DataType[T], d T) Value {
	return *new(Value).Init(t, d)
}

func (self *Value) Init(t Type, d any) *Value {
	self.Type = t
	self.Data = d
	return self
}

func (self Value) Dump(out io.Writer, vm *VM) error {
	return self.Type.Dump(self, out, vm)
}

func (self Value) Dup(vm *VM) Value {
	return self.Type.Dup(self, vm)
}

func (self Value) Emit(sloc Sloc, in *Forms, vm *VM) error {
	return self.Type.Emit(self, sloc, in, vm)
}

func Cast[T any](v Value, t DataType[T]) T {
	if v.Type != t {
		panic(fmt.Sprintf("Expected %v: %v", v.Type, t))
	}
	
	return v.Data.(T)
}
