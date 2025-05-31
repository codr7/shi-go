package forms

import (
	"bufio"
	"shi/src/shi"
)

type TId struct {
	shi.BaseForm
	name shi.Sym
}

func Id(sloc shi.Sloc, name shi.Sym) *TId {
	return new(TId).Init(sloc, name)
}

func (self *TId) Init(sloc shi.Sloc, name shi.Sym) *TId {
	self.BaseForm.Init(sloc)
	self.name = name
	return self
}

func (self *TId) Name() shi.Sym {
	return self.name
}

func (self *TId) Emit(in *shi.Forms, vm *shi.VM) error {
	v := vm.CurrentLib().Find(self.name)

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
