package slc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Iterator", func() {
	Describe("Actual", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("returns actual element", func() {
			Expect(sut.Actual()).To(Equal(2))
		})
	})

	Describe("Next", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("returns next element and increase index", func() {
			Expect(sut.Next()).To(Equal(3))
			Expect(sut.index).To(Equal(2))
		})
	})

	Describe("Previous", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("returns previous element and decrease index", func() {
			Expect(sut.Previous()).To(Equal(1))
			Expect(sut.index).To(Equal(0))
		})
	})

	Describe("Index", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("returns true", func() {
			Expect(sut.Index()).To(Equal(1))
		})
	})

	Describe("HasNext", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("returns true", func() {
			Expect(sut.HasNext()).To(BeTrue())
		})
	})

	Describe("HasPrevious", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("returns true", func() {
			Expect(sut.HasPrevious()).To(BeTrue())
		})
	})
})
