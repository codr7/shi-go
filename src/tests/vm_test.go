package tests

import (
	"testing"

	"shi/src/shi"
	"shi/src/libraries/core"
	"shi/src/operations"
	"shi/src/readers"
)

func newVM() *shi.VM {
	return new(shi.VM).Init(&readers.Form)
}

func TestGet(t *testing.T) {
	vm := newVM()

	r := vm.Allocate(1)
	v := shi.V(&core.Int, 42)
	vm.Registers.Items[r] = &v

	pc := vm.EmitPc()
	vm.Emit(operations.Get(r))
	var stack shi.Values
	vm.Eval(pc, -1, &stack)

	if v := stack.Pop().Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPushValue(t *testing.T) {
	vm := newVM()

	pc := vm.EmitPc()
	vm.Emit(operations.Push(shi.V(&core.Int, 42)))
	var stack shi.Values
	vm.Eval(pc, -1, &stack)

	if v := stack.Pop().Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}

func TestPutValue(t *testing.T) {
	vm := newVM()

	var stack shi.Values
	stack.Push(shi.V(&core.Int, 42))
	r := vm.Allocate(1)
	pc := vm.EmitPc()
	vm.Emit(operations.Put(r))
	vm.Eval(pc, -1, &stack)

	if v := vm.Registers.Items[r].Data; v != 42 {
		t.Fatalf("Expected 42, actual %v", v)
	}
}
