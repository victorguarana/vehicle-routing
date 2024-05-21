package cost

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/src/gps"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mocks"
	"github.com/victorguarana/vehicle-routing/src/routes"
	mockroutes "github.com/victorguarana/vehicle-routing/src/routes/mocks"
	"github.com/victorguarana/vehicle-routing/src/slc"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CalcTotalFuel", func() {
	Context("when itinerary does not have subroutes", func() {
		var mockedCtrl *gomock.Controller
		var mockedItinerary *mockitinerary.MockItinerary
		var mockedMainStop1 *mockroutes.MockIMainStop
		var mockedMainStop2 *mockroutes.MockIMainStop
		var mockedMainStop3 *mockroutes.MockIMainStop
		var mockedMainStop4 *mockroutes.MockIMainStop
		var mockedMainStop5 *mockroutes.MockIMainStop
		var mainPoint1 = gps.Point{Latitude: 0, Name: "MainPoint1"}
		var mainPoint2 = gps.Point{Latitude: 30, Name: "MainPoint2"}
		var mainPoint3 = gps.Point{Latitude: 10, Name: "MainPoint3"}
		var mainPoint4 = gps.Point{Latitude: 40, Name: "MainPoint4"}
		var mainPoint5 = gps.Point{Latitude: 20, Name: "MainPoint5"}
		var carEfficiency = 10.0
		var droneEfficiency = 100.0

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedItinerary = mockitinerary.NewMockItinerary(mockedCtrl)
			mockedMainStop1 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop2 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop3 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop4 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop5 = mockroutes.NewMockIMainStop(mockedCtrl)
		})

		It("should return the total distance of the route", func() {
			mockedMainStop1.EXPECT().Point().Return(mainPoint1).AnyTimes()
			mockedMainStop2.EXPECT().Point().Return(mainPoint2).AnyTimes()
			mockedMainStop3.EXPECT().Point().Return(mainPoint3).AnyTimes()
			mockedMainStop4.EXPECT().Point().Return(mainPoint4).AnyTimes()
			mockedMainStop5.EXPECT().Point().Return(mainPoint5).AnyTimes()

			mockedMainStop1.EXPECT().StartingSubRoutes().Return(nil)
			mockedMainStop2.EXPECT().StartingSubRoutes().Return(nil)
			mockedMainStop3.EXPECT().StartingSubRoutes().Return(nil)
			mockedMainStop4.EXPECT().StartingSubRoutes().Return(nil)

			mockedItinerary.EXPECT().CarEfficiency().Return(carEfficiency)
			mockedItinerary.EXPECT().DroneEfficiency().Return(droneEfficiency)
			mockedItinerary.EXPECT().RouteIterator().Return(slc.NewIterator([]routes.IMainStop{mockedMainStop1, mockedMainStop2, mockedMainStop3, mockedMainStop4, mockedMainStop5}))

			expectedFuelSpent := 10.0
			receivedFuelSpent := CalcTotalFuel(mockedItinerary)
			Expect(receivedFuelSpent).To(Equal(expectedFuelSpent))
		})
	})

	Context("when itinerary has subroutes", func() {
		var mockedCtrl *gomock.Controller
		var mockedItinerary *mockitinerary.MockItinerary
		var mockedSubRoute1 *mockroutes.MockISubRoute
		var mockedSubRoute2 *mockroutes.MockISubRoute
		var mockedMainStop1 *mockroutes.MockIMainStop
		var mockedMainStop2 *mockroutes.MockIMainStop
		var mockedMainStop3 *mockroutes.MockIMainStop
		var mockedMainStop4 *mockroutes.MockIMainStop
		var mockedMainStop5 *mockroutes.MockIMainStop
		var mockedSubStop1 *mockroutes.MockISubStop
		var mockedSubStop2 *mockroutes.MockISubStop
		var mockedSubStop3 *mockroutes.MockISubStop
		var mockedSubStop4 *mockroutes.MockISubStop
		var mainPoint1 = gps.Point{Latitude: 0, Name: "MainPoint1"}
		var mainPoint2 = gps.Point{Latitude: 5, Name: "MainPoint2"}
		var mainPoint3 = gps.Point{Latitude: 10, Name: "MainPoint3"}
		var mainPoint4 = gps.Point{Latitude: 15, Name: "MainPoint4"}
		var mainPoint5 = gps.Point{Latitude: 20, Name: "MainPoint5"}
		var subPoint1 = gps.Point{Latitude: 8, Name: "SubPoint1"}
		var subPoint2 = gps.Point{Latitude: 15, Name: "SubPoint2"}
		var subPoint3 = gps.Point{Latitude: 5, Name: "SubPoint3"}
		var subPoint4 = gps.Point{Latitude: 25, Name: "SubPoint4"}
		var carEfficiency = 10.0
		var droneEfficiency = 100.0

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedItinerary = mockitinerary.NewMockItinerary(mockedCtrl)
			mockedSubRoute1 = mockroutes.NewMockISubRoute(mockedCtrl)
			mockedSubRoute2 = mockroutes.NewMockISubRoute(mockedCtrl)
			mockedMainStop1 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop2 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop3 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop4 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedMainStop5 = mockroutes.NewMockIMainStop(mockedCtrl)
			mockedSubStop1 = mockroutes.NewMockISubStop(mockedCtrl)
			mockedSubStop2 = mockroutes.NewMockISubStop(mockedCtrl)
			mockedSubStop3 = mockroutes.NewMockISubStop(mockedCtrl)
			mockedSubStop4 = mockroutes.NewMockISubStop(mockedCtrl)
		})

		It("should return the total distance of the route", func() {
			mockedMainStop1.EXPECT().Point().Return(mainPoint1).AnyTimes()
			mockedMainStop2.EXPECT().Point().Return(mainPoint2).AnyTimes()
			mockedMainStop3.EXPECT().Point().Return(mainPoint3).AnyTimes()
			mockedMainStop4.EXPECT().Point().Return(mainPoint4).AnyTimes()
			mockedMainStop5.EXPECT().Point().Return(mainPoint5).AnyTimes()

			mockedSubStop1.EXPECT().Point().Return(subPoint1).AnyTimes()
			mockedSubStop2.EXPECT().Point().Return(subPoint2).AnyTimes()
			mockedSubStop3.EXPECT().Point().Return(subPoint3).AnyTimes()
			mockedSubStop4.EXPECT().Point().Return(subPoint4).AnyTimes()

			mockedMainStop1.EXPECT().StartingSubRoutes().Return([]routes.ISubRoute{mockedSubRoute1})
			mockedMainStop2.EXPECT().StartingSubRoutes().Return(nil)
			mockedMainStop3.EXPECT().StartingSubRoutes().Return([]routes.ISubRoute{mockedSubRoute2})
			mockedMainStop4.EXPECT().StartingSubRoutes().Return(nil)

			mockedSubRoute1.EXPECT().StartingStop().Return(mockedMainStop1)
			mockedSubRoute1.EXPECT().ReturningStop().Return(mockedMainStop3)
			mockedSubRoute2.EXPECT().StartingStop().Return(mockedMainStop3)
			mockedSubRoute2.EXPECT().ReturningStop().Return(mockedMainStop5)

			mockedSubRoute1.EXPECT().First().Return(mockedSubStop1)
			mockedSubRoute1.EXPECT().Iterator().Return(slc.NewIterator([]routes.ISubStop{mockedSubStop1, mockedSubStop2}))
			mockedSubRoute2.EXPECT().First().Return(mockedSubStop3)
			mockedSubRoute2.EXPECT().Iterator().Return(slc.NewIterator([]routes.ISubStop{mockedSubStop3, mockedSubStop4}))

			mockedItinerary.EXPECT().CarEfficiency().Return(carEfficiency)
			mockedItinerary.EXPECT().DroneEfficiency().Return(droneEfficiency)
			mockedItinerary.EXPECT().RouteIterator().Return(slc.NewIterator([]routes.IMainStop{mockedMainStop1, mockedMainStop2, mockedMainStop3, mockedMainStop4, mockedMainStop5}))

			carFuelSpent := 2.0
			droneFuelSpent := 0.5
			expectedFuelSpent := carFuelSpent + droneFuelSpent
			receivedFuelSpent := CalcTotalFuel(mockedItinerary)
			Expect(receivedFuelSpent).To(Equal(expectedFuelSpent))
		})
	})
})
