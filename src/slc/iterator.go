package slc

type Iterator[T any] interface {
	Actual() T
	Index() int
	Next() T
	Previous() T
	HasNext() bool
	HasPrevious() bool
}

type iterator[T any] struct {
	list  []T
	index int
}

func NewIterator[T any](list []T) Iterator[T] {
	return &iterator[T]{list: list}
}

func (i *iterator[T]) Actual() T {
	return i.list[i.index]
}

func (i *iterator[T]) Next() T {
	i.index++
	return i.list[i.index]
}

func (i *iterator[T]) Previous() T {
	i.index--
	return i.list[i.index]
}

func (i *iterator[T]) Index() int {
	return i.index
}

func (i *iterator[T]) HasNext() bool {
	return i.index < len(i.list)
}

func (i *iterator[T]) HasPrevious() bool {
	return i.index > 0
}
