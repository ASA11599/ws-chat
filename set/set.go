package set

type Set[T comparable] interface {
	Insert(T)
	Contains(T) bool
	Delete(T)
	Items() []T
}
