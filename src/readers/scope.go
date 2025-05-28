package readers

import (
	"bufio"
	//"fmt"
	"shi/src/shi"
	"shi/src/forms"
	"io"
)

type TScope struct {
}

var Scope TScope

func (self TScope) Read(vm *shi.VM, in *bufio.Reader, out *shi.Forms, sloc *shi.Sloc) (bool, error) {
	formSloc := *sloc
	
	if c, _, err := in.ReadRune(); err != nil {
		if err == io.EOF {
			err = nil
		}

		return false, err
	} else if c != '(' {
		in.UnreadRune()
		return false, nil
	} else {
		sloc.Step(c)
	}

	var buf shi.Forms
	
	for {
		if _, err := Space.Read(vm, in, &buf, sloc); err != nil {
			return false, err
		}
		
		if c, _, err := in.ReadRune(); err != nil {
			if err == io.EOF {
				err = shi.NewReadError(*sloc, "Invalid syntax")
			}
			
			return false, err
		} else if c == ')' {
			sloc.Step(c)
			break
		}

		in.UnreadRune()

		if ok, err := Form.Read(vm, in, &buf, sloc); err != nil {
			return false, err
		} else if !ok {
			err = shi.NewReadError(*sloc, "Invalid syntax")
		}
	}

	out.PushBack(forms.NewScope(formSloc, buf))
	return true, nil
}
