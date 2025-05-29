package ops

import "shi/src/shi"

type TPut struct {
	target shi.Register
	value  shi.Value
}

func Put(target shi.Register, value shi.Value) *TPut {
	return &TPut{target: target, value: value}
}

func (self *TPut) Compile(vm *shi.VM, pc shi.PC) shi.OpEval {
	return func (stack *shi.Values) (shi.PC, error) {
		v := self.value.Dup(vm)
		vm.Registers.Items[self.target] = &v
		return pc+1, nil
	}
}
