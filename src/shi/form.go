package shi

import (
	"bufio"
)

type Forms = Deque[Form]

type Form interface {
	Sloc() Sloc
	Emit(in *Forms, vm *VM) error
	Dump(out *bufio.Writer, vm *VM) error
}

type BaseForm struct {
	sloc Sloc
}

func (self *BaseForm) Init(sloc Sloc) {
	self.sloc = sloc
}

func (self BaseForm) Sloc() Sloc {
	return self.sloc
}
