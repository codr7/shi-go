package libs

import (
	"shi/src/shi"
	"shi/src/libs/core"
)

type TCore struct {
	shi.BaseLib
}

var Core TCore

func init() {
	Core.Init("core")
}

func (self *TCore) Init(name string) {
	self.BaseLib.Init(name, nil)
	self.BindType(&core.Bool)
	self.BindType(&core.Int)
	self.BindType(&core.Meta)
	self.BindType(&core.Nil)
	self.BindType(&core.Sym)

	self.Bind(shi.S("T"), shi.V(&core.Bool, true))
	self.Bind(shi.S("F"), shi.V(&core.Bool, false))
}

func (self *TCore) BindType(t shi.Type) {
	self.Bind(t.Name(), shi.V(&core.Meta, t))
}
