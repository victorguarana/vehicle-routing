package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
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

	Describe("CarSupport", func() {
		var sut info
		var mockedCtrl *gomock.Controller
		var mockedCar *mockvehicle.MockICar
		var nextPoints = []gps.Point{
			{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination1"},
			{Latitude: 7, Longitude: 8, PackageSize: 9, Name: "destination2"},
		}

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

		It("should return true if the car supports the route", func() {
			mockedCar.EXPECT().Support(nextPoints).Return(true)
			Expect(sut.CarSupport(nextPoints...)).To(BeTrue())
		})

		It("should return false if the car does not support the route", func() {
			mockedCar.EXPECT().Support(nextPoints).Return(false)
			Expect(sut.CarSupport(nextPoints...)).To(BeFalse())
		})
	})

	Describe("DroneCanReach", func() {
		var sut info
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var nextPoints = []gps.Point{
			{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination1"},
			{Latitude: 7, Longitude: 8, PackageSize: 9, Name: "destination2"},
		}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)

			sut = info{&itinerary{}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone can reach the route", func() {
			mockedDrone1.EXPECT().CanReach(nextPoints).Return(true)
			Expect(sut.DroneCanReach(mockedDrone1, nextPoints...)).To(BeTrue())
		})

		It("should return false if the drone can not reach the route", func() {
			mockedDrone2.EXPECT().CanReach(nextPoints).Return(false)
			Expect(sut.DroneCanReach(mockedDrone2, nextPoints...)).To(BeFalse())
		})
	})

	Describe("Drones", func() {
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCar *mockvehicle.MockICar

		var sut info

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			sut = info{&itinerary{
				car: mockedCar,
			}}
		})

		It("should return all drone numbers", func() {
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			receivedDroneNumbers := sut.Drones()
			Expect(receivedDroneNumbers).To(HaveLen(2))
			Expect(receivedDroneNumbers).To(ContainElements(mockedDrone1, mockedDrone2))
		})
	})

	Describe("DroneIsFlying", func() {
		var sut info
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)

			sut = info{&itinerary{}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone is flying", func() {
			mockedDrone1.EXPECT().IsFlying().Return(true)
			Expect(sut.DroneIsFlying(mockedDrone1)).To(BeTrue())
		})

		It("should return false if the drone is not flying", func() {
			mockedDrone2.EXPECT().IsFlying().Return(false)
			Expect(sut.DroneIsFlying(mockedDrone2)).To(BeFalse())
		})
	})

	Describe("DroneSupport", func() {
		var sut info
		var mockedCtrl *gomock.Controller
		var mockedDrone *mockvehicle.MockIDrone
		var deliveryPoint = gps.Point{Latitude: 4, Longitude: 5, PackageSize: 6, Name: "destination1"}
		var landingPoint = gps.Point{Latitude: 7, Longitude: 8, PackageSize: 9, Name: "destination2"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone = mockvehicle.NewMockIDrone(mockedCtrl)

			sut = info{&itinerary{}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		Context("when can delivery point and land at the next", func() {
			It("should return true", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(true)
				mockedDrone.EXPECT().CanReach(deliveryPoint, landingPoint).Return(true)
				Expect(sut.DroneSupport(mockedDrone, deliveryPoint, landingPoint)).To(BeTrue())
			})
		})

		Context("when can delivery point but can not reach at the next", func() {
			It("should return false", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(true)
				mockedDrone.EXPECT().CanReach(deliveryPoint, landingPoint).Return(false)
				Expect(sut.DroneSupport(mockedDrone, deliveryPoint, landingPoint)).To(BeFalse())
			})
		})

		Context("when can not delivery point", func() {
			It("should return false", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(false)
				Expect(sut.DroneSupport(mockedDrone, deliveryPoint, landingPoint)).To(BeFalse())
			})
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
