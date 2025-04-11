package vehicle

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewCarUnlimited", func() {
	It("should create car with correct params", func() {
		initialPoint := gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}
		receivedCar := NewCarUnlimited("car1", initialPoint)
		expectedCar := &CarUnlimited{
			car: &car{
				actualPoint: initialPoint,
				efficiency:  CarDefaultEfficiency,
				drones:      []*drone{},
				name:        "car1",
				speed:       CarDefaultSpeed,
			}}

		Expect(receivedCar).To(Equal(expectedCar))
	})
})

var _ = Describe("CarUnlimited{}", func() {
	Describe("Clone", func() {
		var drone1 = &drone{}
		var sut = &CarUnlimited{
			car: &car{
				actualPoint: gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"},
				drones:      []*drone{drone1},
				efficiency:  10,
				name:        "testCar",
				speed:       20,
			},
		}

		It("should clone car with all params", func() {
			receivedCar := sut.Clone()
			Expect(receivedCar).To(Equal(sut))
			Expect(receivedCar).NotTo(BeIdenticalTo(sut))

			sut.actualPoint.Latitude = 10
			Expect(receivedCar).NotTo(Equal(sut))
		})
	})

	Describe("Move", func() {
		var initialPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}
		var dockedDrone = drone{actualPoint: initialPoint, isFlying: false, remaningRange: DroneRange}
		var flyingDrone = drone{actualPoint: gps.Point{}, isFlying: true, remaningRange: DroneRange}
		var sut = &CarUnlimited{
			car: &car{
				actualPoint: initialPoint,
				drones:      []*drone{&dockedDrone, &flyingDrone},
			},
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
})
