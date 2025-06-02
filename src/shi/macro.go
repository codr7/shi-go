package shi

type Macro interface {
	Args() []string
	Call(Sloc, *Forms, *VM) error
	Name() Symbol
}

type BaseMacro struct {
	args []string
	name Symbol
}

func (self *BaseMacro) Init(name Symbol, args []string) {
	self.name = name
	self.args = args
}

func (self *BaseMacro) Args() []string {
	return self.args
}

func (self *BaseMacro) Name() Symbol {
	return self.name
}

type HostMacroBody = func (sloc Sloc, in *Forms, vm *VM) error

type HostMacro struct {
	BaseMacro
	body HostMacroBody
}

func (self *HostMacro) Init(name Symbol, args []string, body HostMacroBody) {
	self.BaseMacro.Init(name, args)
	self.body = body
}

func (self *HostMacro) Call(sloc Sloc, in *Forms, vm *VM) error {
	if in.Len() < len(self.args) {
		return NewEmitError(sloc, "Not enough args")
	}
	
	return self.body(sloc, in, vm)
}
