package tools

import (
	"bufio"
	"bytes"
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

	for {
		t.Printf("%2v ", sloc.Line()).Flush()
		var line bytes.Buffer
		i := 1
		
		for {
			in, _ := t.GetChar()
			
			if len(in) == 1 {
				c := in[0]

				if c == Enter {
					if line.Len() == 0 {
						break
					} else {
						code.WriteString(line.String())
						code.WriteRune('\n')
						line.Reset()
						t.Br().Printf("%2v ", sloc.Line()+i).Flush()
						i++
					}
				} else if c == CtrlD {
					return
				} else {
					line.WriteRune(c)
					t.Out().WriteRune(c)
					t.Flush()
				}
			} else {
				panic("Not implemented")
			}
		}
		
		pc := vm.EmitPC()
		var fs shi.Deque[shi.Form]
		
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
		
		if err := vm.Eval(pc, -1); err != nil {
			t.Br().Println(err).Flush()
			goto DONE
		}
		
	DONE:
		t.Br()
		shi.DumpStack(vm.Stack, t.Out(), vm)
		t.Br().Br().Flush()
	}
}
