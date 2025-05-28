package shi

type Notation int

const (
	Prefix = Notation(0)
	Infix = Notation(1)
	Postfix = Notation(2)
)

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
	Notation() Notation
}

type BaseMethod struct {
	args []MethodArg
	name Sym
	notation Notation
}

func (self *BaseMethod) Init(name Sym, notation Notation, args []MethodArg) {
	self.name = name
	self.notation = notation
	self.args = args
}

func (self BaseMethod) Args() []MethodArg {
	return self.args
}

func (self BaseMethod) Name() Sym {
	return self.name
}

func (self BaseMethod) Notation() Notation {
	return self.notation
}

type GoMethodBody = func (sloc Sloc, vm *VM) error

type GoMethod struct {
	BaseMethod
	body GoMethodBody
}

func (self *GoMethod) Init(name Sym, notation Notation, args []MethodArg, body GoMethodBody) {
	self.BaseMethod.Init(name, notation, args)
	self.body = body
}

func (self *GoMethod) Call(sloc Sloc, pc PC, vm *VM) (PC, error) {
	return pc, self.body(sloc, vm)
}
