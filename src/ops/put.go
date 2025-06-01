package ops

import "shi/src/shi"

type TPut struct {
	target shi.Register
}

func Put(target shi.Register) *TPut {
	return &TPut{target: target}
}

func (self *TPut) Compile(vm *shi.VM, pc shi.PC) shi.OpEval {
	return func (stack *shi.Values) (shi.PC, error) {
		v := stack.Pop()
		vm.Registers.Items[self.target] = &v
		return pc+1, nil
	}
}
