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

var _ = Describe("ShiftDroneToCar", func() {
	var mockedCtrl *gomock.Controller
	var mockedFinder *mockitinerary.MockFinder
	var mockedModifier *mockitinerary.MockModifier
	var mockedDroneStop *mockroute.MockISubStop
	var mockedFlight *mockroute.MockISubRoute

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedFinder = mockitinerary.NewMockFinder(mockedCtrl)
		mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
		mockedDroneStop = mockroute.NewMockISubStop(mockedCtrl)
		mockedFlight = mockroute.NewMockISubRoute(mockedCtrl)
	})

	Context("when a drone stop can be shifted to car", func() {
		It("should shift the drone stop to car", func() {
			mockedDroneStopCost := itinerary.DroneStopCost{Stop: mockedDroneStop, Index: 1, Flight: mockedFlight}
			mockedDroneStop.EXPECT().Point().Return(gps.Point{Latitude: 1})
			mockedFinder.EXPECT().FindWorstDroneStop().Return(mockedDroneStopCost)
			mockedModifier.EXPECT().InsertCarDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(nil)
			mockedModifier.EXPECT().RemoveDroneStopFromFlight(1, mockedFlight)

			receivedErr := ShiftDroneToCar(mockedModifier, mockedFinder)
			Expect(receivedErr).NotTo(HaveOccurred())
		})
	})

	Context("when no drone stop can be shifted to car", func() {
		It("should return an error", func() {
			mockedDroneStopCost := itinerary.DroneStopCost{Stop: mockedDroneStop, Index: 1, Flight: mockedFlight}
			mockedDroneStop.EXPECT().Point().Return(gps.Point{Latitude: 1})
			mockedFinder.EXPECT().FindWorstDroneStop().Return(mockedDroneStopCost)
			mockedModifier.EXPECT().InsertCarDelivery(gps.Point{Latitude: 1}, gomock.Any()).Return(errors.New(""))

			receivedErr := ShiftDroneToCar(mockedModifier, mockedFinder)
			Expect(receivedErr).To(HaveOccurred())
		})
	})
})
