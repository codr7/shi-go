package libraries

import (
	"shi/src/libraries/core"
	"shi/src/shi"
)

func BindMacro(lib shi.Library, name shi.Symbol, args []string, body shi.HostMacroBody) {
	m := new(shi.HostMacro)
	m.Init(name, args, body)
	lib.Bind(name, shi.V(&core.Macro, shi.Macro(m)))
}

func BindMethod(lib shi.Library, name shi.Symbol, args []shi.MethodArg, body shi.HostMethodBody) {
	m := new(shi.HostMethod)
	m.Init(name, args, body)
	lib.Bind(name, shi.V(&core.Method, shi.Method(m)))
}

func BindType(lib shi.Library, it shi.Type) {
	lib.Bind(it.Name(), shi.V(&core.Meta, it))
}
