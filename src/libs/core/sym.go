package core

import "shi/src/shi"

type TSym struct {
	shi.BaseType[shi.Sym]
}

var Sym TSym

func init() {
	Sym.Init(shi.S("Sym"))
}
