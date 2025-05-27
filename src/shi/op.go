package shi

type OpEval func () (PC, error)

type Op interface {
	Compile(vm *VM, pc PC) OpEval
}
