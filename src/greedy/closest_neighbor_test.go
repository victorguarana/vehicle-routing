package greedy

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	mockroutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClosestNeighbor", func() {
	var (
		mockCtrl      *gomock.Controller
		mockedCar1    *mockvehicles.MockICar
		mockedCar2    *mockvehicles.MockICar
		mockedRoute1  *mockroutes.MockIRoute
		mockedRoute2  *mockroutes.MockIRoute
		mockedCarStop *mockroutes.MockICarStop

		routesList []routes.IRoute
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicles.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicles.NewMockICar(mockCtrl)
		mockedRoute1 = mockroutes.NewMockIRoute(mockCtrl)
		mockedRoute2 = mockroutes.NewMockIRoute(mockCtrl)
		mockedCarStop = mockroutes.NewMockICarStop(mockCtrl)

		routesList = []routes.IRoute{mockedRoute1, mockedRoute2}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car supports entire route", func() {
		It("return a route without deposits between clients", func() {
			initialPoint := gps.Point{Latitude: 0, Longitude: 0}
			client1 := gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
			client2 := gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
			client3 := gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
			client4 := gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
			client5 := gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
			client6 := gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
			deposit1 := gps.Point{Latitude: 0, Longitude: 0}
			deposit2 := gps.Point{Latitude: 7, Longitude: 7}

			m := gps.Map{
				Clients:  []gps.Point{client4, client2, client5, client1, client3, client6},
				Deposits: []gps.Point{deposit1, deposit2},
			}

			mockedRoute1.EXPECT().Car().Return(mockedCar1).AnyTimes()
			mockedRoute2.EXPECT().Car().Return(mockedCar2).AnyTimes()

			mockedCar1.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar1.EXPECT().Support(client1, deposit1).Return(true)
			mockedCar1.EXPECT().Move(client1)
			mockedRoute1.EXPECT().Append(client1)

			mockedCar2.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar2.EXPECT().Support(client2, deposit1).Return(true)
			mockedCar2.EXPECT().Move(client2)
			mockedRoute2.EXPECT().Append(client2)

			mockedCar1.EXPECT().ActualPosition().Return(client1)
			mockedCar1.EXPECT().Support(client3, deposit1).Return(true)
			mockedCar1.EXPECT().Move(client3)
			mockedRoute1.EXPECT().Append(client3)

			mockedCar2.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar2.EXPECT().Support(client4, deposit2).Return(true)
			mockedCar2.EXPECT().Move(client4)
			mockedRoute2.EXPECT().Append(client4)

			mockedCar1.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar1.EXPECT().Support(client5, deposit2).Return(true)
			mockedCar1.EXPECT().Move(client5)
			mockedRoute1.EXPECT().Append(client5)

			mockedCar2.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar2.EXPECT().Support(client6, deposit2).Return(true)
			mockedCar2.EXPECT().Move(client6)
			mockedRoute2.EXPECT().Append(client6)

			mockedRoute1.EXPECT().Last().Return(mockedCarStop)
			mockedCarStop.EXPECT().Point().Return(client5)
			mockedCar1.EXPECT().Move(deposit2)
			mockedRoute1.EXPECT().Append(deposit2)

			mockedRoute2.EXPECT().Last().Return(mockedCarStop)
			mockedCarStop.EXPECT().Point().Return(client6)
			mockedCar2.EXPECT().Move(deposit2)
			mockedRoute2.EXPECT().Append(deposit2)

			Expect(ClosestNeighbor(routesList, m)).To(Succeed())
		})
	})
})

var _ = Describe("removePoint", func() {
	var points []gps.Point

	BeforeEach(func() {
		points = []gps.Point{
			{Latitude: 1, Longitude: 1},
			{Latitude: 2, Longitude: 2},
			{Latitude: 3, Longitude: 3},
		}
	})

	Context("when point is in the slice", func() {
		Context("when point is the first element", func() {
			It("remove the point", func() {
				point := points[0]
				expectedPoints := []gps.Point{{Latitude: 2, Longitude: 2}, {Latitude: 3, Longitude: 3}}
				receivedPoints := removePoint(points, point)

				Expect(receivedPoints).To(Equal(expectedPoints))
			})
		})

		Context("when point is in the middle", func() {
			It("remove the point", func() {
				point := points[1]
				expectedPoints := []gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 3, Longitude: 3}}
				receivedPoints := removePoint(points, point)

				Expect(receivedPoints).To(Equal(expectedPoints))
			})
		})

		Context("when point is the last element", func() {
			It("remove the point", func() {
				point := points[2]
				expectedPoints := []gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 2, Longitude: 2}}
				receivedPoints := removePoint(points, point)

				Expect(receivedPoints).To(Equal(expectedPoints))
			})
		})
	})

})
