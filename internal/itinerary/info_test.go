package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("info{}", func() {
	Describe("ActualCarPoint", func() {
		var sut info
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicle.MockICar
		var initialPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			sut = info{&itinerary{
				car: mockedCar,
			}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return the last point of the route", func() {
			mockedCar.EXPECT().ActualPoint().Return(initialPoint)
			Expect(sut.ActualCarPoint()).To(Equal(initialPoint))
		})
	})

	Describe("Car", func() {
		var sut info
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicle.MockICar

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)

			sut = info{&itinerary{
				car: mockedCar,
			}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return the current car", func() {
			Expect(sut.Car()).To(Equal(mockedCar))
		})
	})

	Describe("RouteIterator", func() {
		var sut info
		var mockedCtrl *gomock.Controller
		var mockedRoute *mockroute.MockIMainRoute
		var mockedMainStop1 *mockroute.MockIMainStop
		var mockedMainStop2 *mockroute.MockIMainStop
		var mockedMainStops = []route.IMainStop{mockedMainStop1, mockedMainStop2}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedRoute = mockroute.NewMockIMainRoute(mockedCtrl)

			sut = info{&itinerary{
				route: mockedRoute,
			}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return an iterator for the route", func() {
			expectedIterator := slc.NewIterator[route.IMainStop](mockedMainStops)
			mockedRoute.EXPECT().Iterator().Return(expectedIterator)
			Expect(sut.RouteIterator()).To(Equal(expectedIterator))
		})
	})
})
