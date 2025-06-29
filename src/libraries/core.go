package libraries

import (
	"shi/src/forms"
	"shi/src/libraries/core"
	"shi/src/operations"
	"shi/src/shi"
)

type TCore struct {
	shi.BaseLibrary
}

var Core TCore

func init() {
	Core.Init(shi.S("core"), nil)
}

func (self *TCore) Init(name shi.Symbol, parentLibrary shi.Library) {
	self.BaseLibrary.Init(name, nil)

	BindType(self, &core.Bool)
	BindType(self, &core.Int)
	BindType(self, &core.Macro)
	BindType(self, &core.Meta)
	BindType(self, &core.Method)

	self.Bind(shi.S("T"), core.T)
	self.Bind(shi.S("F"), core.F)

	BindMethod(self, shi.S("+"),
		shi.MethodArgs{}.
			Add(shi.S("x"), &core.Int).
			Add(shi.S("y"), &core.Int),
		func(sloc shi.Sloc, stack *shi.Values, vm *shi.VM) error {
			y := shi.Cast(stack.Pop(), &core.Int)
			x := stack.Peek()
			x.Data = shi.Cast(*x, &core.Int) + y
			return nil
		})

	BindMethod(self, shi.S("-"),
		shi.MethodArgs{}.
			Add(shi.S("x"), &core.Int).
			Add(shi.S("y"), &core.Int),
		func(sloc shi.Sloc, stack *shi.Values, vm *shi.VM) error {
			y := shi.Cast(stack.Pop(), &core.Int)
			x := stack.Peek()
			x.Data = shi.Cast(*x, &core.Int) - y
			return nil
		})

	BindMethod(self, shi.S("<"),
		shi.MethodArgs{}.
			Add(shi.S("x"), &core.Int).
			Add(shi.S("y"), &core.Int),
		func(sloc shi.Sloc, stack *shi.Values, vm *shi.VM) error {
			y := shi.Cast(stack.Pop(), &core.Int)
			x := stack.Peek()
			x.Init(&core.Bool, shi.Cast(*x, &core.Int) < y)
			return nil
		})

	BindMacro(self, shi.S("if"),
		[]string{"cond", "branch"},
		func(sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
			if err := in.PopFront().Emit(in, vm); err != nil {
				return err
			}

			branchEnd := shi.NewLabel()
			vm.Emit(operations.Branch(branchEnd))

			if err := in.PopFront().Emit(in, vm); err != nil {
				return err
			}

			if f, ok := in.PeekFront().(*forms.TId); ok && f.Name() == shi.S("else") {
				in.PopFront()
				elseEnd := shi.NewLabel()
				vm.Emit(operations.Goto(elseEnd))
				branchEnd.Pc = vm.EmitPc()

				if err := in.PopFront().Emit(in, vm); err != nil {
					return err
				}

				elseEnd.Pc = vm.EmitPc()
			} else {
				branchEnd.Pc = vm.EmitPc()
			}

			return nil
		})
}
