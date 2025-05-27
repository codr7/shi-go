package shi

import (
	"bufio"
)

type Reader interface {
	Read(*VM, *bufio.Reader, *Deque[Form], *Sloc) (bool, error)
}
