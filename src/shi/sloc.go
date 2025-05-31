package shi

import (
	"fmt"
)

type Sloc struct {
	source       string
	line, column int
}

func NewSloc(source string) *Sloc {
	return new(Sloc).Init(source, 1, 0)
}

func (self *Sloc) Init(source string, line, column int) *Sloc {
	self.source = source
	self.line = line
	self.column = column
	return self
}

func (self *Sloc) Line() int {
	return self.line
}

func (self *Sloc) Step(c rune) {
	if c == '\n' {
		self.line++;
		self.column = 0;
	} else {
		self.column++
	}
}

func (self Sloc) String() string {
	return fmt.Sprintf("'%v' at line %v, column %v",
		self.source, self.line, self.column)
}
