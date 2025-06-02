package forms

import (
	"bufio"
	"shi/src/shi"
)

type TScope struct {
	shi.BaseForm
	forms shi.Forms
}

func Scope(sloc shi.Sloc, forms shi.Forms) *TScope {
	return new(TScope).Init(sloc, forms)
}

func (self *TScope) Init(sloc shi.Sloc, forms shi.Forms) *TScope {
	self.BaseForm.Init(sloc)
	self.forms = forms
	return self
}

func (self *TScope) Emit(in *shi.Forms, vm *shi.VM) error {
	fs := self.forms
	return vm.WithLibrary(nil, func () error {
		return EmitAll(&fs, vm)
	})
}

func (self *TScope) Dump(out *bufio.Writer, vm *shi.VM) error {
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
