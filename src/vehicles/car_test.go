package vehicles

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/victorguarana/go-vehicle-route/src/gps"
)

var _ = Describe("NewCar", func() {
	Context("when car can be created", func() {
		It("create with correct params", func() {
			p := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := NewCar("car1", p)

			expectedCar := car{
				vehicle: vehicle{
					speed:          defaultSpeed,
					name:           "car1",
					actualPosition: p,
				},
				drones: []IDrone{},
			}

			Expect(sut).To(Equal(&expectedCar))
		})
	})
})

var _ = Describe("Move", func() {
	Context("when car can move to next position", func() {
		It("move car", func() {
			p := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}
			sut := car{
				vehicle: vehicle{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				},
			}

			Expect(sut.Move(p)).To(Succeed())
			Expect(sut.actualPosition).To(Equal(p))
		})
	})

	Context("when next position is nil", func() {
		It("raise error", func() {
			sut := car{
				vehicle: vehicle{actualPosition: &gps.Point{}},
			}
			Expect(sut.Move(nil)).Error().To(MatchError(ErrInvalidParams))
		})
	})

	Context("when car does not have position", func() {
		It("raise error", func() {
			sut := car{}
			Expect(sut.Move(&gps.Point{})).Error().To(MatchError(ErrInvalidParams))
		})
	})
})

var _ = Describe("Support", func() {
	Describe("single destination", func() {
		Context("when car can reach point with plenty", func() {
			It("returns true", func() {
				destination := gps.Point{Latitude: 10}
				sut := car{
					vehicle: vehicle{
						actualPosition: &gps.Point{
							Latitude:  0,
							Longitude: 0,
						},
					},
				}
				Expect(sut.Support(&destination)).To(BeTrue())
			})
		})
	})

	Describe("multi destination", func() {
		Context("when car can reach point with plenty", func() {
			It("returns true", func() {
				destination1 := gps.Point{Latitude: 10}
				destination2 := gps.Point{Latitude: 20}
				sut := car{
					vehicle: vehicle{
						actualPosition: &gps.Point{
							Latitude:  0,
							Longitude: 0,
						},
					},
				}
				Expect(sut.Support(&destination1, &destination2)).To(BeTrue())
			})
		})
	})
})

var _ = Describe("NewDrone", func() {
	Context("when car can create a new drone", func() {
		It("create with correct params", func() {
			sut := car{
				vehicle: vehicle{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
					},
				},
			}

			expectedDrone := drone{
				totalStorage:    defaultStorage,
				remaningStorage: defaultStorage,
				totalRange:      defaultRange,
				remaningRange:   defaultRange,
				vehicle: vehicle{
					speed:          defaultSpeed,
					name:           "drone1",
					actualPosition: sut.actualPosition,
				},
				car: &sut,
			}

			sut.NewDrone("drone1")
			Expect(len(sut.drones)).To(Equal(1))
			Expect(sut.drones[0]).To(Equal(expectedDrone))
		})
	})
})
