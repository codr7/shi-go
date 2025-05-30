package readers

import (
	"bufio"
	"shi/src/shi"
	"io"
)

type TForm struct {
}

var Form TForm

func (self TForm) Read(vm *shi.VM, in *bufio.Reader, out *shi.Forms, sloc *shi.Sloc) (bool, error) {
	if _, err := Space.Read(vm, in, out, sloc); err != nil {
		return false, err
	}

	c, _, err := in.ReadRune()

	if err != nil {
		if err == io.EOF {
			err = nil
		}

		return false, err
	}

	in.UnreadRune()
	
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return Int.Read(vm, in, out, sloc)
	case '(':
		return Scope.Read(vm, in, out, sloc)
	default:
		return Id.Read(vm, in, out, sloc)
	}

	return false, shi.NewReadError(*sloc, "Invalid syntax: %v", c)
}
