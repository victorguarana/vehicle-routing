package vehicle

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewCarLimited", func() {
	It("should create car with correct params", func() {
		initialPoint := gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}
		receivedCar := NewCarLimited(
			CarParams{
				Efficiency:    1.0,
				Speed:         5.0,
				Storage:       20.0,
				Range:         100.0,
				Name:          "car1",
				StartingPoint: initialPoint,
			})
		expectedCar := &CarLimited{
			car: &car{
				actualPoint: initialPoint,
				efficiency:  1.0,
				drones:      []*drone{},
				name:        "car1",
				speed:       5.0,
			},
			remaningRange:   100.0,
			remaningStorage: 20.0,
			totalRange:      100.0,
			totalStorage:    20.0,
		}

		Expect(receivedCar).To(Equal(expectedCar))
	})
})

var _ = Describe("CarLimited{}", func() {

	Describe("Clone", func() {
		var drone1 = &drone{}
		var sut = &CarLimited{
			car: &car{
				actualPoint: gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"},
				drones:      []*drone{drone1},
				efficiency:  10,
				name:        "testCar",
				speed:       20,
			},
			remaningRange:   1.0,
			remaningStorage: 2.0,
			totalRange:      10.0,
			totalStorage:    20.0,
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
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 5, Name: "destination"}
		var dockedDrone = drone{actualPoint: initialPoint, isFlying: false, remaningRange: DroneRange}
		var flyingDrone = drone{actualPoint: gps.Point{}, isFlying: true, remaningRange: DroneRange}
		var sut = &CarLimited{
			car: &car{
				actualPoint: initialPoint,
				drones:      []*drone{&dockedDrone, &flyingDrone},
			},
			remaningRange:   10.0,
			remaningStorage: 20.0,
		}

		It("should move car and docked drones without decrease range to destination", func() {
			sut.Move(destination)
			Expect(sut.actualPoint).To(Equal(destination))
			Expect(sut.remaningRange).To(Equal(4.0))
			Expect(sut.remaningStorage).To(Equal(15.0))
			Expect(flyingDrone.actualPoint).NotTo(Equal(destination))
			Expect(flyingDrone.remaningRange).To(Equal(DroneRange))
			Expect(dockedDrone.actualPoint).To(Equal(destination))
			Expect(dockedDrone.remaningRange).To(Equal(DroneRange))
		})
	})

	Describe("Range", func() {
		var sut = &CarLimited{
			totalRange: 100.0,
		}

		It("should return the total range of the car", func() {
			Expect(sut.Range()).To(Equal(100.0))
		})
	})

	Describe("Storage", func() {
		var sut = &CarLimited{
			totalStorage: 50.0,
		}

		It("should return the total storage of the car", func() {
			Expect(sut.Storage()).To(Equal(50.0))
		})
	})

	Describe("Support", func() {
		var sut = &CarLimited{
			remaningRange:   100.0,
			remaningStorage: 50.0,
		}

		Context("when the car can support the destinations", func() {
			var destinations = []gps.Point{
				{Latitude: 1, Longitude: 2, PackageSize: 10},
				{Latitude: 3, Longitude: 4, PackageSize: 20},
			}

			It("should return true", func() {
				Expect(sut.Support(destinations...)).To(BeTrue())
			})
		})

		Context("when the car cannot support the destinations due to range", func() {
			var destinations = []gps.Point{
				{Latitude: 1, Longitude: 2, PackageSize: 10},
				{Latitude: 100, Longitude: 200, PackageSize: 20},
			}

			It("should return false", func() {
				Expect(sut.Support(destinations...)).To(BeFalse())
			})
		})

		Context("when the car cannot support the destinations due to storage", func() {
			var destinations = []gps.Point{
				{Latitude: 1, Longitude: 2, PackageSize: 30},
				{Latitude: 3, Longitude: 4, PackageSize: 40},
			}

			It("should return false", func() {
				Expect(sut.Support(destinations...)).To(BeFalse())
			})
		})
	})
})
