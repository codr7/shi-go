package shi

type OpEval func (stack *Values) (PC, error)

type Op interface {
	Compile(vm *VM, pc PC) OpEval
}
