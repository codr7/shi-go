package shi

type Stack[T any] struct {
	Deque[T]
}

func (self *Stack[T]) Peek() *T {
	return self.PeekBack()
}

func (self *Stack[T]) Pop() T {
	return self.PopBack()
}

func (self *Stack[T]) Push(it T) {
	self.PushBack(it)
}

