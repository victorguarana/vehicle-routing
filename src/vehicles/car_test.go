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
				speed:          defaultCarSpeed,
				name:           "car1",
				actualPosition: p,
				drones:         []*drone{},
			}

			Expect(sut).To(Equal(&expectedCar))
		})
	})
})

var _ = Describe("Move", func() {
	Context("when car can move to next position", func() {
		It("move car and docked drones", func() {
			p := &gps.Point{
				Latitude:  10,
				Longitude: 10,
			}

			drone1 := drone{
				isFlying:       true,
				actualPosition: &gps.Point{},
			}

			drone2 := drone{
				isFlying:       false,
				actualPosition: &gps.Point{},
			}

			sut := car{
				drones: []*drone{&drone1, &drone2},
				actualPosition: &gps.Point{
					Latitude:  0,
					Longitude: 0,
				},
			}

			sut.Move(p)
			Expect(sut.actualPosition).To(Equal(p))
			Expect(drone1.actualPosition).To(Equal(&gps.Point{}))
			Expect(drone2.actualPosition).To(Equal(sut.actualPosition))
		})
	})
})

var _ = Describe("Support", func() {
	Describe("single destination", func() {
		Context("when car can reach point with plenty", func() {
			It("returns true", func() {
				destination := gps.Point{Latitude: 10}
				sut := car{
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
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
					actualPosition: &gps.Point{
						Latitude:  0,
						Longitude: 0,
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
				actualPosition: &gps.Point{
					Latitude:  0,
					Longitude: 0,
				},
			}

			expectedDrone := newDrone("drone1", &sut)

			sut.NewDrone("drone1")
			Expect(len(sut.drones)).To(Equal(1))
			Expect(sut.drones).To(Equal([]*drone{expectedDrone}))
		})
	})
})

var _ = Describe("ActualPosition", func() {
	It("returns car position", func() {
		p := &gps.Point{
			Latitude:  10,
			Longitude: 10,
		}
		sut := car{
			actualPosition: p,
		}
		Expect(sut.ActualPosition()).To(Equal(p))
	})
})
