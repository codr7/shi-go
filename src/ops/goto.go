package ops

import "shi/src/shi"

type TGoto struct {
	pc shi.Register
}

func Goto(pc shi.PC) *TGoto {
	return &TGoto{pc: pc}
}

func (self *TGoto) Compile(vm *shi.VM, pc shi.PC) shi.OpEval {
	return func() (shi.PC, error) {
		return self.pc, nil
	}
}
