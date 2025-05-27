package shi

import (
	"log"
	"os"
	"unicode/utf8"
	t "golang.org/x/crypto/ssh/terminal"
)

type Term struct {
	fd int
	height int
	state *t.State
	width int
}

func (self *Term) Init(in *os.File) *Term {
	self.fd = int(in.Fd())
	var err error
	
	if self.state, err = t.GetState(self.fd); err != nil {
		log.Fatal(err)
	}
		
	t.MakeRaw(self.fd)

	if self.width, self.height, err = t.GetSize(self.fd); err != nil {
		log.Fatal(err)
	}
	
	return self
}

func (self Term) GetChar() ([]rune, error) {
	bs := make([]byte, 64)
	readBytes, err := os.Stdin.Read(bs)

	if err != nil {
		return nil, err
	}
	
	i := 0
	var out []rune
	
	for i < readBytes {
		c, n := utf8.DecodeRune(bs[i:])
		out = append(out, c)
		i += n
	}

	return out, nil
}

func (self Term) Height() int {
	return self.height
}

func (self *Term) Restore() {	
	t.Restore(self.fd, self.state)
}

func (self Term) Width() int {
	return self.width
}
