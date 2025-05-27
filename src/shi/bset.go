package shi

import (
	"cmp"
)

type BSetCompare[K cmp.Ordered] = func(x, y K) int
type BSetKey[K cmp.Ordered, V any] = func(V) K

func BSetDefaultKey[K cmp.Ordered, V any](v V) K {
	return any(v).(K)	
}

type BSet[K cmp.Ordered, V any] struct {
	Deque[V]
	compare BSetCompare[K]
	key BSetKey[K, V]
}

func (self *BSet[K, V]) Init(c BSetCompare[K], k BSetKey[K, V]) {
	self.compare = c
	self.key = k

	if self.key == nil {
		self.key = BSetDefaultKey
	}
}

func (self *BSet[K, V]) Add(v V, force bool) bool {
	k := self.key(v)
	i, found := self.IndexOf(k)

	if found == nil {
		self.Insert(i, v)
	} else {
		if !force {
			return false
		}

		self.Items[i] = v
	}

	return true
}

func (self BSet[K, V]) Find(k K) *V {
	_, v := self.IndexOf(k)
	return v
}

func (self BSet[K, V]) IndexOf(k K) (int, *V) {
	min, max := 0, len(self.Items)

	for min < max {
		i := (min + max) / 2
		it := self.Items[i]
		
		switch self.compare(k, self.key(it)) {
		case -1:
			max = i
		case 0:
			return i, &it
		case 1:
			min = i+1
		}
	}

	return min, nil
}

func (self *BSet[K, V]) Remove(k K) *V {
	i, v := self.IndexOf(k)

	if v != nil {
		self.Delete(i, 1)
	}

	return v
}
