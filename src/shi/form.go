package shi

import (
	"bufio"
)

type Form interface {
	Sloc() Sloc
	Emit(in *Deque[Form], vm *VM) error
	Quote(vm *VM) Value
	Dump(out *bufio.Writer) error
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
