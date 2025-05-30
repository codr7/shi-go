package forms

import (
	"bufio"
	"shi/src/shi"
)

type Scope struct {
	shi.BaseForm
	forms shi.Forms
}

func NewScope(sloc shi.Sloc, forms shi.Forms) *Scope {
	return new(Scope).Init(sloc, forms)
}

func (self *Scope) Init(sloc shi.Sloc, forms shi.Forms) *Scope {
	self.BaseForm.Init(sloc)
	self.forms = forms
	return self
}

func (self Scope) Emit(in *shi.Forms, vm *shi.VM) error {
	fs := self.forms
	return vm.WithLib(nil, func () error {
		return EmitAll(&fs, vm)
	})
}

func (self Scope) Dump(out *bufio.Writer, vm *shi.VM) error {
	if _, err := out.WriteRune('('); err != nil {
		return err
	}

	for i, f := range self.forms.Items {
		if i > 0 {
			if _, err := out.WriteRune(' '); err != nil {
				return err
			}
		}

		if err := f.Dump(out, vm); err != nil {
			return err
		}
	}
	
	if _, err := out.WriteRune(')'); err != nil {
		return err
	}

	return nil
}
