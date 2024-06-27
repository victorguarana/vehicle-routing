package slc

import "log"

type Iterator[T any] interface {
	Actual() T
	ForEach(f func())
	GoToNext()
	GoToPrevious()
	HasNext() bool
	HasPrevious() bool
	Index() int
	Next() T
	Previous() T
	RemoveActualIndex()
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

func (i *iterator[T]) ForEach(f func()) {
	for i.index < len(i.list) {
		f()
		i.index++
	}
}

func (i *iterator[T]) GoToNext() {
	if !i.HasNext() {
		log.Printf("Iterator can not go to next element\n")
		return
	}
	i.index++
}

func (i *iterator[T]) GoToPrevious() {
	if !i.HasPrevious() {
		log.Printf("Iterator can not go to previous element\n")
		return
	}

	i.index--
}

func (i *iterator[T]) HasNext() bool {
	return i.index < len(i.list)-1
}

func (i *iterator[T]) HasPrevious() bool {
	return i.index > 0
}

func (i *iterator[T]) Index() int {
	return i.index
}

func (i *iterator[T]) Next() T {
	if !i.HasNext() {
		log.Panic("Iterator can not get next element\n")
	}
	return i.list[i.index+1]
}

func (i *iterator[T]) Previous() T {
	if !i.HasPrevious() {
		log.Panic("Iterator can not get previous element\n")
	}
	return i.list[i.index-1]
}

func (i *iterator[T]) RemoveActualIndex() {
	i.list = append(i.list[:i.index], i.list[i.index+1:]...)
}
