package vehicles

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	mockRoutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("newDrone", func() {
	var car = &car{}
	var droneParams = DroneParams{
		Name: "drone1",
		car:  car,
	}

	It("should create drone with correct params", func() {
		expectedDrone := drone{
			car:             car,
			name:            droneParams.Name,
			speed:           defaultDroneSpeed,
			remaningRange:   defaultDroneRange,
			remaningStorage: defaultDroneStorage,
			totalRange:      defaultDroneRange,
			totalStorage:    defaultDroneStorage,
		}
		receivedDrone := newDrone(droneParams)
		Expect(receivedDrone).To(Equal(&expectedDrone))
	})
})

var _ = Describe("drone{}", func() {
	Describe("Land", func() {
		var sut drone
		var mockCtrl *gomock.Controller
		var mockedCarStop *mockRoutes.MockIMainStop
		var mockedFlight *mockRoutes.MockISubRoute

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockedCarStop = mockRoutes.NewMockIMainStop(mockCtrl)
			mockedFlight = mockRoutes.NewMockISubRoute(mockCtrl)

			sut = drone{
				isFlying:     true,
				totalStorage: 10,
				totalRange:   100,
				flight:       mockedFlight,
			}
		})

		AfterEach(func() {
			defer mockCtrl.Finish()
		})

		It("should land drone and reset attributes", func() {
			mockedFlight.EXPECT().Return(mockedCarStop)
			sut.Land(mockedCarStop)
			Expect(sut.isFlying).To(BeFalse())
			Expect(sut.remaningRange).To(Equal(sut.totalRange))
			Expect(sut.remaningStorage).To(Equal(defaultDroneStorage))
			Expect(sut.flight).To(BeNil())
		})
	})

	var _ = Describe("Move", func() {
		Context("when drone is not flying", func() {
			var sut drone
			var mockCtrl *gomock.Controller
			var mockedCarStop *mockRoutes.MockIMainStop
			var mockedDroneStop *mockRoutes.MockISubStop
			var mockedRoute *mockRoutes.MockIMainRoute
			var mockedFlight *mockRoutes.MockISubRoute
			var mockedFlightFactory func(routes.IMainStop) routes.ISubRoute

			BeforeEach(func() {
				mockCtrl = gomock.NewController(GinkgoT())
				mockedCarStop = mockRoutes.NewMockIMainStop(mockCtrl)
				mockedDroneStop = mockRoutes.NewMockISubStop(mockCtrl)
				mockedRoute = mockRoutes.NewMockIMainRoute(mockCtrl)
				mockedFlight = mockRoutes.NewMockISubRoute(mockCtrl)
				mockedFlightFactory = func(routes.IMainStop) routes.ISubRoute { return mockedFlight }

				c := car{
					route: mockedRoute,
				}
				sut = drone{
					remaningRange: defaultDroneRange,
					isFlying:      false,
					car:           &c,
					flightFactory: mockedFlightFactory,
				}
			})

			It("should create flight and move drone", func() {
				takeoffPoint := gps.Point{Latitude: 5}
				destinationPoint := gps.Point{Latitude: 10}
				distance := gps.DistanceBetweenPoints(takeoffPoint, destinationPoint)

				mockedFlight.EXPECT().Append(mockedDroneStop)
				mockedRoute.EXPECT().Last().Return(mockedCarStop)
				mockedCarStop.EXPECT().Point().Return(takeoffPoint)
				mockedDroneStop.EXPECT().Point().Return(destinationPoint)

				sut.Move(mockedDroneStop)
				Expect(sut.remaningRange).To(Equal(defaultDroneRange - distance))
				Expect(sut.isFlying).To(BeTrue())
				Expect(sut.flight).NotTo(BeNil())
			})
		})

		Context("when drone is flying", func() {
			var sut drone
			var mockCtrl *gomock.Controller
			var mockedDestinationDroneStop *mockRoutes.MockISubStop
			var mockedActualDroneStop *mockRoutes.MockISubStop
			var mockedFlight *mockRoutes.MockISubRoute

			BeforeEach(func() {
				mockCtrl = gomock.NewController(GinkgoT())
				mockedDestinationDroneStop = mockRoutes.NewMockISubStop(mockCtrl)
				mockedActualDroneStop = mockRoutes.NewMockISubStop(mockCtrl)
				mockedFlight = mockRoutes.NewMockISubRoute(mockCtrl)

				sut = drone{
					remaningRange: defaultDroneRange,
					isFlying:      true,
					flight:        mockedFlight,
				}
			})

			It("should append stop to flight and update remaning range", func() {
				destinationPoint := gps.Point{Latitude: 10}
				actualPoint := gps.Point{Latitude: 5}
				distance := gps.DistanceBetweenPoints(actualPoint, destinationPoint)

				mockedFlight.EXPECT().Last().Return(mockedActualDroneStop)
				mockedActualDroneStop.EXPECT().Point().Return(actualPoint)
				mockedDestinationDroneStop.EXPECT().Point().Return(destinationPoint)
				mockedFlight.EXPECT().Append(mockedDestinationDroneStop)

				sut.Move(mockedDestinationDroneStop)
				Expect(sut.remaningRange).To(Equal(defaultDroneRange - distance))
			})
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
		var mockCtrl *gomock.Controller
		var mockedFlight *mockRoutes.MockISubRoute
		var sut drone

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockedFlight = mockRoutes.NewMockISubRoute(mockCtrl)
			mockedDroneStop := mockRoutes.NewMockISubStop(mockCtrl)

			mockedFlight.EXPECT().Last().Return(mockedDroneStop)
			mockedDroneStop.EXPECT().Point().Return(gps.Point{})

			sut = drone{
				remaningRange:   10,
				remaningStorage: 10,
				isFlying:        true,
				flight:          mockedFlight,
			}
		})

		Describe("single destination cases", func() {
			Context("when drone can support point", func() {
				It("returns true", func() {
					destination := gps.Point{Latitude: 10, PackageSize: 10}
					Expect(sut.Support(destination)).To(BeTrue())
				})
			})

			Context("when drone can not support point because of range", func() {
				It("returns false", func() {
					destination := gps.Point{Latitude: 1}
					sut.remaningRange = 0
					Expect(sut.Support(destination)).To(BeFalse())
				})
			})

			Context("when drone can not support point because of storage", func() {
				It("returns false", func() {
					destination := gps.Point{Latitude: 1, PackageSize: 1}
					sut.remaningStorage = 0
					Expect(sut.Support(destination)).To(BeFalse())
				})
			})
		})

		Describe("multi destinations cases", func() {
			Context("when drone can support points", func() {
				It("returns true", func() {
					destination1 := gps.Point{Latitude: 5}
					destination2 := gps.Point{Latitude: 10}
					Expect(sut.Support(destination1, destination2)).To(BeTrue())
				})
			})

			Context("when drone can not support first point", func() {
				It("returns false", func() {
					destination1 := gps.Point{Latitude: 15}
					destination2 := gps.Point{Latitude: 20}
					Expect(sut.Support(destination1, destination2)).To(BeFalse())
				})
			})

			Context("when drone can not support second point", func() {
				It("returns false", func() {
					destination1 := gps.Point{Latitude: 5}
					destination2 := gps.Point{Latitude: 15}
					Expect(sut.Support(destination1, destination2)).To(BeFalse())
				})
			})
		})
	})
})
