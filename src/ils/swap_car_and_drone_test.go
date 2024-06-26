package ils

import (
	"errors"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mock"
	mockroute "github.com/victorguarana/vehicle-routing/src/route/mock"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SwapCarAndDrone", func() {
	Context("when there is no swappable car stops", func() {
		var mockedCtrl *gomock.Controller
		var mockedFinder *mockitinerary.MockFinder
		var mockedModifier *mockitinerary.MockModifier

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedFinder = mockitinerary.NewMockFinder(mockedCtrl)
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
		})

		It("should return an error", func() {
			mockedFinder.EXPECT().FindWorstDroneStop().Return(itinerary.DroneStopCost{})
			mockedFinder.EXPECT().FindWorstSwappableCarStopsOrdered().Return([]itinerary.CarStopCost{})

			receivedErr := SwapCarAndDrone(mockedModifier, mockedFinder)
			Expect(receivedErr).To(HaveOccurred())
		})
	})

	Context("when no swappable car stop can be shifted to drone", func() {
		var mockedCtrl *gomock.Controller
		var mockedFinder *mockitinerary.MockFinder
		var mockedModifier *mockitinerary.MockModifier
		var mockedCarStop1 *mockroute.MockIMainStop
		var mockedCarStop2 *mockroute.MockIMainStop
		var mockedDroneStop *mockroute.MockISubStop
		var mockedFlight *mockroute.MockISubRoute

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedFinder = mockitinerary.NewMockFinder(mockedCtrl)
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCarStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCarStop2 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedDroneStop = mockroute.NewMockISubStop(mockedCtrl)
			mockedFlight = mockroute.NewMockISubRoute(mockedCtrl)
		})

		It("should return an error", func() {
			mockedCarStopCosts := []itinerary.CarStopCost{{Stop: mockedCarStop1, Index: 1}, {Stop: mockedCarStop2, Index: 2}}
			mockedDroneStopCost := itinerary.DroneStopCost{Stop: mockedDroneStop, Index: 1, Flight: mockedFlight}
			mockedFinder.EXPECT().FindWorstDroneStop().Return(mockedDroneStopCost)
			mockedFinder.EXPECT().FindWorstSwappableCarStopsOrdered().Return(mockedCarStopCosts)
			mockedModifier.EXPECT().RemoveDroneStopFromFlight(1, mockedFlight)
			mockedCarStop1.EXPECT().Point().Return(gps.Point{Latitude: 1})
			mockedCarStop2.EXPECT().Point().Return(gps.Point{Latitude: 2})
			mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(errors.New(""))
			mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 2}, gomock.Any()).Return(errors.New(""))

			receivedErr := SwapCarAndDrone(mockedModifier, mockedFinder)
			Expect(receivedErr).To(HaveOccurred())
		})
	})

	Context("when drone stop can not be shifted to car", func() {
		var mockedCtrl *gomock.Controller
		var mockedFinder *mockitinerary.MockFinder
		var mockedModifier *mockitinerary.MockModifier
		var mockedCarStop1 *mockroute.MockIMainStop
		var mockedCarStop2 *mockroute.MockIMainStop
		var mockedDroneStop *mockroute.MockISubStop
		var mockedFlight *mockroute.MockISubRoute

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedFinder = mockitinerary.NewMockFinder(mockedCtrl)
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCarStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCarStop2 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedDroneStop = mockroute.NewMockISubStop(mockedCtrl)
			mockedFlight = mockroute.NewMockISubRoute(mockedCtrl)
		})

		It("should return an error", func() {
			mockedCarStopCosts := []itinerary.CarStopCost{{Stop: mockedCarStop1, Index: 1}, {Stop: mockedCarStop2, Index: 2}}
			mockedDroneStopCost := itinerary.DroneStopCost{Stop: mockedDroneStop, Index: 1, Flight: mockedFlight}
			mockedFinder.EXPECT().FindWorstDroneStop().Return(mockedDroneStopCost)
			mockedFinder.EXPECT().FindWorstSwappableCarStopsOrdered().Return(mockedCarStopCosts)
			mockedModifier.EXPECT().RemoveDroneStopFromFlight(1, mockedFlight)
			mockedCarStop1.EXPECT().Point().Return(gps.Point{Latitude: 1})
			mockedCarStop2.EXPECT().Point().Return(gps.Point{Latitude: 2})
			mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(errors.New(""))
			mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 2}, gomock.Any()).Return(nil)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(2)
			mockedDroneStop.EXPECT().Point().Return(gps.Point{Latitude: 1})
			mockedModifier.EXPECT().InsertCarDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(errors.New(""))

			receivedErr := SwapCarAndDrone(mockedModifier, mockedFinder)
			Expect(receivedErr).To(HaveOccurred())
		})
	})

	Context("when a stop can be shifted", func() {
		var mockedCtrl *gomock.Controller
		var mockedFinder *mockitinerary.MockFinder
		var mockedModifier *mockitinerary.MockModifier
		var mockedCarStop1 *mockroute.MockIMainStop
		var mockedCarStop2 *mockroute.MockIMainStop
		var mockedDroneStop *mockroute.MockISubStop
		var mockedFlight *mockroute.MockISubRoute

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedFinder = mockitinerary.NewMockFinder(mockedCtrl)
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCarStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCarStop2 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedDroneStop = mockroute.NewMockISubStop(mockedCtrl)
			mockedFlight = mockroute.NewMockISubRoute(mockedCtrl)
		})

		It("should shift the car stop to drone", func() {
			mockedCarStopCosts := []itinerary.CarStopCost{{Stop: mockedCarStop1, Index: 1}, {Stop: mockedCarStop2, Index: 2}}
			mockedDroneStopCost := itinerary.DroneStopCost{Stop: mockedDroneStop, Index: 1, Flight: mockedFlight}
			mockedFinder.EXPECT().FindWorstDroneStop().Return(mockedDroneStopCost)
			mockedFinder.EXPECT().FindWorstSwappableCarStopsOrdered().Return(mockedCarStopCosts)
			mockedModifier.EXPECT().RemoveDroneStopFromFlight(1, mockedFlight)
			mockedCarStop1.EXPECT().Point().Return(gps.Point{Latitude: 1})
			mockedCarStop2.EXPECT().Point().Return(gps.Point{Latitude: 2})
			mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(errors.New(""))
			mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 2}, gomock.Any()).Return(nil)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(2)
			mockedDroneStop.EXPECT().Point().Return(gps.Point{Latitude: 1})
			mockedModifier.EXPECT().InsertCarDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(nil)

			receivedErr := SwapCarAndDrone(mockedModifier, mockedFinder)
			Expect(receivedErr).NotTo(HaveOccurred())
		})
	})
})
