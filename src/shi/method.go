package shi

type MethodArg struct {
	Name Sym
	Type Type
}

type MethodArgs []MethodArg

func (self MethodArgs) Add(name Sym, t Type) MethodArgs {
	return append(self, MethodArg{name, t})
}

type Method interface {
	Args() []MethodArg
	Call(Sloc, PC, *Values, *VM) (PC, error)
	Name() Sym
}

type BaseMethod struct {
	args []MethodArg
	name Sym
}

func (self *BaseMethod) Init(name Sym, args []MethodArg) {
	self.name = name
	self.args = args
}

func (self BaseMethod) Args() []MethodArg {
	return self.args
}

func (self BaseMethod) Name() Sym {
	return self.name
}

type HostMethodBody = func (sloc Sloc, stack *Values, vm *VM) error

type HostMethod struct {
	BaseMethod
	body HostMethodBody
}

func (self *HostMethod) Init(name Sym, args []MethodArg, body HostMethodBody) {
	self.BaseMethod.Init(name, args)
	self.body = body
}

func (self HostMethod) Call(sloc Sloc, pc PC, stack *Values, vm *VM) (PC, error) {
	return pc, self.body(sloc, stack, vm)
}
