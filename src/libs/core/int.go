package core

import (
	"shi/src/ops"
	"shi/src/shi"
)

type TInt struct {
	shi.BaseType[int]
}

var Int TInt

func init() {
	Int.Init(shi.S("Int"))
}

func (_ *TInt) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	vm.Emit(ops.Push(value))
	return nil
}
