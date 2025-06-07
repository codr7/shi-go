package main

import (
	"shi/src/libraries"
	"shi/src/readers"
	"shi/src/shi"
	"shi/src/tools"
)

func main() {
	var vm shi.VM
	vm.Init(&readers.Form)
	vm.CurrentLibrary().ImportFrom(&libraries.Core)
	tools.Repl(&vm)
}
