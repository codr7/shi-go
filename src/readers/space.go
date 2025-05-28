package readers

import (
	"bufio"
	"shi/src/shi"
	"io"
	"unicode"
)

type TSpace struct {
}

var Space TSpace

func (self TSpace) Read(vm *shi.VM, in *bufio.Reader, out *shi.Forms, sloc *shi.Sloc) (bool, error) {
	found := false
	
	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return false, err
		}

		if unicode.IsSpace(c) {
			sloc.Step(c)
			found = true
		} else {
			in.UnreadRune()
			break
		}
	}

	return found, nil
}
