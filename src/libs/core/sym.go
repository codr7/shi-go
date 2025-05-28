package core

import (
	"shi/src/ops"
	"shi/src/shi"
)

type TSym struct {
	shi.BaseType[shi.Sym]
}

var Sym TSym

func init() {
	Sym.Init(shi.S("Sym"))
}

func (_ TSym) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	vm.Emit(ops.Push(value))
	return nil

}

