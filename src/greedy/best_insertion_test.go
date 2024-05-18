package greedy

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BestInsertion", func() {
	var itineraryList []itinerary.Itinerary
	var mockCtrl *gomock.Controller
	var mockedItinerary1 *mockitinerary.MockItinerary
	var mockedItinerary2 *mockitinerary.MockItinerary

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedItinerary1 = mockitinerary.NewMockItinerary(mockCtrl)
		mockedItinerary2 = mockitinerary.NewMockItinerary(mockCtrl)
		itineraryList = []itinerary.Itinerary{mockedItinerary1, mockedItinerary2}
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
			mockedItinerary1.EXPECT().ActualCarPoint().Return(initialPoint).AnyTimes()
			mockedItinerary1.EXPECT().CarSupport(client3, deposit1).Return(true)
			mockedItinerary1.EXPECT().MoveCar(client3)
			mockedItinerary1.EXPECT().CarSupport(client5, deposit2).Return(true)
			mockedItinerary1.EXPECT().MoveCar(client5)
			mockedItinerary1.EXPECT().CarSupport(client4, deposit2).Return(true)
			mockedItinerary1.EXPECT().MoveCar(client4)
			mockedItinerary1.EXPECT().MoveCar(deposit1)

			mockedItinerary2.EXPECT().ActualCarPoint().Return(initialPoint).AnyTimes()
			mockedItinerary2.EXPECT().CarSupport(client1, deposit1).Return(true)
			mockedItinerary2.EXPECT().MoveCar(client1)
			mockedItinerary2.EXPECT().CarSupport(client6, deposit2).Return(true)
			mockedItinerary2.EXPECT().MoveCar(client6)
			mockedItinerary2.EXPECT().CarSupport(client2, deposit1).Return(true)
			mockedItinerary2.EXPECT().MoveCar(client2)
			mockedItinerary2.EXPECT().MoveCar(deposit1)

			BestInsertion(itineraryList, m)
		})
	})
})

var _ = Describe("orderClientsByItinerary", func() {
	var itineraryList []itinerary.Itinerary
	var mockCtrl *gomock.Controller
	var mockedItinerary1 *mockitinerary.MockItinerary
	var mockedItinerary2 *mockitinerary.MockItinerary

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedItinerary1 = mockitinerary.NewMockItinerary(mockCtrl)
		mockedItinerary2 = mockitinerary.NewMockItinerary(mockCtrl)
		itineraryList = []itinerary.Itinerary{mockedItinerary1, mockedItinerary2}

		mockedItinerary1.EXPECT().ActualCarPoint().Return(gps.Point{Latitude: 0, Longitude: 0}).AnyTimes()
		mockedItinerary2.EXPECT().ActualCarPoint().Return(gps.Point{Latitude: 0, Longitude: 0}).AnyTimes()
	})

	Context("when clients is empty", func() {
		It("return empty array", func() {
			expectedOrderedClients := map[int][]gps.Point{}
			receivedOrderedClients := orderClientsByItinerary(itineraryList, []gps.Point{})
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
			receivedOrderedClients := orderClientsByItinerary(itineraryList, clients)
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
