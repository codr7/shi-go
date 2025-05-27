package tests

import (
	"testing"

	"shi/src/shi"
	"shi/src/libs/core"
	"shi/src/ops"
	"shi/src/readers"
)

func newVM() *shi.VM {
	return new(shi.VM).Init(&readers.Form)
}

func TestGet(t *testing.T) {
	vm := newVM()

	r := vm.AllocateRegisters(1)
	vm.Registers.Items[r] = shi.V(&core.Int, 42)

	pc := vm.EmitPC()
	vm.Emit(ops.Get(r))
	vm.Eval(pc, -1)

	if v := vm.Stack.Pop().Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPushValue(t *testing.T) {
	vm := newVM()

	pc := vm.EmitPC()
	vm.Emit(ops.Push(shi.V(&core.Int, 42))) 
	vm.Eval(pc, -1)

	if v := vm.Stack.Pop().Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPutValue(t *testing.T) {
	vm := newVM()

	r := vm.AllocateRegisters(1)
	pc := vm.EmitPC()
	vm.Emit(ops.Put(r, shi.V(&core.Int, 42))) 
	vm.Eval(pc, -1)

	if v := vm.Registers.Items[r].Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}
