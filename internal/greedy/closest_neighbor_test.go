package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("ClosestNeighbor", func() {
	var constructorsList []itinerary.Constructor
	var mockCtrl *gomock.Controller
	var mockedConstructor1 *mockitinerary.MockConstructor
	var mockedConstructor2 *mockitinerary.MockConstructor
	var mockedCar1 *mockvehicle.MockICar
	var mockedCar2 *mockvehicle.MockICar

	var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
	var customer1 = gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
	var customer2 = gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
	var customer3 = gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
	var customer4 = gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
	var customer5 = gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
	var customer6 = gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
	var warehouse1 = gps.Point{Latitude: 0, Longitude: 0}
	var warehouse2 = gps.Point{Latitude: 7, Longitude: 7}
	var m = gps.Map{
		Customers:  []gps.Point{customer4, customer2, customer5, customer1, customer3, customer6},
		Warehouses: []gps.Point{warehouse1, warehouse2},
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockCtrl)
		mockedConstructor2 = mockitinerary.NewMockConstructor(mockCtrl)
		mockedCar1 = mockvehicle.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicle.NewMockICar(mockCtrl)
		constructorsList = []itinerary.Constructor{mockedConstructor1, mockedConstructor2}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car supports entire route", func() {
		It("return a route without warehouses between customers", func() {
			mockedConstructor1.EXPECT().Car().Return(mockedCar1).AnyTimes()
			mockedConstructor1.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedCar1.EXPECT().Support(customer1, warehouse1).Return(true)
			mockedConstructor1.EXPECT().MoveCar(customer1)
			mockedConstructor1.EXPECT().ActualCarPoint().Return(customer1)
			mockedCar1.EXPECT().Support(customer3, warehouse1).Return(true)
			mockedConstructor1.EXPECT().MoveCar(customer3)
			mockedConstructor1.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedCar1.EXPECT().Support(customer5, warehouse2).Return(true)
			mockedConstructor1.EXPECT().MoveCar(customer5)
			mockedConstructor1.EXPECT().ActualCarPoint().Return(customer5)
			mockedConstructor1.EXPECT().MoveCar(warehouse2)

			mockedConstructor2.EXPECT().Car().Return(mockedCar2).AnyTimes()
			mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedCar2.EXPECT().Support(customer2, warehouse1).Return(true)
			mockedConstructor2.EXPECT().MoveCar(customer2)
			mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedCar2.EXPECT().Support(customer4, warehouse2).Return(true)
			mockedConstructor2.EXPECT().MoveCar(customer4)
			mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedCar2.EXPECT().Support(customer6, warehouse2).Return(true)
			mockedConstructor2.EXPECT().MoveCar(customer6)
			mockedConstructor2.EXPECT().ActualCarPoint().Return(customer6)
			mockedConstructor2.EXPECT().MoveCar(warehouse2)

			ClosestNeighbor(constructorsList, m)
		})
	})
})
