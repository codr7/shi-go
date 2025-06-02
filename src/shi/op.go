package shi

type Eval func (stack *Values) (PC, error)

type Operation interface {
	Compile(vm *VM, pc PC) Eval
}
