package forms

import (
	"bufio"
	"shi/src/shi"
	"shi/src/libs/core"
)

type Scope struct {
	shi.BaseForm
	forms shi.Deque[shi.Form]
}

func NewScope(sloc shi.Sloc, forms shi.Deque[shi.Form]) *Scope {
	return new(Scope).Init(sloc, forms)
}

func (self *Scope) Init(sloc shi.Sloc, forms shi.Deque[shi.Form]) *Scope {
	self.BaseForm.Init(sloc)
	self.forms = forms
	return self
}

func (self *Scope) Emit(in *shi.Deque[shi.Form], vm *shi.VM) error {
	fs := self.forms
	return EmitAll(&fs, vm)
}

func (self Scope) Quote(vm *shi.VM) shi.Value {
	//TODO emit list
	return core.NIL
}

func (self Scope) Dump(out *bufio.Writer) error {
	if _, err := out.WriteRune('('); err != nil {
		return err
	}

	for i, f := range self.forms.Items {
		if i > 0 {
			if _, err := out.WriteRune(' '); err != nil {
				return err
			}
		}

		if err := f.Dump(out); err != nil {
			return err
		}
	}
	
	if _, err := out.WriteRune(')'); err != nil {
		return err
	}

	return nil
}
