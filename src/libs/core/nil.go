package core

import (
	"shi/src/ops"
	"shi/src/shi"
)

type TNil struct {
	shi.BaseType[any]
}

var Nil TNil
var N shi.Value

func init() {
	Nil.Init(shi.S("Nil"))
	N.Init(&Nil, nil)
}

func (_ *TNil) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	vm.Emit(ops.Push(value))
	return nil
}
