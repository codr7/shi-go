package core

import "shi/src/shi"

type TNil struct {
	shi.BaseType[any]
}

var Nil TNil
var NIL shi.Value

func init() {
	Nil.Init(shi.S("Nil"))
	NIL.Init(&Nil, nil)
}
