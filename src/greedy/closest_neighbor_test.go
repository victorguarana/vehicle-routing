package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mocks"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("ClosestNeighbor", func() {
	var carsList []itinerary.Itinerary
	var mockCtrl *gomock.Controller
	var mockedItinerary1 *mockitinerary.MockItinerary
	var mockedItinerary2 *mockitinerary.MockItinerary

	var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
	var client1 = gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
	var client2 = gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
	var client3 = gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
	var client4 = gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
	var client5 = gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
	var client6 = gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
	var warehouse1 = gps.Point{Latitude: 0, Longitude: 0}
	var warehouse2 = gps.Point{Latitude: 7, Longitude: 7}
	var m = gps.Map{
		Clients:    []gps.Point{client4, client2, client5, client1, client3, client6},
		Warehouses: []gps.Point{warehouse1, warehouse2},
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedItinerary1 = mockitinerary.NewMockItinerary(mockCtrl)
		mockedItinerary2 = mockitinerary.NewMockItinerary(mockCtrl)
		carsList = []itinerary.Itinerary{mockedItinerary1, mockedItinerary2}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car supports entire route", func() {
		It("return a route without warehouses between clients", func() {
			mockedItinerary1.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedItinerary1.EXPECT().CarSupport(client1, warehouse1).Return(true)
			mockedItinerary1.EXPECT().MoveCar(client1)
			mockedItinerary1.EXPECT().ActualCarPoint().Return(client1)
			mockedItinerary1.EXPECT().CarSupport(client3, warehouse1).Return(true)
			mockedItinerary1.EXPECT().MoveCar(client3)
			mockedItinerary1.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedItinerary1.EXPECT().CarSupport(client5, warehouse2).Return(true)
			mockedItinerary1.EXPECT().MoveCar(client5)
			mockedItinerary1.EXPECT().ActualCarPoint().Return(client5)
			mockedItinerary1.EXPECT().MoveCar(warehouse2)

			mockedItinerary2.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedItinerary2.EXPECT().CarSupport(client2, warehouse1).Return(true)
			mockedItinerary2.EXPECT().MoveCar(client2)
			mockedItinerary2.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedItinerary2.EXPECT().CarSupport(client4, warehouse2).Return(true)
			mockedItinerary2.EXPECT().MoveCar(client4)
			mockedItinerary2.EXPECT().ActualCarPoint().Return(initialPoint)
			mockedItinerary2.EXPECT().CarSupport(client6, warehouse2).Return(true)
			mockedItinerary2.EXPECT().MoveCar(client6)
			mockedItinerary2.EXPECT().ActualCarPoint().Return(client6)
			mockedItinerary2.EXPECT().MoveCar(warehouse2)

			ClosestNeighbor(carsList, m)
		})
	})
})
