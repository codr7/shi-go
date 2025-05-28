package libs

import (
	"shi/src/shi"
	"shi/src/libs/core"
)

type TCore struct {
	shi.BaseLib
}

var Core TCore

func init() {
	Core.Init(shi.S("core"), nil)
}

func (self *TCore) Init(name shi.Sym, parentLib shi.Lib) {
	self.BaseLib.Init(name, nil)

	BindType(self, &core.Bool)
	BindType(self, &core.Int)
	BindType(self, &core.Meta)
	BindType(self, &core.Method)
	BindType(self, &core.Nil)
	BindType(self, &core.Sym)

	self.Bind(shi.S("T"), shi.V(&core.Bool, true))
	self.Bind(shi.S("F"), shi.V(&core.Bool, false))

	BindMethod(self, shi.S("+"), shi.Infix,
		shi.MethodArgs{}.
			Add(shi.S("x"), &core.Int).
			Add(shi.S("y"), &core.Int),
		func (sloc shi.Sloc, vm *shi.VM) error {
			y := shi.Cast(vm.Stack.Pop(), &core.Int)
			x := vm.Stack.Peek()
			x.Data = shi.Cast(*x, &core.Int) + y
			return nil
		})
}
