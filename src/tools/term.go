package tools

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	t "golang.org/x/crypto/ssh/terminal"
)

const (
	BACKSPACE = rune(127)
	CTRL_D = rune(4)
	ENTER = rune(13)
)

type Term struct {
	fd int
	height int
	inFile *os.File
	lineBreak string
	outFile *os.File
	state *t.State
	width int
	Buffer bytes.Buffer
}

func (self *Term) Init(in *os.File, out *os.File) *Term {
	self.inFile = in
	self.outFile = out
	self.fd = int(in.Fd())
	var err error
	
	if self.state, err = t.GetState(self.fd); err != nil {
		log.Fatal(err)
	}
		
	t.MakeRaw(self.fd)

	if self.width, self.height, err = t.GetSize(self.fd); err != nil {
		log.Fatal(err)
	}

	self.lineBreak = "\r\n"
	return self
}

func (self *Term) Backspace() *Term {
	self.Buffer.WriteString("\b \b")
	return self
}

func (self *Term) Br() *Term {
	self.Buffer.WriteString(self.lineBreak)
	return self
}

func (self *Term) Ctrl(args...any) {
        self.Buffer.WriteRune(rune(27));
        self.Buffer.WriteRune('[');

	for _, a := range args {
		fmt.Fprint(&self.Buffer, a)
	}
}

func (self *Term) Flush() *Term {
	self.outFile.WriteString(self.Buffer.String())
	self.Buffer.Reset()
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

func (self *Term) Printf(spec string, values...any) *Term {
	fmt.Fprintf(&self.Buffer, spec, values...)
	return self
}

func (self *Term) Println(values...any) *Term {
	fmt.Fprintln(&self.Buffer, values...)
	return self
}

func (self *Term) Restore() {	
	t.Restore(self.fd, self.state)
}

func (self Term) Width() int {
	return self.width
}
