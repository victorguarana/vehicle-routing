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
	var mockedCar1 *mockvehicle.MockICar
	var mockedCar2 *mockvehicle.MockICar
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedDrones []vehicle.IDrone
	var initialPoint = gps.Point{Latitude: 0}
	var customer1 = gps.Point{Latitude: 10}
	var customer2 = gps.Point{Latitude: 20}
	var customer3 = gps.Point{Latitude: 30}
	var customer4 = gps.Point{Latitude: 40}
	var customer5 = gps.Point{Latitude: 50}
	var customer6 = gps.Point{Latitude: 60}
	var customer7 = gps.Point{Latitude: 70}
	var customers = []gps.Point{customer1, customer2, customer3, customer4, customer5, customer6, customer7}
	var warehouse = gps.Point{Latitude: 0}
	var gpsMap = gps.Map{Customers: customers, Warehouses: []gps.Point{warehouse}}
	var neighborhoodDistance = 10.0

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedConstructor2 = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		constructorList = []itinerary.Constructor{mockedConstructor1, mockedConstructor2}
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedCar1 = mockvehicle.NewMockICar(mockedCtrl)
		mockedCar2 = mockvehicle.NewMockICar(mockedCtrl)
		mockedDrones = []vehicle.IDrone{mockedDrone1, mockedDrone2}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should cover customers with drones", func() {
		// First iteration
		mockedConstructor1.EXPECT().Car().Return(mockedCar1)
		mockedConstructor1.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedConstructor1.EXPECT().MoveCar(customer2)
		mockedCar1.EXPECT().Drones().Return(mockedDrones)
		mockedConstructor1.EXPECT().ActualCarPoint().Return(customer2)
		mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedDrone1.EXPECT().Support(customer1, customer2).Return(true)
		mockedDrone1.EXPECT().IsFlying().Return(false)
		mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor1.EXPECT().MoveDrone(mockedDrone1, customer1)
		mockedDrone2.EXPECT().Support(customer3, customer2).Return(true)
		mockedDrone2.EXPECT().IsFlying().Return(false)
		mockedConstructor1.EXPECT().StartDroneFlight(mockedDrone2, mockedCarStop)
		mockedConstructor1.EXPECT().MoveDrone(mockedDrone2, customer3)
		mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop)

		// Second iteration
		mockedConstructor2.EXPECT().Car().Return(mockedCar2)
		mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint)
		mockedConstructor2.EXPECT().MoveCar(customer5)
		mockedCar2.EXPECT().Drones().Return(mockedDrones)
		mockedConstructor2.EXPECT().ActualCarPoint().Return(customer5)
		mockedConstructor2.EXPECT().ActualCarStop().Return(mockedCarStop)
		mockedDrone1.EXPECT().Support(customer4, customer5).Return(true)
		mockedDrone1.EXPECT().IsFlying().Return(false)
		mockedConstructor2.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor2.EXPECT().MoveDrone(mockedDrone1, customer4)
		mockedDrone2.EXPECT().Support(customer6, customer5).Return(true)
		mockedDrone2.EXPECT().IsFlying().Return(false)
		mockedConstructor2.EXPECT().StartDroneFlight(mockedDrone2, mockedCarStop)
		mockedConstructor2.EXPECT().MoveDrone(mockedDrone2, customer6)
		mockedConstructor2.EXPECT().LandAllDrones(mockedCarStop)

		// Third iteration
		mockedConstructor1.EXPECT().ActualCarPoint().Return(customer6)
		mockedConstructor1.EXPECT().MoveCar(customer7)

		// Finish routes on closest warehouses
		mockedConstructor1.EXPECT().ActualCarPoint().Return(customer7)
		mockedConstructor1.EXPECT().MoveCar(warehouse)
		mockedConstructor2.EXPECT().ActualCarPoint().Return(customer6)
		mockedConstructor2.EXPECT().MoveCar(warehouse)

		CoveringWithDrones(constructorList, gpsMap, neighborhoodDistance)
	})
})

