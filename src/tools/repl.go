package tools

import (
	"bufio"
	"bytes"
	"shi/src/shi"
	"shi/src/forms"
	"fmt"
	"log"
	"os"
)

func Repl(vm *shi.VM) {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var b bytes.Buffer

	for {
		if wd, err := os.Getwd(); err != nil {
			log.Fatal(err)
			break
		} else {
			fmt.Fprintf(out, "%v$ ", wd)
			out.Flush()
		}
				
		if !in.Scan() {
			if err := in.Err(); err != nil {
				log.Fatal(err)
			}

			break
		}

		line := in.Text()

		if line == "" {
			sloc := shi.NewSloc("repl")
			pc := vm.EmitPC()
			var fs shi.Deque[shi.Form]

			if err := vm.ReadAll(bufio.NewReader(&b), &fs, sloc);
			err != nil {
				fmt.Println(err)
				b.Reset()
				goto NEXT
			}

			b.Reset()

			if err := forms.EmitAll(&fs, vm); err != nil {
				fmt.Println(err)
				goto NEXT
			}

			if err := vm.Eval(pc, -1); err != nil {
				fmt.Println(err)
				goto NEXT
			}
		NEXT:
			shi.DumpStack(vm.Stack, out, vm)
			out.WriteRune('\n')
		} else {
			fmt.Fprintln(&b, line)
		}
	}
}
