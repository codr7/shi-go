package core

import (
	"shi/src/ops"
	"shi/src/shi"
)

type TMeta struct {
	shi.BaseType[shi.Type]
}

var Meta TMeta

func init() {
	Meta.Init(shi.S("Meta"))
}

func (_ *TMeta) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	vm.Emit(ops.Push(value))
	return nil
}
