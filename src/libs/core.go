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

	self.Bind(shi.S("T"), core.T)
	self.Bind(shi.S("F"), core.F)
	self.Bind(shi.S("N"), core.N)

	BindMethod(self, shi.S("+"),
		shi.MethodArgs{}.
			Add(shi.S("x"), &core.Int).
			Add(shi.S("y"), &core.Int),
		func (sloc shi.Sloc, stack *shi.Values, vm *shi.VM) error {
			y := shi.Cast(stack.Pop(), &core.Int)
			x := stack.Peek()
			x.Data = shi.Cast(*x, &core.Int) + y
			return nil
		})
}
