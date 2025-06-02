package core

import (
	"shi/src/operations"
	"shi/src/shi"
)

type TInt struct {
	shi.BaseType[int]
}

var Int TInt

func init() {
	Int.Init(shi.S("Int"))
}

func (self *TInt) AsBool(v shi.Value) bool {
	return shi.Cast(v, self) != 0
}

func (_ *TInt) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	vm.Emit(operations.Push(value))
	return nil
}
