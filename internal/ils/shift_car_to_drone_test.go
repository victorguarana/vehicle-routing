package ils

import (
	"errors"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShiftCarToDrone", func() {
	Context("when there is swappable car stops", func() {
		var mockedCtrl *gomock.Controller
		var mockedFinder *mockitinerary.MockFinder
		var mockedModifier *mockitinerary.MockModifier
		var mockedCarStop1 *mockroute.MockIMainStop
		var mockedCarStop2 *mockroute.MockIMainStop

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedFinder = mockitinerary.NewMockFinder(mockedCtrl)
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCarStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCarStop2 = mockroute.NewMockIMainStop(mockedCtrl)
		})

		Context("when a car stop can be shifted to drone", func() {
			It("should shift the car stop to drone", func() {
				mockedCarStopCosts := []itinerary.CarStopCost{{Stop: mockedCarStop1, Index: 1}, {Stop: mockedCarStop2, Index: 2}}
				mockedCarStop1.EXPECT().Point().Return(gps.Point{Latitude: 1})
				mockedCarStop2.EXPECT().Point().Return(gps.Point{Latitude: 2})
				mockedFinder.EXPECT().FindWorstSwappableCarStopsOrdered().Return(mockedCarStopCosts)
				mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(errors.New(""))
				mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 2}, gomock.Any()).Return(nil)
				mockedModifier.EXPECT().RemoveMainStopFromRoute(2)

				receivedErr := ShiftCarToDrone(mockedModifier, mockedFinder)
				Expect(receivedErr).NotTo(HaveOccurred())
			})

			Context("when no car stop can be shifted to drone", func() {
				It("should return an error", func() {
					mockedCarStopCosts := []itinerary.CarStopCost{{Stop: mockedCarStop1, Index: 1}, {Stop: mockedCarStop2, Index: 2}}
					mockedCarStop1.EXPECT().Point().Return(gps.Point{Latitude: 1})
					mockedCarStop2.EXPECT().Point().Return(gps.Point{Latitude: 2})
					mockedFinder.EXPECT().FindWorstSwappableCarStopsOrdered().Return(mockedCarStopCosts)
					mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(errors.New(""))
					mockedModifier.EXPECT().InsertDroneDelivery(gps.Point{Latitude: 2}, gomock.Any()).Return(errors.New(""))

					receivedErr := ShiftCarToDrone(mockedModifier, mockedFinder)
					Expect(receivedErr).To(HaveOccurred())
				})
			})
		})
	})

	Context("when there is no swappable car stops", func() {
		var mockedCtrl *gomock.Controller
		var mockedFinder *mockitinerary.MockFinder

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedFinder = mockitinerary.NewMockFinder(mockedCtrl)
		})

		It("should return an error", func() {
			mockedFinder.EXPECT().FindWorstSwappableCarStopsOrdered().Return([]itinerary.CarStopCost{})

			receivedErr := ShiftCarToDrone(nil, mockedFinder)
			Expect(receivedErr).To(HaveOccurred())
		})
	})
})
