package shi

type Label struct {
	Pc PC
}

func NewLabel() *Label {
	return &Label{-1}
}
