package forms

import (
	"bufio"
	"shi/src/ops"
	"shi/src/shi"
)

type TLiteral struct {
	shi.BaseForm
	value shi.Value
}

func Literal(sloc shi.Sloc, value shi.Value) *TLiteral {
	return new(TLiteral).Init(sloc, value)
}

func (self *TLiteral) Init(sloc shi.Sloc, value shi.Value) *TLiteral {
	self.BaseForm.Init(sloc)
	self.value = value
	return self
}

func (self *TLiteral) Emit(in *shi.Forms, vm *shi.VM) error {
	vm.Emit(ops.Push(self.value))
	return nil
}

func (self *TLiteral) Dump(out *bufio.Writer, vm *shi.VM) error {
	return self.value.Dump(out, vm)
}
