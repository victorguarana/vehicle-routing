package vehicle

import (
	"github.com/victorguarana/vehicle-routing/src/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewCar", func() {
	Context("when car can be created", func() {
		var initialPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}
		It("should create car with correct params", func() {
			receivedCar := NewCar("car1", initialPoint)
			expectedCar := car{
				actualPoint: initialPoint,
				efficiency:  CarEfficiency,
				drones:      []*drone{},
				name:        "car1",
				speed:       CarSpeed,
			}

			Expect(receivedCar).To(Equal(&expectedCar))
		})
	})
})

var _ = Describe("car{}", func() {
	Describe("ActualPoint", func() {
		var actualPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}
		var sut = &car{
			actualPoint: actualPoint,
		}

		It("should return actual point", func() {
			Expect(sut.ActualPoint()).To(Equal(actualPoint))
		})
	})

	Describe("Drones", func() {
		var drone1 = &drone{}
		var drone2 = &drone{}
		var sut = &car{
			drones: []*drone{drone1, drone2},
		}

		It("should return all drones", func() {
			Expect(sut.Drones()).To(Equal([]IDrone{drone1, drone2}))
		})
	})

	Describe("Move", func() {
		var initialPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}
		var dockedDrone = drone{actualPoint: initialPoint, isFlying: false, remaningRange: DroneRange}
		var flyingDrone = drone{actualPoint: gps.Point{}, isFlying: true, remaningRange: DroneRange}
		var sut = &car{
			actualPoint: initialPoint,
			drones:      []*drone{&dockedDrone, &flyingDrone},
		}

		It("should move car and docked drones without decrease range to destination", func() {
			sut.Move(destination)
			Expect(sut.actualPoint).To(Equal(destination))
			Expect(flyingDrone.actualPoint).NotTo(Equal(destination))
			Expect(flyingDrone.remaningRange).To(Equal(DroneRange))
			Expect(dockedDrone.actualPoint).To(Equal(destination))
			Expect(dockedDrone.remaningRange).To(Equal(DroneRange))
		})
	})

	Describe("Name", func() {
		var sut = &car{
			name: "car1",
		}

		It("should return car name", func() {
			Expect(sut.Name()).To(Equal("car1"))
		})
	})

	Describe("NewDrone", func() {
		var sut = &car{
			drones: []*drone{},
			name:   "car1",
		}

		It("should create new drone", func() {
			sut.NewDrone("drone1")
			Expect(len(sut.drones)).To(Equal(1))
		})
	})

	Describe("Speed", func() {
		var sut = &car{
			speed: CarSpeed,
		}

		It("should return car speed", func() {
			Expect(sut.Speed()).To(Equal(CarSpeed))
		})
	})
})
