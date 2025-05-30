package forms

import (
	"bufio"
	"shi/src/ops"
	"shi/src/shi"
)

type Literal struct {
	shi.BaseForm
	value shi.Value
}

func NewLiteral(sloc shi.Sloc, value shi.Value) *Literal {
	return new(Literal).Init(sloc, value)
}

func (self *Literal) Init(sloc shi.Sloc, value shi.Value) *Literal {
	self.BaseForm.Init(sloc)
	self.value = value
	return self
}

func (self Literal) Emit(in *shi.Forms, vm *shi.VM) error {
	vm.Emit(ops.Push(self.value))
	return nil
}

func (self Literal) Dump(out *bufio.Writer, vm *shi.VM) error {
	return self.value.Dump(out, vm)
}
