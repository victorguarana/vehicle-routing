package csp

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"
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
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedDrones []vehicle.IDrone
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
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrones = []vehicle.IDrone{mockedDrone1, mockedDrone2}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should cover clients with drones", func() {
		// First iteration
		mockedConstructor1.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedConstructor1.EXPECT().MoveCar(client2)
		mockedConstructor1.EXPECT().Drones().Return(mockedDrones)
		mockedConstructor1.EXPECT().ActualCarPoint().Return(client2)
		mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedConstructor1.EXPECT().DroneSupport(mockedDrone1, client1, client2).Return(true)
		mockedConstructor1.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
		mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, client1)
		mockedConstructor1.EXPECT().DroneSupport(mockedDrone2, client3, client2).Return(true)
		mockedConstructor1.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
		mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone2, mockedCarStop)
		mockedConstructor1.EXPECT().MoveDrone(mockedDrone2, client3)
		mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop)

		// Second iteration
		mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedConstructor2.EXPECT().MoveCar(client5)
		mockedConstructor2.EXPECT().Drones().Return(mockedDrones)
		mockedConstructor2.EXPECT().ActualCarPoint().Return(client5)
		mockedConstructor2.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedConstructor2.EXPECT().DroneSupport(mockedDrone1, client4, client5).Return(true)
		mockedConstructor2.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
		mockedConstructor2.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor2.EXPECT().MoveDrone(mockedDrone1, client4)
		mockedConstructor2.EXPECT().DroneSupport(mockedDrone2, client6, client5).Return(true)
		mockedConstructor2.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
		mockedConstructor2.EXPECT().StartDroneFlight(mockedDrone2, mockedCarStop)
		mockedConstructor2.EXPECT().MoveDrone(mockedDrone2, client6)
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
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedDrones []vehicle.IDrone
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
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrones = []vehicle.IDrone{mockedDrone1, mockedDrone2}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should deliver neighbors with drones", func() {
		mockedConstructor.EXPECT().Drones().Return(mockedDrones)
		mockedConstructor.EXPECT().ActualCarPoint().Return(actualCarPoint)
		mockedConstructor.EXPECT().ActualCarStop().Return(mockedCarStop)

		// Drone1 supports client1: Start flight and move to client1
		mockedConstructor.EXPECT().DroneSupport(mockedDrone1, client1, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone1, client1)

		// Drone2 supports client2: Start flight and move to client2
		mockedConstructor.EXPECT().DroneSupport(mockedDrone2, client2, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(mockedDrone2, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone2, client2)

		// Drone1 does not support client3: Land drone1
		mockedConstructor.EXPECT().DroneSupport(mockedDrone1, client3, actualCarPoint).Return(false)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
		mockedConstructor.EXPECT().LandDrone(mockedDrone1, mockedCarStop)

		// Drone2 supports client3: Move to client3
		mockedConstructor.EXPECT().DroneSupport(mockedDrone2, client3, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone2, client3)

		// Drone1 supports client4: Start flight and move to client4
		mockedConstructor.EXPECT().DroneSupport(mockedDrone1, client4, actualCarPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone1, client4)

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
