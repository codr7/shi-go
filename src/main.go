package main

import (
	"shi/src/shi"
	"shi/src/libraries"
	"shi/src/readers"
	"shi/src/tools"
)

func main() {		
	var vm shi.VM
	vm.Init(&readers.Form)
	vm.CurrentLibrary().ImportFrom(&libraries.Core)
	tools.Repl(&vm)
}
