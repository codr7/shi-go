package ops

import (
	"shi/src/shi"
)

type TPush struct {
	value  shi.Value
}

func Push(value shi.Value) *TPush {
	return &TPush{value: value}
}

func (self *TPush) Compile(vm *shi.VM, pc shi.PC) shi.OpEval {
	return func () (shi.PC, error) {
		vm.Stack.Push(self.value.Dup(vm))
		return pc+1, nil
	}
}
