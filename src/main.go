package main

import (
	"fmt"
	"os"
	"shi/src/shi"
	"shi/src/libs"
	"shi/src/readers"
	"shi/src/tools"
)

func main() {
	var t shi.Term
	t.Init(os.Stdin)
	c, _ := t.GetChar()
	fmt.Printf("%v:%v %v\n", t.Width(), t.Height(), c)
	t.Restore()
		
	var vm shi.VM
	vm.Init(&readers.Form)
	vm.CurrentLib().Import(&libs.Core)
	tools.Repl(&vm)
}
