package shi

import (
	"bufio"
)

type Reader interface {
	Read(*VM, *bufio.Reader, *Forms, *Sloc) (bool, error)
}
