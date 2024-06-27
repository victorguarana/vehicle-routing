package measure

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	"github.com/victorguarana/vehicle-routing/internal/route"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/slc"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TimeSpent", func() {
	Context("when itinerary does not have subroutes", func() {
		var mockedCtrl *gomock.Controller
		var mockedItineraryInfo *mockitinerary.MockInfo
		var mockedMainStop1 *mockroute.MockIMainStop
		var mockedMainStop2 *mockroute.MockIMainStop
		var mockedMainStop3 *mockroute.MockIMainStop
		var mockedMainStop4 *mockroute.MockIMainStop
		var mockedMainStop5 *mockroute.MockIMainStop
		var mainPoint1 = gps.Point{Latitude: 0, Name: "MainPoint1"}
		var mainPoint2 = gps.Point{Latitude: 30, Name: "MainPoint2"}
		var mainPoint3 = gps.Point{Latitude: 10, Name: "MainPoint3"}
		var mainPoint4 = gps.Point{Latitude: 40, Name: "MainPoint4"}
		var mainPoint5 = gps.Point{Latitude: 20, Name: "MainPoint5"}
		var carSpeed = 10.0
		var droneSpeed = 20.0

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedItineraryInfo = mockitinerary.NewMockInfo(mockedCtrl)
			mockedMainStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop2 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop3 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop4 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop5 = mockroute.NewMockIMainStop(mockedCtrl)
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
			mockedMainStop5.EXPECT().StartingSubRoutes().Return(nil)

			mockedMainStop1.EXPECT().ReturningSubRoutes().Return(nil)
			mockedMainStop2.EXPECT().ReturningSubRoutes().Return(nil)
			mockedMainStop3.EXPECT().ReturningSubRoutes().Return(nil)
			mockedMainStop4.EXPECT().ReturningSubRoutes().Return(nil)
			mockedMainStop5.EXPECT().ReturningSubRoutes().Return(nil)

			mockedItineraryInfo.EXPECT().RouteIterator().Return(slc.NewIterator([]route.IMainStop{mockedMainStop1, mockedMainStop2, mockedMainStop3, mockedMainStop4, mockedMainStop5}))
			mockedItineraryInfo.EXPECT().CarSpeed().Return(carSpeed)
			mockedItineraryInfo.EXPECT().DroneSpeed().Return(droneSpeed)

			expectedTime := 100.0 / carSpeed
			receivedTime := TimeSpent(mockedItineraryInfo)
			Expect(receivedTime).To(Equal(expectedTime))
		})
	})

	Context("when itinerary has subroutes", func() {
		var mockedCtrl *gomock.Controller
		var mockedItineraryInfo *mockitinerary.MockInfo
		var mockedSubRoute1 *mockroute.MockISubRoute
		var mockedSubRoute2 *mockroute.MockISubRoute
		var mockedMainStop1 *mockroute.MockIMainStop
		var mockedMainStop2 *mockroute.MockIMainStop
		var mockedMainStop3 *mockroute.MockIMainStop
		var mockedMainStop4 *mockroute.MockIMainStop
		var mockedMainStop5 *mockroute.MockIMainStop
		var mockedSubStop1 *mockroute.MockISubStop
		var mockedSubStop2 *mockroute.MockISubStop
		var mockedSubStop3 *mockroute.MockISubStop
		var mockedSubStop4 *mockroute.MockISubStop
		var mainPoint1 = gps.Point{Latitude: 0, Name: "MainPoint1"}
		var mainPoint2 = gps.Point{Latitude: 5, Name: "MainPoint2"}
		var mainPoint3 = gps.Point{Latitude: 10, Name: "MainPoint3"}
		var mainPoint4 = gps.Point{Latitude: 15, Name: "MainPoint4"}
		var mainPoint5 = gps.Point{Latitude: 20, Name: "MainPoint5"}
		var subPoint1 = gps.Point{Latitude: -5, Name: "SubPoint1"}
		var subPoint2 = gps.Point{Latitude: 8, Name: "SubPoint2"}
		var subPoint3 = gps.Point{Latitude: -10, Name: "SubPoint3"}
		var subPoint4 = gps.Point{Latitude: 40, Name: "SubPoint4"}
		var carSpeed = 10.0
		var droneSpeed = 20.0

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedItineraryInfo = mockitinerary.NewMockInfo(mockedCtrl)
			mockedSubRoute1 = mockroute.NewMockISubRoute(mockedCtrl)
			mockedSubRoute2 = mockroute.NewMockISubRoute(mockedCtrl)
			mockedMainStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop2 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop3 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop4 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedMainStop5 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedSubStop1 = mockroute.NewMockISubStop(mockedCtrl)
			mockedSubStop2 = mockroute.NewMockISubStop(mockedCtrl)
			mockedSubStop3 = mockroute.NewMockISubStop(mockedCtrl)
			mockedSubStop4 = mockroute.NewMockISubStop(mockedCtrl)
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

			mockedMainStop1.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{mockedSubRoute1})
			mockedMainStop2.EXPECT().StartingSubRoutes().Return(nil)
			mockedMainStop3.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{mockedSubRoute2})
			mockedMainStop4.EXPECT().StartingSubRoutes().Return(nil)
			mockedMainStop5.EXPECT().StartingSubRoutes().Return(nil)

			mockedMainStop1.EXPECT().ReturningSubRoutes().Return(nil)
			mockedMainStop2.EXPECT().ReturningSubRoutes().Return(nil)
			mockedMainStop3.EXPECT().ReturningSubRoutes().Return([]route.ISubRoute{mockedSubRoute1})
			mockedMainStop4.EXPECT().ReturningSubRoutes().Return(nil)
			mockedMainStop5.EXPECT().ReturningSubRoutes().Return([]route.ISubRoute{mockedSubRoute2})

			mockedSubRoute1.EXPECT().StartingStop().Return(mockedMainStop1)
			mockedSubRoute1.EXPECT().ReturningStop().Return(mockedMainStop3)
			mockedSubRoute2.EXPECT().StartingStop().Return(mockedMainStop3)
			mockedSubRoute2.EXPECT().ReturningStop().Return(mockedMainStop5)

			mockedSubRoute1.EXPECT().First().Return(mockedSubStop1)
			mockedSubRoute1.EXPECT().Iterator().Return(slc.NewIterator([]route.ISubStop{mockedSubStop1, mockedSubStop2}))
			mockedSubRoute2.EXPECT().First().Return(mockedSubStop3)
			mockedSubRoute2.EXPECT().Iterator().Return(slc.NewIterator([]route.ISubStop{mockedSubStop3, mockedSubStop4}))

			mockedItineraryInfo.EXPECT().RouteIterator().Return(slc.NewIterator([]route.IMainStop{mockedMainStop1, mockedMainStop2, mockedMainStop3, mockedMainStop4, mockedMainStop5}))
			mockedItineraryInfo.EXPECT().CarSpeed().Return(carSpeed)
			mockedItineraryInfo.EXPECT().DroneSpeed().Return(droneSpeed)

			secondSubRouteAddicionalTime := (90.0 / droneSpeed) - (10 / carSpeed)
			expectedTime := (20 / carSpeed) + secondSubRouteAddicionalTime
			receivedTime := TimeSpent(mockedItineraryInfo)
			Expect(receivedTime).To(Equal(expectedTime))
		})
	})
})

