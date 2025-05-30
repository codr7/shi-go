package forms

import (
	"bufio"
	"shi/src/shi"
)

type Id struct {
	shi.BaseForm
	name shi.Sym
}

func NewId(sloc shi.Sloc, name shi.Sym) *Id {
	return new(Id).Init(sloc, name)
}

func (self *Id) Init(sloc shi.Sloc, name shi.Sym) *Id {
	self.BaseForm.Init(sloc)
	self.name = name
	return self
}

func (self *Id) Emit(in *shi.Forms, vm *shi.VM) error {
	v := vm.CurrentLib().Find(self.name)

	if v == nil {
		return shi.NewEmitError(self.Sloc(),
			"Unknown identifier: %v",
			self.name.Value())
	}

	return v.Emit(self.Sloc(), in, vm)
}

func (self Id) Dump(out *bufio.Writer, vm *shi.VM) error {
	_, err := out.WriteString(self.name.Value())
	return err
}
