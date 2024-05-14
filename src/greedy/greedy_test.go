package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"

	"go.uber.org/mock/gomock"

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
	var mockCtrl *gomock.Controller
	var mockedCar *mockvehicles.MockICar
	var carsList []vehicles.ICar
	var closestDeposit = gps.Point{Latitude: 1}
	var actualCarPoint = gps.Point{Latitude: 0}
	var gpsMap = gps.Map{Deposits: []gps.Point{closestDeposit}}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
		carsList = []vehicles.ICar{mockedCar}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest deposit and append it to the route", func() {
			mockedCar.EXPECT().ActualPoint().Return(actualCarPoint)
			mockedCar.EXPECT().Move(routes.NewMainStop(closestDeposit))
			finishItineraryOnClosestDeposits(carsList, gpsMap)
		})
	})
})
