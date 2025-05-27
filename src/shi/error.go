package shi

import (
	"fmt"
)

type BaseError struct {
	message string
}

func (self *BaseError) Init(spec string, args...any) {
	self.message = fmt.Sprintf(spec, args...)
}

func (self BaseError) Error() string {
	return self.message
}

type EmitError struct {
	BaseError
}

func NewEmitError(sloc Sloc, spec string, args...any) EmitError {
	var e EmitError
	e.Init("Emit Error in %v: " + spec, append([]any{sloc}, args...)...)
	return e
}

type ReadError struct {
	BaseError
}

func NewReadError(sloc Sloc, spec string, args...any) ReadError {
	var e ReadError
	e.Init("Read Error in %v: " + spec, append([]any{sloc}, args...)...)
	return e
}