var _ = Describe("calcSubRouteTimes", func() {
	var mockedCtrl *gomock.Controller
	var mockedSubRoute1 *mockroute.MockISubRoute
	var mockedSubRoute2 *mockroute.MockISubRoute
	var mockedStartingStop *mockroute.MockIMainStop
	var mockedReturningStop *mockroute.MockIMainStop
	var mockedSubStop1 *mockroute.MockISubStop
	var mockedSubStop2 *mockroute.MockISubStop
	var mockedSubStop3 *mockroute.MockISubStop
	var mockedSubStop4 *mockroute.MockISubStop
	var startingPoint = gps.Point{Latitude: 0, Name: "StartingPoint1"}
	var returningPoint = gps.Point{Latitude: 20, Name: "ReturningPoint"}
	var point1 = gps.Point{Latitude: 5, Name: "Point1"}
	var point2 = gps.Point{Latitude: 25, Name: "Point2"}
	var point3 = gps.Point{Latitude: -10, Name: "Point3"}
	var point4 = gps.Point{Latitude: 10, Name: "Point4"}
	var droneSpeed = 10.0

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedSubRoute1 = mockroute.NewMockISubRoute(mockedCtrl)
		mockedSubRoute2 = mockroute.NewMockISubRoute(mockedCtrl)
		mockedStartingStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedReturningStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedSubStop1 = mockroute.NewMockISubStop(mockedCtrl)
		mockedSubStop2 = mockroute.NewMockISubStop(mockedCtrl)
		mockedSubStop3 = mockroute.NewMockISubStop(mockedCtrl)
		mockedSubStop4 = mockroute.NewMockISubStop(mockedCtrl)
	})

	Context("when subroute has points", func() {
		It("should return the total distance of the subroute", func() {
			mockedSubStop1.EXPECT().Point().Return(point1).AnyTimes()
			mockedSubStop2.EXPECT().Point().Return(point2).AnyTimes()
			mockedSubStop3.EXPECT().Point().Return(point3).AnyTimes()
			mockedSubStop4.EXPECT().Point().Return(point4).AnyTimes()

			mockedSubRoute1.EXPECT().First().Return(mockedSubStop1)
			mockedSubRoute1.EXPECT().Iterator().Return(slc.NewIterator([]route.ISubStop{mockedSubStop1, mockedSubStop2}))
			mockedSubRoute2.EXPECT().First().Return(mockedSubStop3)
			mockedSubRoute2.EXPECT().Iterator().Return(slc.NewIterator([]route.ISubStop{mockedSubStop3, mockedSubStop4}))

			mockedSubRoute1.EXPECT().StartingStop().Return(mockedStartingStop)
			mockedSubRoute2.EXPECT().StartingStop().Return(mockedStartingStop)
			mockedStartingStop.EXPECT().Point().Return(startingPoint).Times(2)
			mockedSubRoute1.EXPECT().ReturningStop().Return(mockedReturningStop)
			mockedSubRoute2.EXPECT().ReturningStop().Return(mockedReturningStop)
			mockedReturningStop.EXPECT().Point().Return(returningPoint).Times(2)

			subRoutes := []route.ISubRoute{mockedSubRoute1, mockedSubRoute2}
			subRouteFlyingTimes := make(subRouteTimes)
			mainRouteTravelTime := make(subRouteTimes)
			expectedSubRoutesFlyingTimes := subRouteTimes{
				mockedSubRoute1: 30.0 / droneSpeed, mockedSubRoute2: 40.0 / droneSpeed,
			}
			expectedMainRouteTravelTime := subRouteTimes{
				mockedSubRoute1: 0, mockedSubRoute2: 0,
			}
			calcSubRouteTimes(mainRouteTravelTime, subRouteFlyingTimes, subRoutes, droneSpeed)
			Expect(subRouteFlyingTimes).To(Equal(expectedSubRoutesFlyingTimes))
			Expect(mainRouteTravelTime).To(Equal(expectedMainRouteTravelTime))
		})
	})
})

