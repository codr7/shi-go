package shi

import (
	"bufio"
)

type PC = int
type Register = int
type Values = Stack[Value]

type VM struct {
	Registers Stack[*Value]

	currentLibrary Library
	userLibrary BaseLibrary

	reader Reader
	operations Stack[Operation]
	code Stack[Eval]
}

func (self *VM) Init(reader Reader) *VM {
	self.userLibrary.Init(S("user"), nil)
	self.currentLibrary = &self.userLibrary
	self.reader = reader
	return self
}

func (self *VM) Allocate(n int) Register {
	result := self.Registers.Len()

	for i := 0; i < n; i++ {
		self.Registers.Push(nil)
	}

	return result
}

func (self *VM) Compile(from PC) {
	for pc := from; pc < self.operations.Len(); pc++ {
		self.code.Push(self.operations.Items[pc].Compile(self, pc))
	}
}

func (self *VM) CurrentLibrary() Library {
	return self.currentLibrary
}

func (self *VM) Emit(operation Operation) {
	self.operations.Push(operation)
}

func (self *VM) EmitPC() PC {
	return self.operations.Len()
}

func (self *VM) Eval(from, to PC, stack *Values) error {
	if to == -1 {
		to = self.operations.Len() - 1
	}

	if self.code.Len() < self.operations.Len() {
		self.Compile(self.code.Len())
	}

	var err error;
	
	for pc := from;
	err == nil && pc <= to;
	pc, err = self.code.Items[pc](stack) {
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

func (self *VM) WithLibrary(lib Library, body func() error) error {
	prev := self.currentLibrary

	if lib == nil {
		lib = new(BaseLibrary)
		lib.Init(self.currentLibrary.Name(), self.currentLibrary) 
	}
	
	self.currentLibrary = lib
	defer func () { self.currentLibrary = prev }()
	return body()
}
