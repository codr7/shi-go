package shi

type MethodArg struct {
	Name Symbol
	Type Type
}

type MethodArgs []MethodArg

func (self MethodArgs) Add(name Symbol, t Type) MethodArgs {
	return append(self, MethodArg{name, t})
}

type Method interface {
	Args() []MethodArg
	Call(Sloc, PC, *Values, *VM) (PC, error)
	Name() Symbol
}

type BaseMethod struct {
	args []MethodArg
	name Symbol
}

func (self *BaseMethod) Init(name Symbol, args []MethodArg) {
	self.name = name
	self.args = args
}

func (self *BaseMethod) Args() []MethodArg {
	return self.args
}

func (self *BaseMethod) Name() Symbol {
	return self.name
}

type HostMethodBody = func(sloc Sloc, stack *Values, vm *VM) error

type HostMethod struct {
	BaseMethod
	body HostMethodBody
}

func (self *HostMethod) Init(name Symbol, args []MethodArg, body HostMethodBody) {
	self.BaseMethod.Init(name, args)
	self.body = body
}

func (self *HostMethod) Call(sloc Sloc, pc PC, stack *Values, vm *VM) (PC, error) {
	return pc, self.body(sloc, stack, vm)
}

type ScriptMethod struct {
	BaseMethod
	startPc PC
}

func (self *ScriptMethod) Init(name Symbol, args []MethodArg, startPc PC) {
	self.BaseMethod.Init(name, args)
	self.startPc = startPc
}

func (self *ScriptMethod) Call(sloc Sloc, pc PC, stack *Values, vm *VM) (PC, error) {
	return self.startPc, nil
}
