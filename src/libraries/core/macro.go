package core

import (
	"shi/src/shi"
)

type TMacro struct {
	shi.BaseType[shi.Macro]
}

var Macro TMacro

func init() {
	Macro.Init(shi.S("Macro"))
}

func (self *TMacro) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	return shi.Cast(value, self).Call(sloc, in, vm)
}
