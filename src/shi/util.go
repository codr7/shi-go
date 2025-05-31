package shi

import (
	"fmt"
	"io"
	"unique"
)

func DumpStack(stack Values, out io.Writer, vm *VM) error {
	fmt.Fprint(out, "[")

	for i, v := range stack.Items {
		if i > 0 {
			if _, err := fmt.Fprint(out, " "); err != nil {
				return err
			}
			
		}

		if err := v.Dump(out, vm); err != nil {
			return err
		}
	}
	
	if _, err := fmt.Fprint(out, "]"); err != nil {
		return err
	}

	return nil
}

type Sym = unique.Handle[string]

func S(name string) Sym {
	return unique.Make(name)
}