var _ = Describe("deliverNeighborsWithDrones", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor *mockitinerary.MockConstructor
	var mockedCarStop *mockroute.MockIMainStop
	var mockedCar *mockvehicle.MockICar
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedDrones []vehicle.IDrone
	var actualCarPoint = gps.Point{Latitude: 0}
	var customer1 = gps.Point{Latitude: 1}
	var customer2 = gps.Point{Latitude: 2}
	var customer3 = gps.Point{Latitude: 3}
	var customer4 = gps.Point{Latitude: 4}
	var neighbors = []gps.Point{customer1, customer2, customer3, customer4}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedCar = mockvehicle.NewMockICar(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrones = []vehicle.IDrone{mockedDrone1, mockedDrone2}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should deliver neighbors with drones", func() {
		mockedConstructor.EXPECT().Car().Return(mockedCar)
		mockedCar.EXPECT().Drones().Return(mockedDrones)
		mockedConstructor.EXPECT().ActualCarPoint().Return(actualCarPoint)
		mockedConstructor.EXPECT().ActualCarStop().Return(mockedCarStop)

		// Drone1 supports customer1: Start flight and move to customer1
		mockedDrone1.EXPECT().Support(customer1, actualCarPoint).Return(true)
		mockedDrone1.EXPECT().IsFlying().Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone1, customer1)

		// Drone2 supports customer2: Start flight and move to customer2
		mockedDrone2.EXPECT().Support(customer2, actualCarPoint).Return(true)
		mockedDrone2.EXPECT().IsFlying().Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(mockedDrone2, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone2, customer2)

		// Drone1 does not support customer3: Land drone1
		mockedDrone1.EXPECT().Support(customer3, actualCarPoint).Return(false)
		mockedDrone1.EXPECT().IsFlying().Return(true)
		mockedConstructor.EXPECT().LandDrone(mockedDrone1, mockedCarStop)

		// Drone2 supports customer3: Move to customer3
		mockedDrone2.EXPECT().Support(customer3, actualCarPoint).Return(true)
		mockedDrone2.EXPECT().IsFlying().Return(true)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone2, customer3)

		// Drone1 supports customer4: Start flight and move to customer4
		mockedDrone1.EXPECT().Support(customer4, actualCarPoint).Return(true)
		mockedDrone1.EXPECT().IsFlying().Return(false)
		mockedConstructor.EXPECT().StartDroneFlight(mockedDrone1, mockedCarStop)
		mockedConstructor.EXPECT().MoveDrone(mockedDrone1, customer4)

		// Land all drones
		mockedConstructor.EXPECT().LandAllDrones(mockedCarStop)

		deliverNeighborsWithDrones(mockedConstructor, neighbors)
	})
})

var _ = Describe("removeCustomerAndItsNeighborsFromMap", func() {
	var customer0 = gps.Point{Latitude: 0}
	var customer1 = gps.Point{Latitude: 10}
	var customer2 = gps.Point{Latitude: 20}
	var customer3 = gps.Point{Latitude: 30}
	var customer4 = gps.Point{Latitude: 40}
	var customer5 = gps.Point{Latitude: 50}
	var customer6 = gps.Point{Latitude: 60}
	var neighborhood = gps.Neighborhood{
		customer0: {customer1},
		customer1: {customer0, customer2},
		customer2: {customer1, customer3},
		customer3: {customer2, customer4},
		customer4: {customer3, customer5},
		customer5: {customer4, customer6},
		customer6: {customer5},
	}

	It("should remove customer and its neighbors from map", func() {
		expectedNeighborhood := gps.Neighborhood{
			customer0: {customer1},
			customer1: {customer0},
			customer5: {customer6},
			customer6: {customer5},
		}
		removeCustomerAndItsNeighborsFromMap(customer3, neighborhood)
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
