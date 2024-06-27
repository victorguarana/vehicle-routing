package slc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CircularSelection", func() {
	Context("when index is less than the length of the list", func() {
		It("return the element at the index", func() {
			list := []int{1, 2, 3}
			receivedElement := CircularSelection(list, 1)
			Expect(receivedElement).To(Equal(2))
		})
	})

	Context("when index is greater than the length of the list", func() {
		It("return the element at the index modulo the length of the list", func() {
			list := []int{1, 2, 3}
			receivedElement := CircularSelection(list, 4)
			Expect(receivedElement).To(Equal(2))
		})
	})
})

var _ = Describe("CircularSelectionWithIndex", func() {
	Context("when index is less than the length of the list", func() {
		It("return the element at the index and the index", func() {
			list := []int{1, 2, 3}
			receivedElement, receivedIndex := CircularSelectionWithIndex(list, 1)
			Expect(receivedElement).To(Equal(2))
			Expect(receivedIndex).To(Equal(1))
		})
	})

	Context("when index is greater than the length of the list", func() {
		It("return the element at the index modulo the length of the list and the index", func() {
			list := []int{1, 2, 3}
			receivedElement, receivedIndex := CircularSelectionWithIndex(list, 4)
			Expect(receivedElement).To(Equal(2))
			Expect(receivedIndex).To(Equal(1))
		})
	})
})

var _ = Describe("Copy", func() {
	Context("when slice is empty", func() {
		It("return a new empty slice", func() {
			elements := []int{}
			receivedElements := Copy(elements)
			Expect(receivedElements).To(BeEmpty())
			Expect(&receivedElements).NotTo(BeIdenticalTo(&elements))
		})
	})

	Context("when slice is not empty", func() {
		It("return a copy of the slice", func() {
			elements := []int{1, 2, 3}
			receivedElements := Copy(elements)
			Expect(receivedElements).To(Equal(elements))
			Expect(&receivedElements).NotTo(BeIdenticalTo(&elements))
		})
	})
})

var _ = Describe("InsertAt", func() {
	Context("when index is 0", func() {
		var elements = []int{1, 2, 3}

		It("insert the element at the beginning of the slice", func() {
			expectedElements := []int{0, 1, 2, 3}
			receivedElements := InsertAt(elements, 0, 0)
			Expect(receivedElements).To(Equal(expectedElements))
		})
	})

	Context("when index is in the middle", func() {
		var elements = []int{1, 2, 3}

		It("insert the element in the middle of the slice", func() {
			expectedElements := []int{1, 0, 2, 3}
			receivedElements := InsertAt(elements, 0, 1)
			Expect(receivedElements).To(Equal(expectedElements))
		})
	})

	Context("when index is the last element", func() {
		var elements = []int{1, 2, 3}

		It("insert the element at the end of the slice", func() {
			expectedElements := []int{1, 2, 3, 0}
			receivedElements := InsertAt(elements, 0, 3)
			Expect(receivedElements).To(Equal(expectedElements))
		})
	})
})

var _ = Describe("RemoveElement", func() {
	Context("when element exists once in the slice", func() {
		Context("when is the first element", func() {
			var elements = []int{1, 2, 3}

			It("remove the element", func() {
				expectedElements := []int{2, 3}
				receivedElements := RemoveElement(elements, 1)
				Expect(receivedElements).To(Equal(expectedElements))
			})
		})

		Context("when is in the middle", func() {
			var elements = []int{1, 2, 3}

			It("remove the element", func() {
				expectedElements := []int{1, 3}
				receivedElements := RemoveElement(elements, 2)
				Expect(receivedElements).To(Equal(expectedElements))
			})
		})

		Context("when is the last element", func() {
			var elements = []int{1, 2, 3}

			It("remove the element", func() {
				expectedElements := []int{1, 2}
				receivedElements := RemoveElement(elements, 3)
				Expect(receivedElements).To(Equal(expectedElements))
			})
		})
	})

	Context("when element exists more than once in the slice", func() {
		var elements = []int{1, 2, 3, 1}

		It("remove the element", func() {
			expectedElements := []int{2, 3}
			receivedElements := RemoveElement(elements, 1)
			Expect(receivedElements).To(Equal(expectedElements))
		})
	})

	Context("when element does not exist in the slice", func() {
		var elements = []int{1, 2, 3}

		It("return the original slice", func() {
			receivedElements := RemoveElement(elements, 4)
			Expect(receivedElements).To(Equal(elements))
		})
	})
})
