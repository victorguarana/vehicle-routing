package vehicle

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("car{}", func() {
	Describe("ActualPoint", func() {
		var actualPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}
		var sut = &car{actualPoint: actualPoint}

		It("should return actual point", func() {
			Expect(sut.ActualPoint()).To(BeIdenticalTo(actualPoint))
		})
	})

	Describe("Drones", func() {
		var drone1 = &drone{}
		var drone2 = &drone{}
		var sut = &car{
			drones: []*drone{drone1, drone2},
		}

		It("should return all drones", func() {
			receivedDrones := sut.Drones()
			Expect(receivedDrones).To(HaveLen(2))
			Expect(receivedDrones).To(ContainElements(BeIdenticalTo(drone1), BeIdenticalTo(drone2)))
		})
	})

	Describe("Efficiency", func() {
		var sut = &car{efficiency: 10.0}

		It("should return car efficiency", func() {
			Expect(sut.Efficiency()).To(Equal(10.0))
		})
	})

	Describe("Name", func() {
		var sut = &car{name: "car1"}

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
			sut.NewDefaultDrone("drone1")
			Expect(len(sut.drones)).To(Equal(1))
		})
	})

	Describe("Speed", func() {
		var sut = &car{speed: 10.0}

		It("should return car speed", func() {
			Expect(sut.Speed()).To(Equal(10.0))
		})
	})

	Describe("moveDockedDrones", func() {
		var droneFlying = &drone{isFlying: true}
		var droneDocked = &drone{isFlying: false}
		var sut = &car{drones: []*drone{droneFlying, droneDocked}}

		It("should move only docked drones", func() {
			destination := gps.Point{Latitude: 10}
			sut.moveDockedDrones(destination)

			Expect(droneDocked.actualPoint).To(BeIdenticalTo(destination))
			Expect(droneFlying.actualPoint).NotTo(Equal(destination))
		})
	})
})
