package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/routes"
	mockroutes "github.com/victorguarana/vehicle-routing/src/routes/mocks"
	"github.com/victorguarana/vehicle-routing/src/slc"
	"github.com/victorguarana/vehicle-routing/src/vehicles"
	mockvehicles "github.com/victorguarana/vehicle-routing/src/vehicles/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("New", func() {
	var mockedCtrl *gomock.Controller
	var mockedCar *mockvehicles.MockICar
	var mockedDrone1 *mockvehicles.MockIDrone
	var mockedDrone2 *mockvehicles.MockIDrone
	var initialPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockedCtrl)
		mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)
	})

	It("should return an itinerary", func() {
		mockedCar.EXPECT().Drones().Return([]vehicles.IDrone{mockedDrone1, mockedDrone2})
		mockedCar.EXPECT().ActualPoint().Return(initialPoint)
		expectedItinerary := itinerary{
			car: mockedCar,
			dronesAndFlights: map[DroneNumber]subItinerary{
				1: {drone: mockedDrone1},
				2: {drone: mockedDrone2},
			},
			route: routes.NewMainRoute(routes.NewMainStop(initialPoint)),
		}
		receivedItinerary := New(mockedCar)
		Expect(receivedItinerary).To(Equal(expectedItinerary))
	})
})

