package core

import (
	"shi/src/operations"
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
	var emitArgs shi.Forms

	for range m.Args() {
		if err := in.PopFront().Emit(&emitArgs, vm); err != nil {
			return err
		}
	}

	vm.Emit(operations.Call(sloc, m))
	return nil
}
