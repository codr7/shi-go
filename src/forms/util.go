package forms

import (
	"shi/src/shi"
)

func EmitAll(in *shi.Forms, vm *shi.VM) error {
	for in.Len() > 0 {
		if err := in.PopFront().Emit(in, vm); err != nil {
			return err
		}
	}

	return nil
}
