package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("ClosestNeighbor", func() {
	var carsList []vehicles.ICar
	var mockCtrl *gomock.Controller
	var mockedCar1 *mockvehicles.MockICar
	var mockedCar2 *mockvehicles.MockICar

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
		carsList = []vehicles.ICar{mockedCar1, mockedCar2}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car supports entire route", func() {
		It("return a route without deposits between clients", func() {
			mockedCar1.EXPECT().ActualPoint().Return(initialPoint)
			mockedCar1.EXPECT().Support(client1, deposit1).Return(true)
			mockedCar1.EXPECT().Move(routes.NewMainStop(client1))
			mockedCar1.EXPECT().ActualPoint().Return(client1)
			mockedCar1.EXPECT().Support(client3, deposit1).Return(true)
			mockedCar1.EXPECT().Move(routes.NewMainStop(client3))
			mockedCar1.EXPECT().ActualPoint().Return(initialPoint)
			mockedCar1.EXPECT().Support(client5, deposit2).Return(true)
			mockedCar1.EXPECT().Move(routes.NewMainStop(client5))
			mockedCar1.EXPECT().ActualPoint().Return(client5)
			mockedCar1.EXPECT().Move(routes.NewMainStop(deposit2))

			mockedCar2.EXPECT().ActualPoint().Return(initialPoint)
			mockedCar2.EXPECT().Support(client2, deposit1).Return(true)
			mockedCar2.EXPECT().Move(routes.NewMainStop(client2))
			mockedCar2.EXPECT().ActualPoint().Return(initialPoint)
			mockedCar2.EXPECT().Support(client4, deposit2).Return(true)
			mockedCar2.EXPECT().Move(routes.NewMainStop(client4))
			mockedCar2.EXPECT().ActualPoint().Return(initialPoint)
			mockedCar2.EXPECT().Support(client6, deposit2).Return(true)
			mockedCar2.EXPECT().Move(routes.NewMainStop(client6))
			mockedCar2.EXPECT().ActualPoint().Return(client6)
			mockedCar2.EXPECT().Move(routes.NewMainStop(deposit2))

			ClosestNeighbor(carsList, m)
		})
	})
})
