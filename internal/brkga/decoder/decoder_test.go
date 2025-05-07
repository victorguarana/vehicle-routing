package decoder

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"
	"go.uber.org/mock/gomock"
)

var _ = Describe("decoder", func() {
	var mockedCtrl *gomock.Controller
	var mockedCar1 *mockvehicle.MockICar
	var mockedCar2 *mockvehicle.MockICar
	var mockedItinerary1 *mockitinerary.MockItinerary
	var mockedItinerary2 *mockitinerary.MockItinerary
	var mockedConstructor1 *mockitinerary.MockConstructor
	var mockedConstructor2 *mockitinerary.MockConstructor

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicle.NewMockICar(mockedCtrl)
		mockedCar2 = mockvehicle.NewMockICar(mockedCtrl)
		mockedItinerary1 = mockitinerary.NewMockItinerary(mockedCtrl)
		mockedItinerary2 = mockitinerary.NewMockItinerary(mockedCtrl)
		mockedConstructor1 = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedConstructor2 = mockitinerary.NewMockConstructor(mockedCtrl)
	})

	Describe("finalizeItineraries", func() {
		var warehouse1 gps.Point
		var warehouse2 gps.Point
		var carPoint1 gps.Point
		var carPoint2 gps.Point

		var mockedCarStop1 *mockroute.MockIMainStop
		var mockedCarStop2 *mockroute.MockIMainStop

		BeforeEach(func() {
			warehouse1 = gps.Point{Latitude: 1, Longitude: 1}
			warehouse2 = gps.Point{Latitude: 10, Longitude: 10}
			carPoint1 = gps.Point{Latitude: 0, Longitude: 0}
			carPoint2 = gps.Point{Latitude: 15, Longitude: 15}

			mockedCarStop1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCarStop2 = mockroute.NewMockIMainStop(mockedCtrl)
		})

		It("should move cars to closest warehouses", func() {
			itineraryList := []itinerary.Itinerary{mockedItinerary1, mockedItinerary2}
			gpsMap := gps.Map{
				Warehouses: []gps.Point{warehouse1, warehouse2},
			}

			mockedItinerary1.EXPECT().Constructor().Return(mockedConstructor1).AnyTimes()
			mockedItinerary2.EXPECT().Constructor().Return(mockedConstructor2).AnyTimes()

			mockedConstructor1.EXPECT().MoveCar(warehouse1)
			mockedConstructor1.EXPECT().ActualCarPoint().Return(carPoint1)
			mockedConstructor1.EXPECT().ActualCarStop().Return(mockedCarStop1)
			mockedConstructor1.EXPECT().LandAllDrones(mockedCarStop1)

			mockedConstructor2.EXPECT().MoveCar(warehouse2)
			mockedConstructor2.EXPECT().ActualCarPoint().Return(carPoint2)
			mockedConstructor2.EXPECT().ActualCarStop().Return(mockedCarStop2)
			mockedConstructor2.EXPECT().LandAllDrones(mockedCarStop2)

			finalizeItineraries(itineraryList, gpsMap)
		})
	})

	Describe("cloneCars", func() {
		var clonedCar1 *mockvehicle.MockICar
		var clonedCar2 *mockvehicle.MockICar

		BeforeEach(func() {
			clonedCar1 = mockvehicle.NewMockICar(mockedCtrl)
			clonedCar2 = mockvehicle.NewMockICar(mockedCtrl)
		})

		It("should add cloned cars to cars list", func() {
			carList := []vehicle.ICar{mockedCar1, mockedCar2}
			mockedCar1.EXPECT().Clone().Return(clonedCar1)
			mockedCar2.EXPECT().Clone().Return(clonedCar2)

			receivedCarList := cloneCars(carList)

			Expect(receivedCarList).To(HaveExactElements(BeIdenticalTo(clonedCar1), BeIdenticalTo(clonedCar2)))
		})
	})

	Describe("isValidSolution", func() {
		var mockedValidator1 *mockitinerary.MockValidator
		var mockedValidator2 *mockitinerary.MockValidator

		BeforeEach(func() {
			mockedValidator1 = mockitinerary.NewMockValidator(mockedCtrl)
			mockedValidator2 = mockitinerary.NewMockValidator(mockedCtrl)
		})

		It("should return false when any itinerary is invalid", func() {
			itineraryList := []itinerary.Itinerary{mockedItinerary1, mockedItinerary2}

			mockedItinerary1.EXPECT().Validator().Return(mockedValidator1)
			mockedValidator1.EXPECT().IsValid().Return(true)
			mockedItinerary2.EXPECT().Validator().Return(mockedValidator2)
			mockedValidator2.EXPECT().IsValid().Return(false)

			Expect(isValidSolution(itineraryList)).To(BeFalse())
		})

		It("should return true when all itineraries are valid", func() {
			itineraryList := []itinerary.Itinerary{mockedItinerary1, mockedItinerary2}

			mockedItinerary1.EXPECT().Validator().Return(mockedValidator1)
			mockedValidator1.EXPECT().IsValid().Return(true)
			mockedItinerary2.EXPECT().Validator().Return(mockedValidator2)
			mockedValidator2.EXPECT().IsValid().Return(true)

			Expect(isValidSolution(itineraryList)).To(BeTrue())
		})
	})

})
