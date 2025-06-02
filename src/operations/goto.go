package operations

import "shi/src/shi"

type TGoto struct {
	target *shi.Label
}

func Goto(target *shi.Label) *TGoto {
	return &TGoto{target: target}
}

func (self *TGoto) Compile(vm *shi.VM, pc shi.PC) shi.Eval {
	return func(stack *shi.Values) (shi.PC, error) {
		return self.target.PC, nil
	}
}
