package ops

import (
	"shi/src/shi"
)

type TGet struct {
	source shi.Register
}

func Get(source shi.Register) *TGet {
	return &TGet{source: source}
}

func (self *TGet) Compile(vm *shi.VM, pc shi.PC) shi.OpEval {
	return func (stack *shi.Values) (shi.PC, error) {
		stack.Push(vm.Registers.Items[self.source])
		return pc+1, nil
	}
}
