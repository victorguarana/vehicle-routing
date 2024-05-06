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

var _ = Describe("closestPoint", func() {
	DescribeTable("when receive cadidates", func(candidatePoints []gps.Point, expectedPoint gps.Point) {
		originPoint := gps.Point{Latitude: 0, Longitude: 0}
		receivedPoint := closestPoint(originPoint, candidatePoints)

		Expect(receivedPoint).To(Equal(expectedPoint))
	},
		Entry("when latitude is equal, return closest point", []gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 1, Longitude: 2}}, gps.Point{Latitude: 1, Longitude: 1}),
		Entry("when longitude is equal, return closest point", []gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 2, Longitude: 1}}, gps.Point{Latitude: 1, Longitude: 1}),
	)

	Context("when there are no candidate points", func() {
		It("return empty", func() {
			originPoint := gps.Point{Latitude: 10, Longitude: 10}
			receivedPoint := closestPoint(originPoint, []gps.Point{})

			Expect(receivedPoint).To(Equal(gps.Point{}))
		})
	})
})

var _ = Describe("finishRoutesOnClosestDeposits", func() {
	var (
		mockCtrl      *gomock.Controller
		mockedCar     *mockvehicles.MockICar
		mockedRoute   *mockroutes.MockIRoute
		mockedCarStop *mockroutes.MockICarStop

		routesList     []routes.IRoute
		closestDeposit gps.Point
		m              gps.Map
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
		mockedRoute = mockroutes.NewMockIRoute(mockCtrl)
		mockedCarStop = mockroutes.NewMockICarStop(mockCtrl)

		routesList = []routes.IRoute{mockedRoute}
		closestDeposit = gps.Point{Latitude: 1, Longitude: 1}
		m = gps.Map{Deposits: []gps.Point{closestDeposit}}

		mockedRoute.EXPECT().Last().Return(mockedCarStop)
		mockedCarStop.EXPECT().Point().Return(closestDeposit)
		mockedRoute.EXPECT().Car().Return(mockedCar)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest deposit and append it to the route", func() {
			mockedCar.EXPECT().Move(closestDeposit)
			mockedRoute.EXPECT().Append(closestDeposit)

			finishRoutesOnClosestDeposits(routesList, m)
		})
	})
})
