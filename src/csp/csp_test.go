package csp

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mocks"
	mockroutes "github.com/victorguarana/vehicle-routing/src/routes/mocks"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CoveringWithDrones", func() {
	var mockedCtrl *gomock.Controller
	var mockedItinerary1 *mockitinerary.MockItinerary
	var mockedItinerary2 *mockitinerary.MockItinerary
	var mockedCarStop *mockroutes.MockIMainStop
	var itineraryList []itinerary.Itinerary
	var mockedDroneNumber1 = itinerary.DroneNumber(1)
	var mockedDroneNumber2 = itinerary.DroneNumber(2)
	var mockedDroneNumbers = []itinerary.DroneNumber{mockedDroneNumber1, mockedDroneNumber2}
	var initialPoint = gps.Point{Latitude: 0}
	var client1 = gps.Point{Latitude: 10}
	var client2 = gps.Point{Latitude: 20}
	var client3 = gps.Point{Latitude: 30}
	var client4 = gps.Point{Latitude: 40}
	var client5 = gps.Point{Latitude: 50}
	var client6 = gps.Point{Latitude: 60}
	var client7 = gps.Point{Latitude: 70}
	var clients = []gps.Point{client1, client2, client3, client4, client5, client6, client7}
	var warehouse = gps.Point{Latitude: 0}
	var gpsMap = gps.Map{Clients: clients, Warehouses: []gps.Point{warehouse}}
	var neighborhoodDistance = 10.0

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedItinerary1 = mockitinerary.NewMockItinerary(mockedCtrl)
		mockedItinerary2 = mockitinerary.NewMockItinerary(mockedCtrl)
		mockedCarStop = mockroutes.NewMockIMainStop(mockedCtrl)
		itineraryList = []itinerary.Itinerary{mockedItinerary1, mockedItinerary2}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should cover clients with drones", func() {
		// First iteration
		mockedItinerary1.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedItinerary1.EXPECT().MoveCar(client2)
		mockedItinerary1.EXPECT().DroneNumbers().Return(mockedDroneNumbers)
		mockedItinerary1.EXPECT().ActualCarPoint().Return(client2)
		mockedItinerary1.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedItinerary1.EXPECT().DroneSupport(mockedDroneNumber1, client1, client2).Return(true)
		mockedItinerary1.EXPECT().DroneIsFlying(mockedDroneNumber1).Return(false)
		mockedItinerary1.EXPECT().StartDroneFlight(mockedDroneNumber1, mockedCarStop)
		mockedItinerary1.EXPECT().MoveDrone(mockedDroneNumber1, client1)
		mockedItinerary1.EXPECT().DroneSupport(mockedDroneNumber2, client3, client2).Return(true)
		mockedItinerary1.EXPECT().DroneIsFlying(mockedDroneNumber2).Return(false)
		mockedItinerary1.EXPECT().StartDroneFlight(mockedDroneNumber2, mockedCarStop)
		mockedItinerary1.EXPECT().MoveDrone(mockedDroneNumber2, client3)
		mockedItinerary1.EXPECT().LandAllDrones(mockedCarStop)

		// Second iteration
		mockedItinerary2.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedItinerary2.EXPECT().MoveCar(client5)
		mockedItinerary2.EXPECT().DroneNumbers().Return(mockedDroneNumbers)
		mockedItinerary2.EXPECT().ActualCarPoint().Return(client5)
		mockedItinerary2.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedItinerary2.EXPECT().DroneSupport(mockedDroneNumber1, client4, client5).Return(true)
		mockedItinerary2.EXPECT().DroneIsFlying(mockedDroneNumber1).Return(false)
		mockedItinerary2.EXPECT().StartDroneFlight(mockedDroneNumber1, mockedCarStop)
		mockedItinerary2.EXPECT().MoveDrone(mockedDroneNumber1, client4)
		mockedItinerary2.EXPECT().DroneSupport(mockedDroneNumber2, client6, client5).Return(true)
		mockedItinerary2.EXPECT().DroneIsFlying(mockedDroneNumber2).Return(false)
		mockedItinerary2.EXPECT().StartDroneFlight(mockedDroneNumber2, mockedCarStop)
		mockedItinerary2.EXPECT().MoveDrone(mockedDroneNumber2, client6)
		mockedItinerary2.EXPECT().LandAllDrones(mockedCarStop)

		// Third iteration
		mockedItinerary1.EXPECT().ActualCarPoint().Return(client6)
		mockedItinerary1.EXPECT().MoveCar(client7)

		// Finish routes on closest warehouses
		mockedItinerary1.EXPECT().ActualCarPoint().Return(client7)
		mockedItinerary1.EXPECT().MoveCar(warehouse)
		mockedItinerary2.EXPECT().ActualCarPoint().Return(client6)
		mockedItinerary2.EXPECT().MoveCar(warehouse)

		CoveringWithDrones(itineraryList, gpsMap, neighborhoodDistance)
	})
})

