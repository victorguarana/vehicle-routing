package greedy

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	mockroutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("ClosestNeighbor", func() {
	var itineraryList []routes.Itinerary
	var mockCtrl *gomock.Controller
	var mockedCar1 *mockvehicles.MockICar
	var mockedCar2 *mockvehicles.MockICar
	var mockedMainRoute1 *mockroutes.MockIMainRoute
	var mockedMainRoute2 *mockroutes.MockIMainRoute
	var mockedMainStop *mockroutes.MockIMainStop

	var initialPoint = gps.Point{Latitude: 0, Longitude: 0}
	var client1 = gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
	var client2 = gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
	var client3 = gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
	var client4 = gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
	var client5 = gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
	var client6 = gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
	var deposit1 = gps.Point{Latitude: 0, Longitude: 0}
	var deposit2 = gps.Point{Latitude: 7, Longitude: 7}
	var m = gps.Map{
		Clients:  []gps.Point{client4, client2, client5, client1, client3, client6},
		Deposits: []gps.Point{deposit1, deposit2},
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicles.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicles.NewMockICar(mockCtrl)
		mockedMainRoute1 = mockroutes.NewMockIMainRoute(mockCtrl)
		mockedMainRoute2 = mockroutes.NewMockIMainRoute(mockCtrl)
		mockedMainStop = mockroutes.NewMockIMainStop(mockCtrl)
		itineraryList = []routes.Itinerary{
			{Car: mockedCar1, Route: mockedMainRoute1},
			{Car: mockedCar2, Route: mockedMainRoute2},
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car supports entire route", func() {
		It("return a route without deposits between clients", func() {
			mockedCar1.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar1.EXPECT().Support(client1, deposit1).Return(true)
			mockedCar1.EXPECT().Move(client1)
			mockedMainRoute1.EXPECT().Append(routes.NewMainStop(client1))

			mockedCar2.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar2.EXPECT().Support(client2, deposit1).Return(true)
			mockedCar2.EXPECT().Move(client2)
			mockedMainRoute2.EXPECT().Append(routes.NewMainStop(client2))

			mockedCar1.EXPECT().ActualPosition().Return(client1)
			mockedCar1.EXPECT().Support(client3, deposit1).Return(true)
			mockedCar1.EXPECT().Move(client3)
			mockedMainRoute1.EXPECT().Append(routes.NewMainStop(client3))

			mockedCar2.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar2.EXPECT().Support(client4, deposit2).Return(true)
			mockedCar2.EXPECT().Move(client4)
			mockedMainRoute2.EXPECT().Append(routes.NewMainStop(client4))

			mockedCar1.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar1.EXPECT().Support(client5, deposit2).Return(true)
			mockedCar1.EXPECT().Move(client5)
			mockedMainRoute1.EXPECT().Append(routes.NewMainStop(client5))

			mockedCar2.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar2.EXPECT().Support(client6, deposit2).Return(true)
			mockedCar2.EXPECT().Move(client6)
			mockedMainRoute2.EXPECT().Append(routes.NewMainStop(client6))

			mockedMainRoute1.EXPECT().Last().Return(mockedMainStop)
			mockedMainStop.EXPECT().Point().Return(client5)
			mockedCar1.EXPECT().Move(deposit2)
			mockedMainRoute1.EXPECT().Append(routes.NewMainStop(deposit2))

			mockedMainRoute2.EXPECT().Last().Return(mockedMainStop)
			mockedMainStop.EXPECT().Point().Return(client6)
			mockedCar2.EXPECT().Move(deposit2)
			mockedMainRoute2.EXPECT().Append(routes.NewMainStop(deposit2))

			ClosestNeighbor(itineraryList, m)
		})
	})
})
