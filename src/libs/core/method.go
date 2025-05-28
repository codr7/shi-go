package core

import (
	"shi/src/shi"
)

type TMethod struct {
	shi.BaseType[shi.Method]
}

var Method TMethod

func init() {
	Method.Init(shi.S("Method"))
}
