package greedy

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BestInsertion", func() {
	var constructorList []itinerary.Constructor
	var mockCtrl *gomock.Controller
	var mockedConstructor1 *mockitinerary.MockConstructor
	var mockedConstructor2 *mockitinerary.MockConstructor
	var mockedCar1 *mockvehicle.MockICar
	var mockedCar2 *mockvehicle.MockICar

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockCtrl)
		mockedConstructor2 = mockitinerary.NewMockConstructor(mockCtrl)
		mockedCar1 = mockvehicle.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicle.NewMockICar(mockCtrl)
		constructorList = []itinerary.Constructor{mockedConstructor1, mockedConstructor2}
	})

	Context("when car supports entire route", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0, Name: "initial"}
		var customer1 = gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
		var customer2 = gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
		var customer3 = gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
		var customer4 = gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
		var customer5 = gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
		var customer6 = gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
		var warehouse1 = gps.Point{Latitude: 0, Longitude: 0, Name: "warehouse1"}
		var warehouse2 = gps.Point{Latitude: 7, Longitude: 7, Name: "warehouse2"}
		var m = gps.Map{
			Customers:  []gps.Point{customer4, customer2, customer5, customer1, customer3, customer6},
			Warehouses: []gps.Point{warehouse1, warehouse2},
		}

		It("return a route without warehouses between customers", func() {
			mockedConstructor1.EXPECT().Car().Return(mockedCar1).AnyTimes()
			mockedConstructor1.EXPECT().ActualCarPoint().Return(initialPoint).AnyTimes()
			mockedCar1.EXPECT().Support(customer3, warehouse1).Return(true)
			mockedConstructor1.EXPECT().MoveCar(customer3)
			mockedCar1.EXPECT().Support(customer5, warehouse2).Return(true)
			mockedConstructor1.EXPECT().MoveCar(customer5)
			mockedCar1.EXPECT().Support(customer4, warehouse2).Return(true)
			mockedConstructor1.EXPECT().MoveCar(customer4)
			mockedConstructor1.EXPECT().MoveCar(warehouse1)

			mockedConstructor2.EXPECT().Car().Return(mockedCar2).AnyTimes()
			mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint).AnyTimes()
			mockedCar2.EXPECT().Support(customer1, warehouse1).Return(true)
			mockedConstructor2.EXPECT().MoveCar(customer1)
			mockedCar2.EXPECT().Support(customer6, warehouse2).Return(true)
			mockedConstructor2.EXPECT().MoveCar(customer6)
			mockedCar2.EXPECT().Support(customer2, warehouse1).Return(true)
			mockedConstructor2.EXPECT().MoveCar(customer2)
			mockedConstructor2.EXPECT().MoveCar(warehouse1)

			BestInsertion(constructorList, m)
		})
	})
})

