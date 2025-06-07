package shi

import (
	"iter"
	"slices"
)

type Deque[T any] struct {
	Items []T
}

func NewDeque[T any](items ...T) *Deque[T] {
	return new(Deque[T]).Init(items...)
}

func (self *Deque[T]) Init(items ...T) *Deque[T] {
	self.Items = items
	return self
}

func (self *Deque[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range self.Items {
			if !yield(v) {
				return
			}
		}
	}
}

func (self *Deque[T]) Delete(i, n int) {
	self.Items = slices.Delete(self.Items, i, i+n)
}

func (self *Deque[T]) Insert(i int, items ...T) {
	self.Items = slices.Insert(self.Items, i, items...)
}

func (self *Deque[T]) Len() int {
	return len(self.Items)
}

func (self *Deque[T]) PeekFront() T {
	return self.Items[0]
}

func (self *Deque[T]) PeekBack() *T {
	return &self.Items[len(self.Items)-1]
}

func (self *Deque[T]) PopFront() T {
	it := self.Items[0]
	self.Delete(0, 1)
	return it
}

func (self *Deque[T]) PopBack() T {
	i := len(self.Items) - 1
	it := self.Items[i]
	self.Items = self.Items[:i]
	return it
}

func (self *Deque[T]) PushFront(it T) {
	self.Insert(0, it)
}

func (self *Deque[T]) PushBack(it T) {
	self.Items = append(self.Items, it)
}
