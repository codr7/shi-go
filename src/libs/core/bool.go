package core

import (
	"bufio"
	"shi/src/ops"
	"shi/src/shi"
)

type TBool struct {
	shi.BaseType[bool]
}

var Bool TBool
var T shi.Value
var F shi.Value

func init() {
	Bool.Init(shi.S("Bool"))
	T.Init(&Bool, true)
	F.Init(&Bool, false)
}

func (self TBool) Dump(v shi.Value, out *bufio.Writer, vm *shi.VM) error {
	if v.Data.(bool) {
		_, err := out.WriteRune('T')
		return err
	}
	
	_, err := out.WriteRune('F')
	return err
}

func (_ TBool) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	vm.Emit(ops.Push(value))
	return nil
}

func (self TBool) Write(v shi.Value, out *bufio.Writer, vm *shi.VM) error {
	return self.Dump(v, out, vm)
}


