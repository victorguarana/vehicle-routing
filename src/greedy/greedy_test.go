package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/itinerary"
	mockitinerary "github.com/victorguarana/go-vehicle-route/src/itinerary/mocks"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("finishRoutesOnClosestDeposits", func() {
	var mockCtrl *gomock.Controller
	var mockedItinerary *mockitinerary.MockItinerary
	var itineraryList []itinerary.Itinerary
	var closestDeposit = gps.Point{Latitude: 1}
	var actualCarPoint = gps.Point{Latitude: 0}
	var gpsMap = gps.Map{Deposits: []gps.Point{closestDeposit}}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedItinerary = mockitinerary.NewMockItinerary(mockCtrl)
		itineraryList = []itinerary.Itinerary{mockedItinerary}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest deposit and append it to the route", func() {
			mockedItinerary.EXPECT().ActualCarPoint().Return(actualCarPoint)
			mockedItinerary.EXPECT().MoveCar(closestDeposit)
			finishItineraryOnClosestDeposits(itineraryList, gpsMap)
		})
	})
})
