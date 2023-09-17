package vehicles

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/go-vehicle-route/gps"
)

var _ = Describe("NewVehicle", func() {
	It("create vehicle with default values", func() {
		p := &gps.Point{}
		v := NewVehicle("test car", p).(*vehicle)

		Expect(v.totalRange).To(Equal(defaultRange))
		Expect(v.remaningRange).To(Equal(defaultRange))
		Expect(v.totalStorage).To(Equal(defaultStorage))
		Expect(v.remaningRange).To(Equal(defaultRange))
		Expect(v.speed).To(Equal(defaultSpeed))
		Expect(v.name).To(Equal("test car"))
		Expect(v.actualPosition).To(Equal(p))
	})
})

var _ = Describe("Move", func() {
	Context("when vehicle can move to next position", func() {
		It("move vehicle", func() {
			initialRange := 100.0
			p := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := vehicle{
				remaningRange: initialRange,
				actualPosition: &gps.Point{
					Latitude:  0,
					Longitude: 0,
				},
			}
			distance := gps.DistanceBetweenPoints(*p, *sut.actualPosition)

			Expect(sut.Move(p)).To(Succeed())
			Expect(sut.actualPosition).To(Equal(p))
			Expect(sut.remaningRange).To(Equal(initialRange - distance))
		})
	})

	Context("when vehicle can not move to next position", func() {
		It("return correct error", func() {
			initialRange := 1.0
			p := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := vehicle{
				remaningRange: initialRange,
				actualPosition: &gps.Point{
					Latitude:  0,
					Longitude: 0,
				},
			}

			Expect(sut.Move(p)).To(MatchError(ErrSoFar))
			Expect(sut.actualPosition).NotTo(Equal(p))
			Expect(sut.remaningRange).To(Equal(initialRange))
		})
	})

	Context("when next position is nil", func() {
		It("raise error", func() {
			sut := vehicle{actualPosition: &gps.Point{}}
			Expect(sut.Move(nil)).Error().To(MatchError(ErrInvalidParams))
		})
	})

	Context("when vehicle does not have position", func() {
		It("raise error", func() {
			sut := vehicle{}
			Expect(sut.Move(&gps.Point{})).Error().To(MatchError(ErrInvalidParams))
		})
	})
})