var _ = Describe("deliverNeighborsWithDrones", func() {
	var mockedCtrl *gomock.Controller
	var mockedItinerary *mockitinerary.MockItinerary
	var mockedCarStop *mockroutes.MockIMainStop
	var drone1 = itinerary.DroneNumber(1)
	var drone2 = itinerary.DroneNumber(2)
	var droneNumbers = []itinerary.DroneNumber{drone1, drone2}
	var actualCarPoint = gps.Point{Latitude: 0}
	var client1 = gps.Point{Latitude: 1}
	var client2 = gps.Point{Latitude: 2}
	var client3 = gps.Point{Latitude: 3}
	var client4 = gps.Point{Latitude: 4}
	var neighbors = []gps.Point{client1, client2, client3, client4}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedItinerary = mockitinerary.NewMockItinerary(mockedCtrl)
		mockedCarStop = mockroutes.NewMockIMainStop(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should deliver neighbors with drones", func() {
		mockedItinerary.EXPECT().DroneNumbers().Return(droneNumbers)
		mockedItinerary.EXPECT().ActualCarPoint().Return(actualCarPoint)
		mockedItinerary.EXPECT().ActualCarStop().Return(mockedCarStop)

		// Drone1 supports client1: Start flight and move to client1
		mockedItinerary.EXPECT().DroneSupport(drone1, client1, actualCarPoint).Return(true)
		mockedItinerary.EXPECT().DroneIsFlying(drone1).Return(false)
		mockedItinerary.EXPECT().StartDroneFlight(drone1, mockedCarStop)
		mockedItinerary.EXPECT().MoveDrone(drone1, client1)

		// Drone2 supports client2: Start flight and move to client2
		mockedItinerary.EXPECT().DroneSupport(drone2, client2, actualCarPoint).Return(true)
		mockedItinerary.EXPECT().DroneIsFlying(drone2).Return(false)
		mockedItinerary.EXPECT().StartDroneFlight(drone2, mockedCarStop)
		mockedItinerary.EXPECT().MoveDrone(drone2, client2)

		// Drone1 does not support client3: Land drone1
		mockedItinerary.EXPECT().DroneSupport(drone1, client3, actualCarPoint).Return(false)
		mockedItinerary.EXPECT().LandDrone(drone1, mockedCarStop)

		// Drone2 supports client3: Move to client3
		mockedItinerary.EXPECT().DroneSupport(drone2, client3, actualCarPoint).Return(true)
		mockedItinerary.EXPECT().DroneIsFlying(drone2).Return(true)
		mockedItinerary.EXPECT().MoveDrone(drone2, client3)

		// Drone1 supports client4: Start flight and move to client4
		mockedItinerary.EXPECT().DroneSupport(drone1, client4, actualCarPoint).Return(true)
		mockedItinerary.EXPECT().DroneIsFlying(drone1).Return(false)
		mockedItinerary.EXPECT().StartDroneFlight(drone1, mockedCarStop)
		mockedItinerary.EXPECT().MoveDrone(drone1, client4)

		// Land all drones
		mockedItinerary.EXPECT().LandAllDrones(mockedCarStop)

		deliverNeighborsWithDrones(mockedItinerary, neighbors)
	})
})

var _ = Describe("removeClientAndItsNeighborsFromMap", func() {
	var client0 = gps.Point{Latitude: 0}
	var client1 = gps.Point{Latitude: 10}
	var client2 = gps.Point{Latitude: 20}
	var client3 = gps.Point{Latitude: 30}
	var client4 = gps.Point{Latitude: 40}
	var client5 = gps.Point{Latitude: 50}
	var client6 = gps.Point{Latitude: 60}
	var neighborhood = gps.Neighborhood{
		client0: {client1},
		client1: {client0, client2},
		client2: {client1, client3},
		client3: {client2, client4},
		client4: {client3, client5},
		client5: {client4, client6},
		client6: {client5},
	}

	It("should remove client and its neighbors from map", func() {
		expectedNeighborhood := gps.Neighborhood{
			client0: {client1},
			client1: {client0},
			client5: {client6},
			client6: {client5},
		}
		removeClientAndItsNeighborsFromMap(client3, neighborhood)
		Expect(neighborhood).To(Equal(expectedNeighborhood))
	})
})

var _ = Describe("finishRoutesOnClosestWarehouses", func() {
	var mockCtrl *gomock.Controller
	var mockedItinerary *mockitinerary.MockItinerary
	var itineraryList []itinerary.Itinerary
	var closestWarehouse = gps.Point{Latitude: 1}
	var actualCarPoint = gps.Point{Latitude: 0}
	var gpsMap = gps.Map{Warehouses: []gps.Point{closestWarehouse}}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedItinerary = mockitinerary.NewMockItinerary(mockCtrl)
		itineraryList = []itinerary.Itinerary{mockedItinerary}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest warehouse and append it to the route", func() {
			mockedItinerary.EXPECT().ActualCarPoint().Return(actualCarPoint)
			mockedItinerary.EXPECT().MoveCar(closestWarehouse)
			finishItineraryOnClosestWarehouses(itineraryList, gpsMap)
		})
	})
})
