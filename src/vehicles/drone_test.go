package vehicles

import (
	"github.com/victorguarana/vehicle-routing/src/gps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("newDrone", func() {
	var car = &car{name: "car1"}
	var droneName = "drone1"

	It("should create drone with correct params", func() {
		expectedDrone := drone{
			car:             car,
			efficiency:      DroneEfficiency,
			name:            droneName,
			speed:           DroneSpeed,
			remaningRange:   DroneRange,
			remaningStorage: DroneStorage,
			totalRange:      DroneRange,
			totalStorage:    DroneStorage,
		}
		receivedDrone := newDrone(droneName, car)
		Expect(receivedDrone).To(Equal(&expectedDrone))
	})
})

var _ = Describe("drone{}", func() {
	Describe("ActualPoint", func() {
		var actualPoint = gps.Point{Latitude: 1}
		var sut = drone{
			actualPoint: actualPoint,
		}

		It("should return actual point", func() {
			Expect(sut.ActualPoint()).To(Equal(actualPoint))
		})
	})

	Describe("CanReach", func() {
		var sut = drone{
			remaningRange: 10,
		}
		var initialPoint = gps.Point{Latitude: 0}

		Context("when drone can reach destination", func() {
			It("returns true", func() {
				destination := gps.Point{Latitude: 10}
				Expect(sut.CanReach(initialPoint, destination)).To(BeTrue())
			})
		})

		Context("when drone can not reach destination", func() {
			It("returns false", func() {
				destination := gps.Point{Latitude: 11}
				Expect(sut.CanReach(initialPoint, destination)).To(BeFalse())
			})
		})
	})

	Describe("Land", func() {
		var destination = gps.Point{Latitude: 10}
		var sut = drone{
			totalStorage: 10,
			totalRange:   100,
		}

		It("should land drone and reset attributes", func() {
			sut.Land(destination)
			Expect(sut.remaningRange).To(Equal(sut.totalRange))
			Expect(sut.remaningStorage).To(Equal(DroneStorage))
			Expect(sut.isFlying).To(BeFalse())
			Expect(sut.actualPoint).To(Equal(destination))
		})
	})

	var _ = Describe("Move", func() {
		var initialPoint = gps.Point{Latitude: 5}
		var destinationPoint = gps.Point{Latitude: 10, PackageSize: 1}
		var sut = drone{
			isFlying:        true,
			actualPoint:     initialPoint,
			remaningRange:   DroneRange,
			remaningStorage: DroneStorage,
		}

		It("should move drone and decrease range and package storage", func() {
			distance := gps.DistanceBetweenPoints(initialPoint, destinationPoint)
			sut.Move(destinationPoint)
			Expect(sut.remaningRange).To(Equal(DroneRange - distance))
			Expect(sut.remaningStorage).To(Equal(DroneStorage - destinationPoint.PackageSize))
			Expect(sut.actualPoint).To(Equal(destinationPoint))
		})
	})

	var _ = Describe("Name", func() {
		It("should return drone name", func() {
			name := "drone1"
			sut := drone{
				name: name,
			}
			Expect(sut.Name()).To(Equal(name))
		})
	})

	var _ = Describe("Speed", func() {
		It("should return drone speed", func() {
			speed := 10.0
			sut := drone{
				speed: speed,
			}
			Expect(sut.Speed()).To(Equal(speed))
		})
	})

	var _ = Describe("Support", func() {
		var sut drone
		var initialPoint = gps.Point{Latitude: 0}

		BeforeEach(func() {
			sut = drone{
				remaningRange:   10,
				remaningStorage: 10,
			}
		})

		Describe("single destination cases", func() {
			Context("when drone can support point", func() {
				It("returns true", func() {
					destination := gps.Point{Latitude: 10, PackageSize: 10}
					Expect(sut.Support(initialPoint, destination)).To(BeTrue())
				})
			})

			Context("when drone can not support point because of range", func() {
				It("returns false", func() {
					destination := gps.Point{Latitude: 1}
					sut.remaningRange = 0
					Expect(sut.Support(initialPoint, destination)).To(BeFalse())
				})
			})

			Context("when drone can not support point because of storage", func() {
				It("returns false", func() {
					destination := gps.Point{Latitude: 1, PackageSize: 1}
					sut.remaningStorage = 0
					Expect(sut.Support(initialPoint, destination)).To(BeFalse())
				})
			})
		})

		Describe("multi destinations cases", func() {
			Context("when drone can support points", func() {
				It("returns true", func() {
					destination1 := gps.Point{Latitude: 5}
					destination2 := gps.Point{Latitude: 10}
					Expect(sut.Support(initialPoint, destination1, destination2)).To(BeTrue())
				})
			})

			Context("when drone can not support first point", func() {
				It("returns false", func() {
					destination1 := gps.Point{Latitude: 15}
					destination2 := gps.Point{Latitude: 20}
					Expect(sut.Support(initialPoint, destination1, destination2)).To(BeFalse())
				})
			})

			Context("when drone can not support second point", func() {
				It("returns false", func() {
					destination1 := gps.Point{Latitude: 5}
					destination2 := gps.Point{Latitude: 15}
					Expect(sut.Support(initialPoint, destination1, destination2)).To(BeFalse())
				})
			})
		})
	})

	Describe("TakeOff", func() {
		var sut = drone{
			isFlying: false,
		}

		It("should take off drone", func() {
			sut.TakeOff()
			Expect(sut.isFlying).To(BeTrue())
		})
	})

})
