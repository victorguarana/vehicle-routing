package greedy

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BestInsertion", func() {
	var carsList []vehicles.ICar
	var mockCtrl *gomock.Controller
	var mockedCar1 *mockvehicles.MockICar
	var mockedCar2 *mockvehicles.MockICar

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicles.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicles.NewMockICar(mockCtrl)
		carsList = []vehicles.ICar{mockedCar1, mockedCar2}
	})

	Context("when car supports entire route", func() {
		var initialPoint = gps.Point{Latitude: 0, Longitude: 0, Name: "initial"}
		var client1 = gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
		var client2 = gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
		var client3 = gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
		var client4 = gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
		var client5 = gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
		var client6 = gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
		var deposit1 = gps.Point{Latitude: 0, Longitude: 0, Name: "deposit1"}
		var deposit2 = gps.Point{Latitude: 7, Longitude: 7, Name: "deposit2"}
		var m = gps.Map{
			Clients:  []gps.Point{client4, client2, client5, client1, client3, client6},
			Deposits: []gps.Point{deposit1, deposit2},
		}

		It("return a route without deposits between clients", func() {
			mockedCar1.EXPECT().ActualPoint().Return(initialPoint).AnyTimes()
			mockedCar1.EXPECT().Support(client3, deposit1).Return(true)
			mockedCar1.EXPECT().Move(routes.NewMainStop(client3))
			mockedCar1.EXPECT().Support(client5, deposit2).Return(true)
			mockedCar1.EXPECT().Move(routes.NewMainStop(client5))
			mockedCar1.EXPECT().Support(client4, deposit2).Return(true)
			mockedCar1.EXPECT().Move(routes.NewMainStop(client4))
			mockedCar1.EXPECT().Move(routes.NewMainStop(deposit1))

			mockedCar2.EXPECT().ActualPoint().Return(initialPoint).AnyTimes()
			mockedCar2.EXPECT().Support(client1, deposit1).Return(true)
			mockedCar2.EXPECT().Move(routes.NewMainStop(client1))
			mockedCar2.EXPECT().Support(client6, deposit2).Return(true)
			mockedCar2.EXPECT().Move(routes.NewMainStop(client6))
			mockedCar2.EXPECT().Support(client2, deposit1).Return(true)
			mockedCar2.EXPECT().Move(routes.NewMainStop(client2))
			mockedCar2.EXPECT().Move(routes.NewMainStop(deposit1))

			BestInsertion(carsList, m)
		})
	})
})

var _ = Describe("orderedClientsByCars", func() {
	var carsList []vehicles.ICar
	var mockCtrl *gomock.Controller
	var mockedCar1 *mockvehicles.MockICar
	var mockedCar2 *mockvehicles.MockICar

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicles.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicles.NewMockICar(mockCtrl)
		carsList = []vehicles.ICar{mockedCar1, mockedCar2}

		mockedCar1.EXPECT().ActualPoint().Return(gps.Point{Latitude: 0, Longitude: 0}).AnyTimes()
		mockedCar2.EXPECT().ActualPoint().Return(gps.Point{Latitude: 0, Longitude: 0}).AnyTimes()
	})

	Context("when clients is empty", func() {
		It("return empty array", func() {
			expectedOrderedClients := map[vehicles.ICar][]gps.Point{}
			receivedOrderedClients := orderedClientsByCars(carsList, []gps.Point{})
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
			var expectedOrderedClients = map[vehicles.ICar][]gps.Point{
				mockedCar1: {client1, client4, client5},
				mockedCar2: {client6, client3, client2},
			}
			receivedOrderedClients := orderedClientsByCars(carsList, clients)
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
