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
		var client1 = gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
		var client2 = gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
		var client3 = gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
		var client4 = gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
		var client5 = gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
		var client6 = gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
		var warehouse1 = gps.Point{Latitude: 0, Longitude: 0, Name: "warehouse1"}
		var warehouse2 = gps.Point{Latitude: 7, Longitude: 7, Name: "warehouse2"}
		var m = gps.Map{
			Clients:    []gps.Point{client4, client2, client5, client1, client3, client6},
			Warehouses: []gps.Point{warehouse1, warehouse2},
		}

		It("return a route without warehouses between clients", func() {
			mockedConstructor1.EXPECT().Car().Return(mockedCar1).AnyTimes()
			mockedConstructor1.EXPECT().ActualCarPoint().Return(initialPoint).AnyTimes()
			mockedCar1.EXPECT().Support(client3, warehouse1).Return(true)
			mockedConstructor1.EXPECT().MoveCar(client3)
			mockedCar1.EXPECT().Support(client5, warehouse2).Return(true)
			mockedConstructor1.EXPECT().MoveCar(client5)
			mockedCar1.EXPECT().Support(client4, warehouse2).Return(true)
			mockedConstructor1.EXPECT().MoveCar(client4)
			mockedConstructor1.EXPECT().MoveCar(warehouse1)

			mockedConstructor2.EXPECT().Car().Return(mockedCar2).AnyTimes()
			mockedConstructor2.EXPECT().ActualCarPoint().Return(initialPoint).AnyTimes()
			mockedCar2.EXPECT().Support(client1, warehouse1).Return(true)
			mockedConstructor2.EXPECT().MoveCar(client1)
			mockedCar2.EXPECT().Support(client6, warehouse2).Return(true)
			mockedConstructor2.EXPECT().MoveCar(client6)
			mockedCar2.EXPECT().Support(client2, warehouse1).Return(true)
			mockedConstructor2.EXPECT().MoveCar(client2)
			mockedConstructor2.EXPECT().MoveCar(warehouse1)

			BestInsertion(constructorList, m)
		})
	})
})

var _ = Describe("orderClientsByItinerary", func() {
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

	Context("when clients is empty", func() {
		It("return empty array", func() {
			expectedOrderedClients := map[int][]gps.Point{}
			receivedOrderedClients := orderClientsByItinerary(constructorList, []gps.Point{})
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when clients has more than one element", func() {
		var client1 = gps.Point{Latitude: 1, Longitude: 6}
		var client2 = gps.Point{Latitude: 2, Longitude: 5}
		var client3 = gps.Point{Latitude: 3, Longitude: 4}
		var client4 = gps.Point{Latitude: 4, Longitude: 3}
		var client5 = gps.Point{Latitude: 5, Longitude: 2}
		var client6 = gps.Point{Latitude: 6, Longitude: 1}
		var clients = []gps.Point{client5, client2, client4, client6, client1, client3}

		It("return ordered clients", func() {
			var expectedOrderedClients = map[int][]gps.Point{
				0: {client1, client4, client5},
				1: {client6, client3, client2},
			}
			receivedOrderedClients := orderClientsByItinerary(constructorList, clients)
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})
})

var _ = Describe("insertInBestPosition", func() {
	Context("when orderedClients is empty", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newClient = gps.Point{Latitude: 1, Longitude: 1}
		var orderedClients = []gps.Point{}

		It("return slice with new client", func() {
			receivedOrderedClients := insertInBestPosition(initialPoint, newClient, orderedClients)
			expectedOrderedClients := []gps.Point{newClient}
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when best insertion is the first position", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newClient = gps.Point{Latitude: 1, Longitude: 1}
		var orderedClients = []gps.Point{
			{Latitude: 2, Longitude: 2},
			{Latitude: 3, Longitude: 3},
			{Latitude: 4, Longitude: 4},
			{Latitude: 5, Longitude: 5},
		}

		It("insert in first position", func() {
			receivedOrderedClients := insertInBestPosition(initialPoint, newClient, orderedClients)
			expectedOrderedClients := []gps.Point{
				newClient,
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
				{Latitude: 4, Longitude: 4},
				{Latitude: 5, Longitude: 5},
			}
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when best insertion is in the middle", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newClient = gps.Point{Latitude: 3, Longitude: 3}
		var orderedClients = []gps.Point{
			{Latitude: 1, Longitude: 1},
			{Latitude: 2, Longitude: 2},
			{Latitude: 4, Longitude: 4},
			{Latitude: 5, Longitude: 5},
		}

		It("insert in the middle", func() {
			receivedOrderedClients := insertInBestPosition(initialPoint, newClient, orderedClients)
			expectedOrderedClients := []gps.Point{
				{Latitude: 1, Longitude: 1},
				{Latitude: 2, Longitude: 2},
				newClient,
				{Latitude: 4, Longitude: 4},
				{Latitude: 5, Longitude: 5},
			}
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when best insertion is at the end", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newClient = gps.Point{Latitude: 5, Longitude: 1}
		var orderedClients = []gps.Point{
			{Latitude: 1, Longitude: 5},
			{Latitude: 2, Longitude: 4},
			{Latitude: 3, Longitude: 3},
			{Latitude: 4, Longitude: 2},
		}

		It("insert at the end", func() {
			receivedOrderedClients := insertInBestPosition(initialPoint, newClient, orderedClients)
			expectedOrderedClients := []gps.Point{
				{Latitude: 1, Longitude: 5},
				{Latitude: 2, Longitude: 4},
				{Latitude: 3, Longitude: 3},
				{Latitude: 4, Longitude: 2},
				newClient,
			}
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when new client is behind initial point", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newClient = gps.Point{Latitude: -2, Longitude: -2}
		var orderedClients = []gps.Point{
			{Latitude: 1, Longitude: 1},
			{Latitude: 2, Longitude: 2},
			{Latitude: 3, Longitude: 3},
			{Latitude: 4, Longitude: 4},
		}

		It("insert in first position", func() {
			receivedOrderedClients := insertInBestPosition(initialPoint, newClient, orderedClients)
			expectedOrderedClients := []gps.Point{
				newClient,
				{Latitude: 1, Longitude: 1},
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
				{Latitude: 4, Longitude: 4},
			}
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when new client is between initial point and first client", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
		var newClient = gps.Point{Latitude: 4, Longitude: 4}
		var orderedClients = []gps.Point{
			{Latitude: 5, Longitude: 5},
		}

		It("insert in first position", func() {
			receivedOrderedClients := insertInBestPosition(initialPoint, newClient, orderedClients)
			expectedOrderedClients := []gps.Point{
				newClient,
				{Latitude: 5, Longitude: 5},
			}
			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})
})
