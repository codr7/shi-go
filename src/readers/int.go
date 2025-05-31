package readers

import (
	"bufio"
	"io"
	"shi/src/forms"
	"shi/src/libs/core"
	"shi/src/shi"
	"unicode"
)

type TInt struct {
}

var Int TInt

func (self TInt) Read(vm *shi.VM, in *bufio.Reader, out *shi.Forms, sloc *shi.Sloc) (bool, error) {
	formSloc := *sloc
	var v int

	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return false, err
		}

		if !unicode.IsDigit(c) {
			in.UnreadRune()
			break
		}

		v = v * 10 + int(c) - int('0')
		sloc.Step(c)
	}

	out.PushBack(forms.Literal(formSloc, shi.V(&core.Int, v)))
	return true, nil
}
