package shi

import (
	"bufio"
)

type PC = int
type Register = int
	
type VM struct {
	Registers Stack[Value]
	Stack Stack[Value]

	currentLib Lib
	userLib BaseLib

	reader Reader
	ops Stack[Op]
	opEvals Stack[OpEval]
}

func (self *VM) Init(reader Reader) *VM {
	self.userLib.Init("user", nil)
	self.currentLib = &self.userLib
	self.reader = reader
	return self
}

func (self *VM) AllocateRegisters(n int) Register {
	result := self.Registers.Len()

	for i := 0; i < n; i++ {
		self.Registers.Push(Value{})
	}

	return result
}

func (self *VM) Compile(from PC) {
	for pc := from; pc < self.ops.Len(); pc++ {
		self.opEvals.Push(self.ops.Items[pc].Compile(self, pc))
	}
}

func (self *VM) CurrentLib() Lib {
	return self.currentLib
}

func (self *VM) Emit(op Op) int {
	result := self.ops.Len()
	self.ops.Push(op)
	return result
}

func (self VM) EmitPC() PC {
	return self.ops.Len()
}

func (self *VM) Eval(from, to PC) error {
	if to == -1 {
		to = self.ops.Len()
	}

	if self.opEvals.Len() < to {
		self.Compile(self.opEvals.Len())
	}

	var err error;
	
	for pc := from;
	err == nil && pc < to;
	pc, err = self.opEvals.Items[pc]() {
		//Do nothing
	}

	return err
}

func (self *VM) ReadAll(in *bufio.Reader, out *Deque[Form], sloc *Sloc) error {
	for {
		ok, err := self.reader.Read(self, in, out, sloc)

		if err != nil {
			return err
		}

		if !ok {
			break;
		}
	}

	return nil
}
