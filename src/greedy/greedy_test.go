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
	Context("when cadidates latitudes are equal", func() {
		var originPoint = gps.Point{Latitude: 0, Longitude: 0}
		var candidatePoints = []gps.Point{
			{Latitude: 1, Longitude: 2},
			{Latitude: 1, Longitude: 1},
			{Latitude: 1, Longitude: 3},
		}

		It("returns closest point", func() {
			expectedPoint := candidatePoints[1]
			receivedPoint := closestPoint(originPoint, candidatePoints)
			Expect(receivedPoint).To(Equal(expectedPoint))
		})
	})

	Context("when candidates longitudes are equal", func() {
		var originPoint = gps.Point{Latitude: 0, Longitude: 0}
		var candidatePoints = []gps.Point{
			{Latitude: 2, Longitude: 1},
			{Latitude: 1, Longitude: 1},
			{Latitude: 3, Longitude: 1},
		}

		It("returns closest point", func() {
			expectedPoint := candidatePoints[1]
			receivedPoint := closestPoint(originPoint, candidatePoints)
			Expect(receivedPoint).To(Equal(expectedPoint))
		})
	})

	Context("when there are no candidate points", func() {
		It("return empty", func() {
			originPoint := gps.Point{Latitude: 10, Longitude: 10}
			receivedPoint := closestPoint(originPoint, []gps.Point{})

			Expect(receivedPoint).To(Equal(gps.Point{}))
		})
	})
})

var _ = Describe("finishRoutesOnClosestDeposits", func() {
	var itineraryList []routes.Itinerary
	var closestDeposit = gps.Point{Latitude: 1, Longitude: 1}
	var closestDepositMainStop = routes.NewMainStop(closestDeposit)
	var gpsMap = gps.Map{Deposits: []gps.Point{closestDeposit}}

	var mockCtrl *gomock.Controller
	var mockedCar *mockvehicles.MockICar
	var mockedRoute *mockroutes.MockIMainRoute
	var mockedMainStop *mockroutes.MockIMainStop

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
		mockedRoute = mockroutes.NewMockIMainRoute(mockCtrl)
		mockedMainStop = mockroutes.NewMockIMainStop(mockCtrl)
		itineraryList = []routes.Itinerary{{Car: mockedCar, Route: mockedRoute}}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest deposit and append it to the route", func() {
			mockedRoute.EXPECT().Last().Return(mockedMainStop)
			mockedMainStop.EXPECT().Point().Return(closestDeposit)
			mockedCar.EXPECT().Move(closestDeposit)
			mockedRoute.EXPECT().Append(closestDepositMainStop)
			finishItineraryOnClosestDeposits(itineraryList, gpsMap)
		})
	})
})