var _ = Describe("maxAdditionalTimeWaitingSubRoutes", func() {
	var mockedCtrl *gomock.Controller
	var mockedSubRoute1 *mockroute.MockISubRoute
	var mockedSubRoute2 *mockroute.MockISubRoute

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedSubRoute1 = mockroute.NewMockISubRoute(mockedCtrl)
		mockedSubRoute2 = mockroute.NewMockISubRoute(mockedCtrl)
	})

	Context("when both received subroute are slower than main route", func() {
		It("should return the biggest waiting time", func() {
			subRouteFlyingTimes := subRouteTimes{
				mockedSubRoute1: 30.0, mockedSubRoute2: 40.0,
			}
			mainRouteTravelTime := subRouteTimes{
				mockedSubRoute1: 20.0, mockedSubRoute2: 25.0,
			}
			receivedAdditionalTime := maxAdditionalTimeWaitingSubRoutes(
				mainRouteTravelTime, subRouteFlyingTimes,
				[]route.ISubRoute{mockedSubRoute1, mockedSubRoute2},
			)
			Expect(receivedAdditionalTime).To(Equal(15.0))
		})
	})

	Context("when received subroute is slower than main route", func() {
		It("should return the biggest waiting time", func() {
			subRouteFlyingTimes := subRouteTimes{
				mockedSubRoute1: 10.0, mockedSubRoute2: 30.0,
			}
			mainRouteTravelTime := subRouteTimes{
				mockedSubRoute1: 20.0, mockedSubRoute2: 25.0,
			}
			receivedAdditionalTime := maxAdditionalTimeWaitingSubRoutes(
				mainRouteTravelTime, subRouteFlyingTimes,
				[]route.ISubRoute{mockedSubRoute2},
			)
			Expect(receivedAdditionalTime).To(Equal(5.0))
		})
	})

	Context("when both received subroutes are faster than main route", func() {
		It("should return zero", func() {
			subRouteFlyingTimes := subRouteTimes{
				mockedSubRoute1: 20.0, mockedSubRoute2: 30.0,
			}
			mainRouteTravelTime := subRouteTimes{
				mockedSubRoute1: 30.0, mockedSubRoute2: 40.0,
			}
			receivedAdditionalTime := maxAdditionalTimeWaitingSubRoutes(
				mainRouteTravelTime, subRouteFlyingTimes,
				[]route.ISubRoute{mockedSubRoute1, mockedSubRoute2},
			)
			Expect(receivedAdditionalTime).To(BeZero())
		})
	})

	Context("when received subroute is equal main route", func() {
		It("should return zero", func() {
			subRouteFlyingTimes := subRouteTimes{
				mockedSubRoute1: 20.0, mockedSubRoute2: 40.0,
			}
			mainRouteTravelTime := subRouteTimes{
				mockedSubRoute1: 30.0, mockedSubRoute2: 40.0,
			}
			receivedAdditionalTime := maxAdditionalTimeWaitingSubRoutes(
				mainRouteTravelTime, subRouteFlyingTimes,
				[]route.ISubRoute{mockedSubRoute1},
			)
			Expect(receivedAdditionalTime).To(BeZero())
		})
	})
})
