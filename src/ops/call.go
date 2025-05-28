package ops

import (
	"shi/src/shi"
)

type TCall struct {
	sloc shi.Sloc
	target shi.Method
}

func Call(sloc shi.Sloc, target shi.Method) *TCall {
	return &TCall{sloc: sloc, target: target}
}

func (self *TCall) Compile(vm *shi.VM, pc shi.PC) shi.OpEval {
	return func () (shi.PC, error) {
		return self.target.Call(self.sloc, pc+1, vm)
	}
}
