package main

import (
	"shi/src/shi"
	"shi/src/libs"
	"shi/src/readers"
	"shi/src/tools"
)

func main() {		
	var vm shi.VM
	vm.Init(&readers.Form)
	vm.CurrentLib().Import(&libs.Core)
	tools.Repl(&vm)
}
