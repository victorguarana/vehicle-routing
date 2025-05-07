// This package contains functions to manipulate slices.
package slc

func CircularSelection[T any](list []T, index int) T {
	i := index % len(list)
	return list[i]
}

func CircularSelectionWithIndex[T any](list []T, index int) (T, int) {
	i := index % len(list)
	return list[i], i
}

func Copy[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func InsertAt[T any](slice []T, element T, index int) []T {
	newSlice := make([]T, len(slice)+1)
	for i := 0; i < len(newSlice); i++ {
		switch {
		case i < index:
			newSlice[i] = slice[i]
		case i == index:
			newSlice[i] = element
		case i > index:
			newSlice[i] = slice[i-1]
		}
	}
	return newSlice
}

func RemoveElement[T comparable](slice []T, element T) []T {
	newSlice := make([]T, 0)
	for _, e := range slice {
		if e != element {
			newSlice = append(newSlice, e)
		}
	}
	return newSlice
}

func AppendIfNotExists[T comparable](slice []T, element T) []T {
	for _, i := range slice {
		if i == element {
			return slice
		}
	}
	return append(slice, element)
}