var _ = Describe("orderCustomersByItinerary", func() {
	var constructorList []itinerary.Constructor
	var mockCtrl *gomock.Controller
	var mockedConstructor1 *mockitinerary.MockConstructor
	var mockedConstructor2 *mockitinerary.MockConstructor

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockCtrl)
		mockedConstructor2 = mockitinerary.NewMockConstructor(mockCtrl)
		constructorList = []itinerary.Constructor{mockedConstructor1, mockedConstructor2}

		mockedConstructor1.EXPECT().ActualCarPoint().Return(gps.Point{Latitude: 0, Longitude: 0}).AnyTimes()
		mockedConstructor2.EXPECT().ActualCarPoint().Return(gps.Point{Latitude: 0, Longitude: 0}).AnyTimes()
	})

	Context("when customers is empty", func() {
		It("return empty array", func() {
			expectedOrderedCustomers := map[int][]gps.Point{}
			receivedOrderedCustomers := orderCustomersByItinerary(constructorList, []gps.Point{})
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})

	Context("when customers has more than one element", func() {
		var customer1 = gps.Point{Latitude: 1, Longitude: 6}
		var customer2 = gps.Point{Latitude: 2, Longitude: 5}
		var customer3 = gps.Point{Latitude: 3, Longitude: 4}
		var customer4 = gps.Point{Latitude: 4, Longitude: 3}
		var customer5 = gps.Point{Latitude: 5, Longitude: 2}
		var customer6 = gps.Point{Latitude: 6, Longitude: 1}
		var customers = []gps.Point{customer5, customer2, customer4, customer6, customer1, customer3}

		It("return ordered customers", func() {
			var expectedOrderedCustomers = map[int][]gps.Point{
				0: {customer1, customer4, customer5},
				1: {customer6, customer3, customer2},
			}
			receivedOrderedCustomers := orderCustomersByItinerary(constructorList, customers)
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})
})

var _ = Describe("insertInBestPosition", func() {
	Context("when orderedCustomers is empty", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newCustomer = gps.Point{Latitude: 1, Longitude: 1}
		var orderedCustomers = []gps.Point{}

		It("return slice with new customer", func() {
			receivedOrderedCustomers := insertInBestPosition(initialPoint, newCustomer, orderedCustomers)
			expectedOrderedCustomers := []gps.Point{newCustomer}
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})

	Context("when best insertion is the first position", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newCustomer = gps.Point{Latitude: 1, Longitude: 1}
		var orderedCustomers = []gps.Point{
			{Latitude: 2, Longitude: 2},
			{Latitude: 3, Longitude: 3},
			{Latitude: 4, Longitude: 4},
			{Latitude: 5, Longitude: 5},
		}

		It("insert in first position", func() {
			receivedOrderedCustomers := insertInBestPosition(initialPoint, newCustomer, orderedCustomers)
			expectedOrderedCustomers := []gps.Point{
				newCustomer,
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
				{Latitude: 4, Longitude: 4},
				{Latitude: 5, Longitude: 5},
			}
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})

	Context("when best insertion is in the middle", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newCustomer = gps.Point{Latitude: 3, Longitude: 3}
		var orderedCustomers = []gps.Point{
			{Latitude: 1, Longitude: 1},
			{Latitude: 2, Longitude: 2},
			{Latitude: 4, Longitude: 4},
			{Latitude: 5, Longitude: 5},
		}

		It("insert in the middle", func() {
			receivedOrderedCustomers := insertInBestPosition(initialPoint, newCustomer, orderedCustomers)
			expectedOrderedCustomers := []gps.Point{
				{Latitude: 1, Longitude: 1},
				{Latitude: 2, Longitude: 2},
				newCustomer,
				{Latitude: 4, Longitude: 4},
				{Latitude: 5, Longitude: 5},
			}
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})

	Context("when best insertion is at the end", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newCustomer = gps.Point{Latitude: 5, Longitude: 1}
		var orderedCustomers = []gps.Point{
			{Latitude: 1, Longitude: 5},
			{Latitude: 2, Longitude: 4},
			{Latitude: 3, Longitude: 3},
			{Latitude: 4, Longitude: 2},
		}

		It("insert at the end", func() {
			receivedOrderedCustomers := insertInBestPosition(initialPoint, newCustomer, orderedCustomers)
			expectedOrderedCustomers := []gps.Point{
				{Latitude: 1, Longitude: 5},
				{Latitude: 2, Longitude: 4},
				{Latitude: 3, Longitude: 3},
				{Latitude: 4, Longitude: 2},
				newCustomer,
			}
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})

	Context("when new customer is behind initial point", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newCustomer = gps.Point{Latitude: -2, Longitude: -2}
		var orderedCustomers = []gps.Point{
			{Latitude: 1, Longitude: 1},
			{Latitude: 2, Longitude: 2},
			{Latitude: 3, Longitude: 3},
			{Latitude: 4, Longitude: 4},
		}

		It("insert in first position", func() {
			receivedOrderedCustomers := insertInBestPosition(initialPoint, newCustomer, orderedCustomers)
			expectedOrderedCustomers := []gps.Point{
				newCustomer,
				{Latitude: 1, Longitude: 1},
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
				{Latitude: 4, Longitude: 4},
			}
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})

	Context("when new customer is between initial point and first customer", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newCustomer = gps.Point{Latitude: 4, Longitude: 4}
		var orderedCustomers = []gps.Point{
			{Latitude: 5, Longitude: 5},
		}

		It("insert in first position", func() {
			receivedOrderedCustomers := insertInBestPosition(initialPoint, newCustomer, orderedCustomers)
			expectedOrderedCustomers := []gps.Point{
				newCustomer,
				{Latitude: 5, Longitude: 5},
			}
			Expect(receivedOrderedCustomers).To(Equal(expectedOrderedCustomers))
		})
	})
})
