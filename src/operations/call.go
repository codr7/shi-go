package operations

import (
	"shi/src/shi"
)

type TCall struct {
	sloc shi.Sloc
	target shi.Method
}

func Call(sloc shi.Sloc, target shi.Method) *TCall {
	return &TCall{sloc: sloc, target: target}
}

func (self *TCall) Compile(vm *shi.VM, pc shi.PC) shi.Eval {
	return func (stack *shi.Values) (shi.PC, error) {
		as := self.target.Args()
		al := len(as)
		sl := stack.Len()
		
		if sl < al {
			return -1, shi.NewEvalError(self.sloc, "Not enough args")
		}
		
		for i := 0; i < al; i++ {
			at := as[i].Type
			v := stack.Items[sl-i-1]
			
			if !v.Isa(at) {
				return -1, shi.NewEvalError(self.sloc, "Type mismatch: expected %v, actual %v", at, v.Type)
			}
		}
		
		return self.target.Call(self.sloc, pc+1, stack, vm)
	}
}
