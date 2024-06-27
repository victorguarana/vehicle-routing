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
		It("should return actual element", func() {
			Expect(sut.Actual()).To(Equal(2))
		})
	})

	Describe("ForEach", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 0,
		}
		var sum = 0
		It("should iterate over elements", func() {
			sut.ForEach(func() {
				sum += sut.Actual()
			})
			Expect(sum).To(Equal(6))
		})
	})

	Describe("GoToNext", func() {
		Context("when there is no next element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 2,
			}
			It("should not increase index", func() {
				sut.GoToNext()
				Expect(sut.index).To(Equal(2))
			})

		})

		Context("when there is next element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 1,
			}
			It("should increase index", func() {
				sut.GoToNext()
				Expect(sut.index).To(Equal(2))
			})
		})
	})

	Describe("GoToPrevious", func() {
		Context("when there is no previous element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 0,
			}
			It("should not decrease index", func() {
				sut.GoToPrevious()
				Expect(sut.index).To(Equal(0))
			})
		})

		Context("when there is previous element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 1,
			}
			It("should decrease index", func() {
				sut.GoToPrevious()
				Expect(sut.index).To(Equal(0))
			})
		})
	})

	Describe("HasNext", func() {
		Context("when there is next element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 1,
			}
			It("should return true", func() {
				Expect(sut.HasNext()).To(BeTrue())
			})
		})

		Context("when there is no next element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 2,
			}
			It("should return false", func() {
				Expect(sut.HasNext()).To(BeFalse())
			})
		})
	})

	Describe("HasPrevious", func() {
		Context("when there is previous element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 1,
			}
			It("should return true", func() {
				Expect(sut.HasPrevious()).To(BeTrue())
			})
		})

		Context("when there is no previous element", func() {
			var sut = &iterator[int]{
				list:  []int{1, 2, 3},
				index: 0,
			}
			It("should return false", func() {
				Expect(sut.HasPrevious()).To(BeFalse())
			})
		})
	})

	Describe("Index", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("should return index", func() {
			Expect(sut.Index()).To(Equal(1))
		})
	})

	Describe("Next", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("should return next element", func() {
			Expect(sut.Next()).To(Equal(3))
		})
	})

	Describe("Previous", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("should return previous element", func() {
			Expect(sut.Previous()).To(Equal(1))
		})
	})

	Describe("RemoveActualIndex", func() {
		var sut = &iterator[int]{
			list:  []int{1, 2, 3},
			index: 1,
		}
		It("should remove actual index and not increment index", func() {
			sut.RemoveActualIndex()
			Expect(sut.list).To(Equal([]int{1, 3}))
			Expect(sut.index).To(Equal(1))
		})
	})
})
