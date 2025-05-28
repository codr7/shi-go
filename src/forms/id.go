package forms

import (
	"bufio"
	"shi/src/shi"
	"shi/src/libs/core"
	"shi/src/ops"
)

type Id struct {
	shi.BaseForm
	name shi.Sym
}

func NewId(sloc shi.Sloc, name string) *Id {
	return new(Id).Init(sloc, shi.S(name))
}

func (self *Id) Init(sloc shi.Sloc, name shi.Sym) *Id {
	self.BaseForm.Init(sloc)
	self.name = name
	return self
}

func (self *Id) Emit(in *shi.Deque[shi.Form], vm *shi.VM) error {
	v := vm.CurrentLib().Find(self.name)

	if v == nil {
		return shi.NewEmitError(self.Sloc(),
			"Unknown identifier: %v",
			self.name.Value())
	}

	if v.Type == &core.Method {
		m := shi.Cast(*v, &core.Method)

		switch m.Notation() {
		case shi.Prefix:
			var args shi.Deque[shi.Form]
			
			for range m.Args() {
				if err := in.PopFront().Emit(&args, vm);
				err != nil {
					return err
				}
			}
		case shi.Infix:
			var args shi.Deque[shi.Form]

			for range m.Args()[1:] {
				if err := in.PopFront().Emit(&args, vm);
				err != nil {
					return err
				}
			}
		default:
			break
		}

		vm.Emit(ops.Call(self.Sloc(), m))
	} else {
		vm.Emit(ops.Push(*v))
	}
	
	return nil
}

func (self Id) Quote(vm *shi.VM) shi.Value {
	return shi.V(&core.Sym, self.name)
}

func (self Id) Dump(out *bufio.Writer) error {
	_, err := out.WriteString(self.name.Value())
	return err
}
