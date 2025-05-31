package readers

import (
	"bufio"
	//"fmt"
	"shi/src/shi"
	"shi/src/forms"
	"io"
	"strings"
	"unicode"
)

type TId struct {
}

var Id TId

func (_ TId) Read(vm *shi.VM, in *bufio.Reader, out *shi.Forms, sloc *shi.Sloc) (bool, error) {
	formSloc := *sloc
	var buf strings.Builder

	for {
		c, _, err := in.ReadRune()
		
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return false, err
		}

		if unicode.IsSpace(c) ||
			c == '(' || c == ')' {
			in.UnreadRune()
			break
		}

		buf.WriteRune(c)
		sloc.Step(c)
	}

	if buf.Len() == 0 {
		return false, nil
	}
	
	out.PushBack(forms.Id(formSloc, shi.S(buf.String())))
	return true, nil
}
