package forms

import (
	"shi/src/shi"
)

func EmitAll(in *shi.Deque[shi.Form], vm *shi.VM) error {
	for in.Len() > 0 {
		if err := in.PopFront().Emit(in, vm); err != nil {
			return err
		}
	}

	return nil
}
