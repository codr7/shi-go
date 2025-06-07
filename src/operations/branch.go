package operations

import (
	"shi/src/shi"
)

type TBranch struct {
	end *shi.Label
}

func Branch(end *shi.Label) *TBranch {
	return &TBranch{end: end}
}

func (self *TBranch) Compile(vm *shi.VM, pc shi.PC) shi.Eval {
	return func(stack *shi.Values) (shi.PC, error) {
		if stack.Pop().AsBool() {
			return pc + 1, nil
		}

		return self.end.Pc, nil
	}
}
