package csp

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mock"
	mockroute "github.com/victorguarana/vehicle-routing/src/route/mock"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CoveringWithDrones", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor1 *mockitinerary.MockConstructor
	var mockedConstructor2 *mockitinerary.MockConstructor
	var mockedCarStop *mockroute.MockIMainStop
	var constructorList []itinerary.Constructor
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
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedConstructor2 = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		constructorList = []itinerary.Constructor{mockedConstructor1, mockedConstructor2}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should cover clients with drones", func() {
		// First iteration
		mockedConstructor1.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedConstructor1.EXPECT().MoveCar(client2)
		mockedConstructor1.EXPECT().DroneNumbers().Return(mockedDroneNumbers)
		mockedConstructor1.EXPECT().ActualCarPoint().Return(client2)
		mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedConstructor1.EXPECT().DroneSupport(mockedDroneNumber1, client1, client2).Return(true)
		mockedConstructor1.EXPECT().DroneIsFlying(mockedDroneNumber1).Return(false)
		mockedConstructor1.EXPECT().StartDroneFlight(mockedDroneNumber1, mockedCarStop)
		mockedConstructor1.EXPECT().MoveDrone(mockedDroneNumber1, client1)
		mockedConstructor1.EXPECT().DroneSupport(mockedDroneNumber2, client3, client2).Return(true)
		mockedConstructor1.EXPECT().DroneIsFlying(mockedDroneNumber2).Return(false)
		mockedConstructor1.EXPECT().StartDroneFlight(mockedDroneNumber2, mockedCarStop)
		mockedConstructor1.EXPECT().MoveDrone(mockedDroneNumber2, client3)
		mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop)

		// Second iteration
		mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedConstructor2.EXPECT().MoveCar(client5)
		mockedConstructor2.EXPECT().DroneNumbers().Return(mockedDroneNumbers)
		mockedConstructor2.EXPECT().ActualCarPoint().Return(client5)
		mockedConstructor2.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedConstructor2.EXPECT().DroneSupport(mockedDroneNumber1, client4, client5).Return(true)
		mockedConstructor2.EXPECT().DroneIsFlying(mockedDroneNumber1).Return(false)
		mockedConstructor2.EXPECT().StartDroneFlight(mockedDroneNumber1, mockedCarStop)
		mockedConstructor2.EXPECT().MoveDrone(mockedDroneNumber1, client4)
		mockedConstructor2.EXPECT().DroneSupport(mockedDroneNumber2, client6, client5).Return(true)
		mockedConstructor2.EXPECT().DroneIsFlying(mockedDroneNumber2).Return(false)
		mockedConstructor2.EXPECT().StartDroneFlight(mockedDroneNumber2, mockedCarStop)
		mockedConstructor2.EXPECT().MoveDrone(mockedDroneNumber2, client6)
		mockedConstructor2.EXPECT().LandAllDrones(mockedCarStop)

		// Third iteration
		mockedConstructor1.EXPECT().ActualCarPoint().Return(client6)
		mockedConstructor1.EXPECT().MoveCar(client7)

		// Finish routes on closest warehouses
		mockedConstructor1.EXPECT().ActualCarPoint().Return(client7)
		mockedConstructor1.EXPECT().MoveCar(warehouse)
		mockedConstructor2.EXPECT().ActualCarPoint().Return(client6)
		mockedConstructor2.EXPECT().MoveCar(warehouse)

		CoveringWithDrones(constructorList, gpsMap, neighborhoodDistance)
	})
})

var _ = Describe("deliverNeighborsWithDrones", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor *mockitinerary.MockConstructor
	var mockedCarStop *mockroute.MockIMainStop
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
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should deliver neighbors with drones", func() {
		mockedConstructor.EXPECT().DroneNumbers().Return(droneNumbers)
		mockedConstructor.EXPECT().ActualCarPoint().Return(actualCarPoint)
		mockedConstructor.EXPECT().ActualCarStop().Return(mockedCarStop)

		// Drone1 supports client1: Start flight and move to client1
		mockedConstructor.EXPECT().DroneSupport(drone1, client1, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(drone1).Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(drone1, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(drone1, client1)

		// Drone2 supports client2: Start flight and move to client2
		mockedConstructor.EXPECT().DroneSupport(drone2, client2, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(drone2).Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(drone2, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(drone2, client2)

		// Drone1 does not support client3: Land drone1
		mockedConstructor.EXPECT().DroneSupport(drone1, client3, actualCarPoint).Return(false)
		mockedConstructor.EXPECT().LandDrone(drone1, mockedCarStop)

		// Drone2 supports client3: Move to client3
		mockedConstructor.EXPECT().DroneSupport(drone2, client3, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(drone2).Return(true)
		mockedConstructor.EXPECT().MoveDrone(drone2, client3)

		// Drone1 supports client4: Start flight and move to client4
		mockedConstructor.EXPECT().DroneSupport(drone1, client4, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(drone1).Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(drone1, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(drone1, client4)

		// Land all drones
		mockedConstructor.EXPECT().LandAllDrones(mockedCarStop)

		deliverNeighborsWithDrones(mockedConstructor, neighbors)
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
	var mockedConstructor *mockitinerary.MockConstructor
	var constructorList []itinerary.Constructor
	var closestWarehouse = gps.Point{Latitude: 1}
	var actualCarPoint = gps.Point{Latitude: 0}
	var gpsMap = gps.Map{Warehouses: []gps.Point{closestWarehouse}}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockCtrl)
		constructorList = []itinerary.Constructor{mockedConstructor}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest warehouse and append it to the route", func() {
			mockedConstructor.EXPECT().ActualCarPoint().Return(actualCarPoint)
			mockedConstructor.EXPECT().MoveCar(closestWarehouse)
			finishOnClosestWarehouses(constructorList, gpsMap)
		})
	})
})
