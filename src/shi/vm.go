package shi

import (
	"bufio"
)

type PC = int
type Register = int
type Values = Stack[Value]

type VM struct {
	Registers Stack[*Value]

	currentLib Lib
	userLib BaseLib

	reader Reader
	ops Stack[Op]
	opEvals Stack[OpEval]
}

func (self *VM) Init(reader Reader) *VM {
	self.userLib.Init(S("user"), nil)
	self.currentLib = &self.userLib
	self.reader = reader
	return self
}

func (self *VM) Alloc(n int) Register {
	result := self.Registers.Len()

	for i := 0; i < n; i++ {
		self.Registers.Push(nil)
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

func (self *VM) Emit(op Op) {
	self.ops.Push(op)
}

func (self *VM) EmitPC() PC {
	return self.ops.Len()
}

func (self *VM) Eval(from, to PC, stack *Values) error {
	if to == -1 {
		to = self.ops.Len() - 1
	}

	if self.opEvals.Len() < self.ops.Len() {
		self.Compile(self.opEvals.Len())
	}

	var err error;
	
	for pc := from;
	err == nil && pc <= to;
	pc, err = self.opEvals.Items[pc](stack) {
		//Do nothing
	}

	return err
}

func (self *VM) ReadAll(in *bufio.Reader, out *Forms, sloc *Sloc) error {
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

func (self *VM) WithLib(lib Lib, body func() error) error {
	prev := self.currentLib

	if lib == nil {
		lib = new(BaseLib)
		lib.Init(self.currentLib.Name(), self.currentLib) 
	}
	
	self.currentLib = lib
	defer func () { self.currentLib = prev }()
	return body()
}
