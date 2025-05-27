package tools

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"shi/src/shi"
	"shi/src/forms"
)

func Repl(vm *shi.VM) {
	var t shi.Term
	t.Init(os.Stdin, os.Stdout)
	defer t.Restore()
	shellMode := false

	var code bytes.Buffer
	sloc := shi.NewSloc("repl")

	for {
		if shellMode {
			if wd, err := os.Getwd(); err != nil {
				log.Fatal(err)
				break
			} else {
				t.Printf("%v$ ", wd).Flush()
			}
		} else {
			t.Printf("%2v ", sloc.Line()).Flush()
			var line bytes.Buffer
			i := 1
			
			for {
				in, _ := t.GetChar()

				if len(in) == 1 {
					c := in[0]
					
					if c == 13 {
						if line.Len() == 0 {
							break
						} else {
							code.WriteString(line.String())
							code.WriteRune('\n')
							line.Reset()
							t.Printf("\r\n%2v ", sloc.Line()+i).Flush()
							i++
						}
					} else if c == 4 {
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
}