var _ = Describe("itinerary{}", func() {
	Describe("ActualCarPoint", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicles.MockICar
		var initialPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicles.NewMockICar(mockedCtrl)
			sut = itinerary{
				car: mockedCar,
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return the last point of the route", func() {
			mockedCar.EXPECT().ActualPoint().Return(initialPoint)
			Expect(sut.ActualCarPoint()).To(Equal(initialPoint))
		})
	})

	Describe("CarSupport", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicles.MockICar
		var nextPoints = []gps.Point{
			{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination1"},
			{Latitude: 7, Longitude: 8, PackageSize: 9, Name: "destination2"},
		}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicles.NewMockICar(mockedCtrl)

			sut = itinerary{
				car: mockedCar,
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the car supports the route", func() {
			mockedCar.EXPECT().Support(nextPoints).Return(true)
			Expect(sut.CarSupport(nextPoints...)).To(BeTrue())
		})

		It("should return false if the car does not support the route", func() {
			mockedCar.EXPECT().Support(nextPoints).Return(false)
			Expect(sut.CarSupport(nextPoints...)).To(BeFalse())
		})
	})

	Describe("DroneCanReach", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var nextPoints = []gps.Point{
			{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination1"},
			{Latitude: 7, Longitude: 8, PackageSize: 9, Name: "destination2"},
		}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)

			sut = itinerary{
				dronesAndFlights: map[DroneNumber]subItinerary{
					1: {drone: mockedDrone1},
					2: {drone: mockedDrone2},
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone can reach the route", func() {
			mockedDrone1.EXPECT().CanReach(nextPoints).Return(true)
			Expect(sut.DroneCanReach(1, nextPoints...)).To(BeTrue())
		})

		It("should return false if the drone can not reach the route", func() {
			mockedDrone2.EXPECT().CanReach(nextPoints).Return(false)
			Expect(sut.DroneCanReach(2, nextPoints...)).To(BeFalse())
		})
	})

	Describe("DroneNumbers", func() {
		var sut = itinerary{
			dronesAndFlights: map[DroneNumber]subItinerary{1: {}, 2: {}},
		}

		It("should return all drone numbers", func() {
			receivedDroneNumbers := sut.DroneNumbers()
			Expect(receivedDroneNumbers).To(HaveLen(2))
			Expect(receivedDroneNumbers).To(ContainElements(DroneNumber(1), DroneNumber(2)))
		})
	})

	Describe("DroneIsFlying", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var mockedSubRoute *mockroutes.MockISubRoute

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedSubRoute = mockroutes.NewMockISubRoute(mockedCtrl)

			sut = itinerary{
				dronesAndFlights: map[DroneNumber]subItinerary{
					1: {drone: mockedDrone1, flight: mockedSubRoute},
					2: {drone: mockedDrone2},
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone has a flight", func() {
			Expect(sut.DroneIsFlying(1)).To(BeTrue())
		})

		It("should return false if the drone does not have a flight", func() {
			Expect(sut.DroneIsFlying(2)).To(BeFalse())
		})
	})

	Describe("DroneSupport", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var nextPoints = []gps.Point{
			{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination1"},
			{Latitude: 7, Longitude: 8, PackageSize: 9, Name: "destination2"},
		}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)

			sut = itinerary{
				dronesAndFlights: map[DroneNumber]subItinerary{
					1: {drone: mockedDrone1},
					2: {drone: mockedDrone2},
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone supports the route", func() {
			mockedDrone1.EXPECT().Support(nextPoints).Return(true)
			Expect(sut.DroneSupport(1, nextPoints...)).To(BeTrue())
		})

		It("should return false if the drone does not support the route", func() {
			mockedDrone2.EXPECT().Support(nextPoints).Return(false)
			Expect(sut.DroneSupport(2, nextPoints...)).To(BeFalse())
		})
	})

	Describe("LandDrone", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var mockedSubRoute *mockroutes.MockISubRoute
		var mockedMainStop *mockroutes.MockIMainStop
		var landingPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedSubRoute = mockroutes.NewMockISubRoute(mockedCtrl)
			mockedMainStop = mockroutes.NewMockIMainStop(mockedCtrl)

			sut = itinerary{
				dronesAndFlights: map[DroneNumber]subItinerary{
					1: {drone: mockedDrone1, flight: mockedSubRoute},
					2: {drone: mockedDrone2},
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		Context("when drone has a flight", func() {
			It("should land the drone and remove flight from map", func() {
				mockedSubRoute.EXPECT().Return(mockedMainStop)
				mockedMainStop.EXPECT().Point().Return(landingPoint)
				mockedDrone1.EXPECT().Land(landingPoint)
				sut.LandDrone(1, mockedMainStop)
				Expect(sut.dronesAndFlights[1].flight).To(BeNil())
			})
		})

		Context("when drone does not have a flight", func() {
			It("should do nothing", func() {
				sut.LandDrone(2, mockedMainStop)
				Expect(sut.dronesAndFlights[0].flight).To(BeNil())
			})
		})
	})

	Describe("LandAllDrones", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var mockedDrone3 *mockvehicles.MockIDrone
		var mockedSubRoute1 *mockroutes.MockISubRoute
		var mockedSubRoute3 *mockroutes.MockISubRoute
		var mockedMainStop *mockroutes.MockIMainStop
		var landingPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone3 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedSubRoute1 = mockroutes.NewMockISubRoute(mockedCtrl)
			mockedSubRoute3 = mockroutes.NewMockISubRoute(mockedCtrl)
			mockedMainStop = mockroutes.NewMockIMainStop(mockedCtrl)

			sut = itinerary{
				dronesAndFlights: map[DroneNumber]subItinerary{
					1: {drone: mockedDrone1, flight: mockedSubRoute1},
					2: {drone: mockedDrone2},
					3: {drone: mockedDrone3, flight: mockedSubRoute3},
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones that have flights and remove flights from map", func() {
			mockedSubRoute1.EXPECT().Return(mockedMainStop)
			mockedSubRoute3.EXPECT().Return(mockedMainStop)
			mockedMainStop.EXPECT().Point().Return(landingPoint).Times(2)
			mockedDrone1.EXPECT().Land(landingPoint)
			mockedDrone3.EXPECT().Land(landingPoint)
			sut.LandAllDrones(mockedMainStop)
			Expect(sut.dronesAndFlights[1].flight).To(BeNil())
			Expect(sut.dronesAndFlights[2].flight).To(BeNil())
			Expect(sut.dronesAndFlights[3].flight).To(BeNil())
		})
	})

	Describe("MoveCar", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicles.MockICar
		var mockedRoute *mockroutes.MockIMainRoute
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicles.NewMockICar(mockedCtrl)
			mockedRoute = mockroutes.NewMockIMainRoute(mockedCtrl)

			sut = itinerary{
				car:   mockedCar,
				route: mockedRoute,
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should append stop to route and move car to destination", func() {
			mockedRoute.EXPECT().Append(routes.NewMainStop(destination))
			mockedCar.EXPECT().Move(destination)
			sut.MoveCar(destination)
		})
	})

	Describe("MoveDrone", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var mockedRoute *mockroutes.MockIMainRoute
		var mockedMainStop *mockroutes.MockIMainStop
		var mockedSubRoute1 *mockroutes.MockISubRoute
		var mockedSubRoute2 *mockroutes.MockISubRoute
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			flightFactory = func(_ routes.IMainStop) routes.ISubRoute { return mockedSubRoute2 }
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedRoute = mockroutes.NewMockIMainRoute(mockedCtrl)
			mockedMainStop = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedSubRoute1 = mockroutes.NewMockISubRoute(mockedCtrl)
			mockedSubRoute2 = mockroutes.NewMockISubRoute(mockedCtrl)

			sut = itinerary{
				route: mockedRoute,
				dronesAndFlights: map[DroneNumber]subItinerary{
					1: {drone: mockedDrone1, flight: mockedSubRoute1},
					2: {drone: mockedDrone2, flight: nil},
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		Context("when drone has a flight", func() {
			It("should append destination to flight and move drone to destination", func() {
				mockedSubRoute1.EXPECT().Append(routes.NewSubStop(destination))
				mockedDrone1.EXPECT().Move(destination)
				sut.MoveDrone(1, destination)
			})
		})

		Context("when drone does not have a flight", func() {
			It("should create a new flight, append destination to flight and move drone to destination", func() {
				mockedRoute.EXPECT().Last().Return(mockedMainStop)
				mockedSubRoute2.EXPECT().Append(routes.NewSubStop(destination))
				mockedDrone2.EXPECT().Move(destination)
				sut.MoveDrone(2, destination)
				Expect(sut.dronesAndFlights[2].flight).To(Equal(mockedSubRoute2))
			})
		})
	})

	Describe("RemoveMainStopFromRoute", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedRoute *mockroutes.MockIMainRoute
		var index = 1

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedRoute = mockroutes.NewMockIMainRoute(mockedCtrl)

			sut = itinerary{
				route: mockedRoute,
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should remove main stop from route", func() {
			mockedRoute.EXPECT().RemoveMainStop(index)
			sut.RemoveMainStopFromRoute(index)
		})
	})

	Describe("RouteIterator", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedRoute *mockroutes.MockIMainRoute
		var mockedMainStop1 *mockroutes.MockIMainStop
		var mockedMainStop2 *mockroutes.MockIMainStop
		var mockedMainStops = []routes.IMainStop{mockedMainStop1, mockedMainStop2}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedRoute = mockroutes.NewMockIMainRoute(mockedCtrl)

			sut = itinerary{
				route: mockedRoute,
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return an iterator for the route", func() {
			expectedIterator := slc.NewIterator[routes.IMainStop](mockedMainStops)
			mockedRoute.EXPECT().Iterator().Return(expectedIterator)
			Expect(sut.RouteIterator()).To(Equal(expectedIterator))
		})
	})
})
