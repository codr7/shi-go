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
	Call(Sloc, PC, *VM) (PC, error)
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

type GoMethodBody = func (sloc Sloc, vm *VM) error

type GoMethod struct {
	BaseMethod
	body GoMethodBody
}

func (self *GoMethod) Init(name Sym, args []MethodArg, body GoMethodBody) {
	self.BaseMethod.Init(name, args)
	self.body = body
}

func (self *GoMethod) Call(sloc Sloc, pc PC, vm *VM) (PC, error) {
	return pc, self.body(sloc, vm)
}
