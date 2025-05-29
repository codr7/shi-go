package libs

import (
	"shi/src/shi"
	"shi/src/libs/core"
)

func BindMethod(l shi.Lib, name shi.Sym, args []shi.MethodArg, resultType shi.Type, body shi.HostMethodBody) {
	m := new(shi.HostMethod)
	m.Init(name, args, resultType, body)
	l.Bind(name, shi.V(&core.Method, shi.Method(m)))
}

func BindType(l shi.Lib, t shi.Type) {
	l.Bind(t.Name(), shi.V(&core.Meta, t))
}
