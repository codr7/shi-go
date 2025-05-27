package core

import "shi/src/shi"

type TMeta struct {
	shi.BaseType[shi.Type]
}

var Meta TMeta

func init() {
	Meta.Init(shi.S("Meta"))
}
