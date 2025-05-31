package shi

type Label struct {
	PC PC
}

func NewLabel() *Label {
	return &Label{-1}
}
