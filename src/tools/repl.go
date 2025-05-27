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
	sloc := shi.NewSloc("repl")

	for {
		if wd, err := os.Getwd(); err != nil {
			log.Fatal(err)
			break
		} else {
			fmt.Fprintf(out, "%v$ ", wd)
			out.Flush()
		}

		MORE:		

		if !in.Scan() {
			if err := in.Err(); err != nil {
				log.Fatal(err)
				break
			}
		}

		fmt.Fprintln(&b, in.Text())
		pc := vm.EmitPC()
		var fs shi.Deque[shi.Form]
		rb := b
		
		if err := vm.ReadAll(bufio.NewReader(&rb), &fs, sloc);
		err != nil {
			goto MORE
		}
		
		b.Reset()

		if err := forms.EmitAll(&fs, vm); err != nil {
			fmt.Println(err)
			goto DONE
		}

		if err := vm.Eval(pc, -1); err != nil {
			fmt.Println(err)
			goto DONE
		}

	DONE:
		shi.DumpStack(vm.Stack, out, vm)
		out.WriteRune('\n')
	}
}
