package core

import (
	"bufio"
	"shi/src/shi"
)

type TBool struct {
	shi.BaseType[bool]
}

var Bool TBool

func init() {
	Bool.Init(shi.S("Bool"))
}

func (self TBool) Dump(v shi.Value, out *bufio.Writer, vm *shi.VM) error {
	if v.Data.(bool) {
		_, err := out.WriteRune('T')
		return err
	}
	
	_, err := out.WriteRune('F')
	return err
}

func (self TBool) Write(v shi.Value, out *bufio.Writer, vm *shi.VM) error {
	return self.Dump(v, out, vm)
}


