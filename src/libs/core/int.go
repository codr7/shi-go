package core

import "shi/src/shi"

type TInt struct {
	shi.BaseType[int]
}

var Int TInt

func init() {
	Int.Init(shi.S("Int"))
}
