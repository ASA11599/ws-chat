package set

type HashSet[T comparable] struct {
	m map[T]bool
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		m: make(map[T]bool),
	}
}

func (hs *HashSet[T]) Insert(item T) {
	hs.m[item] = true
}

func (hs *HashSet[T]) Contains(item T) bool {
	return hs.m[item]
}

func (hs *HashSet[T]) Delete(item T) {
	delete(hs.m, item)
}

func (hs *HashSet[T]) Items() []T {
	items := make([]T, 0, len(hs.m))
	for k := range hs.m {
		items = append(items, k)
	}
	return items
}
