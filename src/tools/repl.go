package tools

import (
	"bufio"
	"bytes"
	//"fmt"
	"os"
	"shi/src/shi"
	"shi/src/forms"
)

func Repl(vm *shi.VM) {
	var t Term
	t.Init(os.Stdin, os.Stdout)
	defer t.Restore()

	var code bytes.Buffer
	sloc := shi.NewSloc("repl")
	var stack shi.Values
	
	for {
		t.Printf("%2v ", sloc.Line()).Flush()
		var lb bytes.Buffer
		i := 1
		
		for {
			in, _ := t.GetChar()

			if len(in) == 1 {
				if c := in[0]; c == ENTER {
					if lb.Len() == 0 {
						break
					} else {
						code.WriteString(lb.String())
						code.WriteRune('\n')
						lb.Reset()
						t.Br().Printf("%2v ", sloc.Line()+i).Flush()
						i++
					}
				} else if c == BACKSPACE {
					if n := lb.Len(); n > 0 {
						lb.Truncate(n-1)
						t.Backspace().Flush()
					}
				} else if c == CTRL_D {
					return
				} else {
					lb.WriteRune(c)
					t.Buffer.WriteRune(c)
					t.Flush()
				}
			} else {
				panic("Not implemented")
			}
		}
		
		pc := vm.EmitPC()
		var fs shi.Forms
		
		if err := vm.ReadAll(bufio.NewReader(&code), &fs, sloc);
		err != nil {
			t.Br().Println(err).Flush()
			code.Reset()
			goto DONE
		}
		
		if err := forms.EmitAll(&fs, vm); err != nil {
			t.Br().Println(err).Flush()
			goto DONE
		}
		
		if err := vm.Eval(pc, -1, &stack); err != nil {
			t.Br().Println(err).Flush()
			goto DONE
		}
		
	DONE:
		t.Br()
		shi.DumpStack(stack, &t.Buffer, vm)
		t.Br().Br().Flush()
	}
}
