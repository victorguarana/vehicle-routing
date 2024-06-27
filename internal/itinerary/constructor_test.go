package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("contructor{}", func() {
	Describe("LandDrone", func() {
		var sut constructor
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedSubRoute *mockroute.MockISubRoute
		var mockedMainStop *mockroute.MockIMainStop
		var landingPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedSubRoute = mockroute.NewMockISubRoute(mockedCtrl)
			mockedMainStop = mockroute.NewMockIMainStop(mockedCtrl)

			sut = constructor{
				&info{&itinerary{
					activeFlights: map[DroneNumber]route.ISubRoute{
						1: mockedSubRoute,
						2: nil,
					},
					completedSubItineraryList: []SubItinerary{},
					droneNumbersMap: map[DroneNumber]vehicle.IDrone{
						1: mockedDrone1,
						2: mockedDrone2,
					},
				}}}
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
				Expect(sut.completedSubItineraryList).To(Equal([]SubItinerary{{Drone: mockedDrone1, Flight: mockedSubRoute}}))
			})
		})
	})

	Describe("LandAllDrones", func() {
		var sut constructor
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedDrone3 *mockvehicle.MockIDrone
		var mockedSubRoute1 *mockroute.MockISubRoute
		var mockedSubRoute3 *mockroute.MockISubRoute
		var mockedMainStop *mockroute.MockIMainStop
		var landingPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone3 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedSubRoute1 = mockroute.NewMockISubRoute(mockedCtrl)
			mockedSubRoute3 = mockroute.NewMockISubRoute(mockedCtrl)
			mockedMainStop = mockroute.NewMockIMainStop(mockedCtrl)

			sut = constructor{
				&info{&itinerary{
					activeFlights: map[DroneNumber]route.ISubRoute{
						1: mockedSubRoute1,
						2: nil,
						3: mockedSubRoute3,
					},
					completedSubItineraryList: []SubItinerary{},
					droneNumbersMap: map[DroneNumber]vehicle.IDrone{
						1: mockedDrone1,
						2: mockedDrone2,
						3: mockedDrone3,
					},
				}}}
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
			expectedCompletedSubItineraryList := []SubItinerary{
				{Drone: mockedDrone1, Flight: mockedSubRoute1},
				{Drone: mockedDrone3, Flight: mockedSubRoute3},
			}
			sut.LandAllDrones(mockedMainStop)
			Expect(sut.activeFlights[1]).To(BeNil())
			Expect(sut.activeFlights[2]).To(BeNil())
			Expect(sut.activeFlights[3]).To(BeNil())
			Expect(sut.completedSubItineraryList).To(Equal(expectedCompletedSubItineraryList))
		})
	})

	Describe("MoveCar", func() {
		var sut constructor
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicle.MockICar
		var mockedRoute *mockroute.MockIMainRoute
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedRoute = mockroute.NewMockIMainRoute(mockedCtrl)

			sut = constructor{
				&info{&itinerary{
					car:   mockedCar,
					route: mockedRoute,
				}}}
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
		var sut constructor
		var mockedCtrl *gomock.Controller
		var mockedDrone *mockvehicle.MockIDrone
		var mockedSubRoute *mockroute.MockISubRoute
		var destination = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedSubRoute = mockroute.NewMockISubRoute(mockedCtrl)

			sut = constructor{
				&info{&itinerary{
					activeFlights: map[DroneNumber]route.ISubRoute{
						1: mockedSubRoute,
						2: nil,
					},
					droneNumbersMap: map[DroneNumber]vehicle.IDrone{
						1: mockedDrone,
						2: nil,
					},
				}}}
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

	Describe("StartDroneFlight", func() {
		var sut constructor
		var mockedCtrl *gomock.Controller
		var mockedDrone *mockvehicle.MockIDrone
		var mockedMainStop *mockroute.MockIMainStop
		var mockedSubRoute *mockroute.MockISubRoute

		BeforeEach(func() {
			flightFactory = func(_ route.IMainStop) route.ISubRoute { return mockedSubRoute }
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedMainStop = mockroute.NewMockIMainStop(mockedCtrl)
			mockedSubRoute = mockroute.NewMockISubRoute(mockedCtrl)

			sut = constructor{
				&info{&itinerary{
					droneNumbersMap: map[DroneNumber]vehicle.IDrone{
						1: nil,
						2: mockedDrone,
					},
					activeFlights: map[DroneNumber]route.ISubRoute{
						1: nil,
						2: nil,
					},
				}}}
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

})
