package core

import (
	"shi/src/ops"
	"shi/src/shi"
)

type TMethod struct {
	shi.BaseType[shi.Method]
}

var Method TMethod

func init() {
	Method.Init(shi.S("Method"))
}

func (self *TMethod) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	m := shi.Cast(value, self)
	var mas []shi.MethodArg

	switch m.Notation() {
	case shi.Prefix:
		mas = m.Args()
	case shi.Infix:
		mas = m.Args()[1:]
	case shi.Postfix:
		break
	}
	
	var eas shi.Forms

	for range mas {
		if err := in.PopFront().Emit(&eas, vm); err != nil {
			return err
		}
	}

	vm.Emit(ops.Call(sloc, m))
	return nil
}
