package core

import (
	"fmt"
	"io"
	"shi/src/operations"
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

func (self *TBool) AsBool(v shi.Value) bool {
	return shi.Cast(v, self)
}

func (self *TBool) Dump(v shi.Value, out io.Writer, vm *shi.VM) error {
	if shi.Cast(v, self) {
		_, err := fmt.Fprint(out, "T")
		return err
	}

	_, err := fmt.Fprint(out, "F")
	return err
}

func (_ *TBool) Emit(value shi.Value, sloc shi.Sloc, in *shi.Forms, vm *shi.VM) error {
	vm.Emit(operations.Push(value))
	return nil
}

func (self *TBool) Write(v shi.Value, out io.Writer, vm *shi.VM) error {
	return self.Dump(v, out, vm)
}
