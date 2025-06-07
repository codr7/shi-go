package forms

import (
	"bufio"
	"shi/src/shi"
)

type TId struct {
	shi.BaseForm
	name shi.Symbol
}

func Id(sloc shi.Sloc, name shi.Symbol) *TId {
	return new(TId).Init(sloc, name)
}

func (self *TId) Init(sloc shi.Sloc, name shi.Symbol) *TId {
	self.BaseForm.Init(sloc)
	self.name = name
	return self
}

func (self *TId) Name() shi.Symbol {
	return self.name
}

func (self *TId) Emit(in *shi.Forms, vm *shi.VM) error {
	v := vm.CurrentLibrary().Find(self.name)

	if v == nil {
		return shi.NewEmitError(self.Sloc(),
			"Unknown identifier: %v",
			self.name.Value())
	}

	return v.Emit(self.Sloc(), in, vm)
}

func (self *TId) Dump(out *bufio.Writer, vm *shi.VM) error {
	_, err := out.WriteString(self.name.Value())
	return err
}
