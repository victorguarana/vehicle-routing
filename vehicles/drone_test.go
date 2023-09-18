package vehicles

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/victorguarana/go-vehicle-route/gps"
)

var _ = Describe("Move", func() {
	Context("when drone can move to next position", func() {
		It("move drone", func() {
			initialRange := 100.0
			p := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := drone{
				remaningRange: initialRange,
				vehicle: vehicle{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				},
			}
			distance := gps.DistanceBetweenPoints(*p, *sut.actualPosition)

			Expect(sut.Move(p)).To(Succeed())
			Expect(sut.actualPosition).To(Equal(p))
			Expect(sut.remaningRange).To(Equal(initialRange - distance))
		})
	})

	Context("when drone can not move to next position", func() {
		It("return correct error", func() {
			initialRange := 1.0
			p := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := drone{
				remaningRange: initialRange,
				vehicle: vehicle{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				},
			}

			Expect(sut.Move(p)).To(MatchError(ErrWithoutRange))
			Expect(sut.actualPosition).NotTo(Equal(p))
			Expect(sut.remaningRange).To(Equal(initialRange))
		})
	})

	Context("when next position is nil", func() {
		It("raise error", func() {
			sut := drone{
				vehicle: vehicle{actualPosition: &gps.Point{}},
			}
			Expect(sut.Move(nil)).Error().To(MatchError(ErrInvalidParams))
		})
	})

	Context("when drone does not have position", func() {
		It("raise error", func() {
			sut := drone{}
			Expect(sut.Move(&gps.Point{})).Error().To(MatchError(ErrInvalidParams))
		})
	})
})

var _ = Describe("Reachable", func() {
	Context("when drone can reach point with plenty", func() {
		It("returns true", func() {
			sut := drone{
				remaningRange: 100,
				vehicle: vehicle{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				},
			}
			Expect(sut.Reachable(gps.Point{Latitude: 10})).To(BeTrue())
		})
	})

	Context("when drone can reach point without plenty", func() {
		It("returns true", func() {
			sut := drone{
				remaningRange: 10,
				vehicle: vehicle{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				},
			}
			Expect(sut.Reachable(gps.Point{Latitude: 10})).To(BeTrue())
		})
	})

	Context("when drone can not reach point", func() {
		It("returns false", func() {
			sut := drone{
				remaningRange: 0,
				vehicle: vehicle{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				},
			}
			Expect(sut.Reachable(gps.Point{Latitude: 1})).To(BeFalse())
		})
	})
})

var _ = DescribeTable("IsFlying", func(sut drone, expectedResponse bool) {
	Expect(sut.IsFlying()).To(Equal(expectedResponse))
},
	Entry("when drone is flying, returns true",
		drone{
			isFlying: true,
		},
		true,
	),
	Entry("when drone is not flying, returns false",
		drone{
			isFlying: false,
		},
		false,
	),
)
