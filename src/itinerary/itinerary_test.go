package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
	mockroute "github.com/victorguarana/vehicle-routing/src/route/mock"
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
		expectedItinerary := &itinerary{
			activeFlights:             map[DroneNumber]route.ISubRoute{},
			droneNumbersMap:           map[DroneNumber]vehicles.IDrone{1: mockedDrone1, 2: mockedDrone2},
			car:                       mockedCar,
			completedSubItineraryList: []subItinerary{},
			route:                     route.NewMainRoute(route.NewMainStop(initialPoint)),
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
				droneNumbersMap: map[DroneNumber]vehicles.IDrone{
					1: mockedDrone1,
					2: mockedDrone2,
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
			droneNumbersMap: map[DroneNumber]vehicles.IDrone{
				1: nil,
				2: nil,
			},
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

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)

			sut = itinerary{
				droneNumbersMap: map[DroneNumber]vehicles.IDrone{
					1: mockedDrone1,
					2: mockedDrone2,
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone is flying", func() {
			mockedDrone1.EXPECT().IsFlying().Return(true)
			Expect(sut.DroneIsFlying(1)).To(BeTrue())
		})

		It("should return false if the drone is not flying", func() {
			mockedDrone2.EXPECT().IsFlying().Return(false)
			Expect(sut.DroneIsFlying(2)).To(BeFalse())
		})
	})

	Describe("DroneSupport", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone *mockvehicles.MockIDrone
		var deliveryPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination1"}
		var landingPoint = gps.Point{Latitude: 7, Longitude: 8, PackageSize: 9, Name: "destination2"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone = mockvehicles.NewMockIDrone(mockedCtrl)

			sut = itinerary{
				droneNumbersMap: map[DroneNumber]vehicles.IDrone{
					1: mockedDrone,
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		Context("when can delivery point and land at the next", func() {
			It("should return true", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(true)
				mockedDrone.EXPECT().CanReach(deliveryPoint, landingPoint).Return(true)
				Expect(sut.DroneSupport(1, deliveryPoint, landingPoint)).To(BeTrue())
			})
		})

		Context("when can delivery point but can not reach at the next", func() {
			It("should return false", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(true)
				mockedDrone.EXPECT().CanReach(deliveryPoint, landingPoint).Return(false)
				Expect(sut.DroneSupport(1, deliveryPoint, landingPoint)).To(BeFalse())
			})
		})

		Context("when can not delivery point", func() {
			It("should return false", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(false)
				Expect(sut.DroneSupport(1, deliveryPoint, landingPoint)).To(BeFalse())
			})
		})
	})

	Describe("StartDroneFlight", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone *mockvehicles.MockIDrone
		var mockedMainStop *mockroute.MockIMainStop
		var mockedSubRoute *mockroute.MockISubRoute

		BeforeEach(func() {
			flightFactory = func(_ route.IMainStop) route.ISubRoute { return mockedSubRoute }
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedMainStop = mockroute.NewMockIMainStop(mockedCtrl)
			mockedSubRoute = mockroute.NewMockISubRoute(mockedCtrl)

			sut = itinerary{
				droneNumbersMap: map[DroneNumber]vehicles.IDrone{
					1: nil,
					2: mockedDrone,
				},
				activeFlights: map[DroneNumber]route.ISubRoute{
					1: nil,
					2: nil,
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		Context("when drone does not have a flight", func() {
			It("should create a new flight", func() {
				mockedDrone.EXPECT().TakeOff()
				sut.StartDroneFlight(2, mockedMainStop)
				Expect(sut.activeFlights[2]).To(Equal(mockedSubRoute))
			})
		})
	})

	Describe("LandDrone", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var mockedSubRoute *mockroute.MockISubRoute
		var mockedMainStop *mockroute.MockIMainStop
		var landingPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedSubRoute = mockroute.NewMockISubRoute(mockedCtrl)
			mockedMainStop = mockroute.NewMockIMainStop(mockedCtrl)

			sut = itinerary{
				activeFlights: map[DroneNumber]route.ISubRoute{
					1: mockedSubRoute,
					2: nil,
				},
				completedSubItineraryList: []subItinerary{},
				droneNumbersMap: map[DroneNumber]vehicles.IDrone{
					1: mockedDrone1,
					2: mockedDrone2,
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
				Expect(sut.activeFlights[1]).To(BeNil())
				Expect(sut.completedSubItineraryList).To(Equal([]subItinerary{{drone: mockedDrone1, flight: mockedSubRoute}}))
			})
		})
	})

	Describe("LandAllDrones", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicles.MockIDrone
		var mockedDrone2 *mockvehicles.MockIDrone
		var mockedDrone3 *mockvehicles.MockIDrone
		var mockedSubRoute1 *mockroute.MockISubRoute
		var mockedSubRoute3 *mockroute.MockISubRoute
		var mockedMainStop *mockroute.MockIMainStop
		var landingPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedDrone3 = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedSubRoute1 = mockroute.NewMockISubRoute(mockedCtrl)
			mockedSubRoute3 = mockroute.NewMockISubRoute(mockedCtrl)
			mockedMainStop = mockroute.NewMockIMainStop(mockedCtrl)

			sut = itinerary{
				activeFlights: map[DroneNumber]route.ISubRoute{
					1: mockedSubRoute1,
					2: nil,
					3: mockedSubRoute3,
				},
				completedSubItineraryList: []subItinerary{},
				droneNumbersMap: map[DroneNumber]vehicles.IDrone{
					1: mockedDrone1,
					2: mockedDrone2,
					3: mockedDrone3,
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones that have flights, remove flights from active flights map and append on completed subItinerary", func() {
			mockedSubRoute1.EXPECT().Return(mockedMainStop)
			mockedSubRoute3.EXPECT().Return(mockedMainStop)
			mockedMainStop.EXPECT().Point().Return(landingPoint).Times(2)
			mockedDrone1.EXPECT().Land(landingPoint)
			mockedDrone3.EXPECT().Land(landingPoint)
			expectedCompletedSubItineraryList := []subItinerary{
				{drone: mockedDrone1, flight: mockedSubRoute1},
				{drone: mockedDrone3, flight: mockedSubRoute3},
			}
			sut.LandAllDrones(mockedMainStop)
			Expect(sut.activeFlights[1]).To(BeNil())
			Expect(sut.activeFlights[2]).To(BeNil())
			Expect(sut.activeFlights[3]).To(BeNil())
			Expect(sut.completedSubItineraryList).To(Equal(expectedCompletedSubItineraryList))
		})
	})

	Describe("MoveCar", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicles.MockICar
		var mockedRoute *mockroute.MockIMainRoute
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicles.NewMockICar(mockedCtrl)
			mockedRoute = mockroute.NewMockIMainRoute(mockedCtrl)

			sut = itinerary{
				car:   mockedCar,
				route: mockedRoute,
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should append stop to route and move car to destination", func() {
			mockedRoute.EXPECT().Append(route.NewMainStop(destination))
			mockedCar.EXPECT().Move(destination)
			sut.MoveCar(destination)
		})
	})

	Describe("MoveDrone", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedDrone *mockvehicles.MockIDrone
		var mockedSubRoute *mockroute.MockISubRoute
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone = mockvehicles.NewMockIDrone(mockedCtrl)
			mockedSubRoute = mockroute.NewMockISubRoute(mockedCtrl)

			sut = itinerary{
				activeFlights: map[DroneNumber]route.ISubRoute{
					1: mockedSubRoute,
					2: nil,
				},
				droneNumbersMap: map[DroneNumber]vehicles.IDrone{
					1: mockedDrone,
					2: nil,
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		Context("when drone has a flight", func() {
			It("should append destination to flight and move drone to destination", func() {
				mockedSubRoute.EXPECT().Append(route.NewSubStop(destination))
				mockedDrone.EXPECT().Move(destination)
				sut.MoveDrone(1, destination)
			})
		})
	})

	Describe("RemoveMainStopFromRoute", func() {
		var sut itinerary
		var mockedCtrl *gomock.Controller
		var mockedRoute *mockroute.MockIMainRoute
		var index = 1

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedRoute = mockroute.NewMockIMainRoute(mockedCtrl)

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
		var mockedRoute *mockroute.MockIMainRoute
		var mockedMainStop1 *mockroute.MockIMainStop
		var mockedMainStop2 *mockroute.MockIMainStop
		var mockedMainStops = []route.IMainStop{mockedMainStop1, mockedMainStop2}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedRoute = mockroute.NewMockIMainRoute(mockedCtrl)

			sut = itinerary{
				route: mockedRoute,
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return an iterator for the route", func() {
			expectedIterator := slc.NewIterator[route.IMainStop](mockedMainStops)
			mockedRoute.EXPECT().Iterator().Return(expectedIterator)
			Expect(sut.RouteIterator()).To(Equal(expectedIterator))
		})
	})
})
