package shi

import (
	"bufio"
	"unique"
)

func DumpStack(stack Values, out *bufio.Writer, vm *VM) error {
	out.WriteRune('[')

	for i, v := range stack.Items {
		if i > 0 {
			if _, err := out.WriteRune(' '); err != nil {
				return err
			}
			
		}

		if err := v.Dump(out, vm); err != nil {
			return err
		}
	}
	
	if _, err := out.WriteRune(']'); err != nil {
		return err
	}

	return nil
}

type Sym = unique.Handle[string]

func S(name string) Sym {
	return unique.Make(name)
}
