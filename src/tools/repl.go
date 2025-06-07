package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"shi/src/shi"
	"shi/src/forms"
)

func Repl(vm *shi.VM) {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var code bytes.Buffer
	sloc := shi.NewSloc("repl")
	var stack shi.Values
	
	for {
		fmt.Fprintf(out, "%2v ", sloc.Line())
		out.Flush()
		
		if !in.Scan() {
			if err := in.Err(); err != nil {
				log.Fatal(err)
			}

			break
		}

		line := in.Text()
		
		if line == "" {
			pc := vm.EmitPc()
			var fs shi.Forms
			
			if err := vm.ReadAll(bufio.NewReader(&code), &fs, sloc);
			err != nil {
				fmt.Fprintln(out, err)
				code.Reset()
				goto NEXT
			}
			
			if err := forms.EmitAll(&fs, vm); err != nil {
				fmt.Fprintln(out, err)
				goto NEXT
			}
			
			if err := vm.Eval(pc, -1, &stack); err != nil {
				fmt.Fprintln(out, err)
			}
		NEXT:
			shi.DumpStack(stack, out, vm)
			out.WriteString("\n\n")
			out.Flush()
		} else {
			fmt.Fprintln(&code, line)
		}
	}
}
