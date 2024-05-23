package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
	mockroute "github.com/victorguarana/vehicle-routing/src/route/mock"
	"github.com/victorguarana/vehicle-routing/src/slc"
	"github.com/victorguarana/vehicle-routing/src/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/src/vehicle/mock"

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

			sut = info{&itinerary{
				droneNumbersMap: map[DroneNumber]vehicle.IDrone{
					1: mockedDrone1,
					2: mockedDrone2,
				},
			}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone can reach the route", func() {
			mockedDrone1.EXPECT().CanReach(nextPoints).Return(true)
			Expect(sut.DroneCanReach(1, nextPoints...)).To(BeTrue())
		})

		It("should return false if the drone can not reach the route", func() {
			mockedDrone2.EXPECT().CanReach(nextPoints).Return(false)
			Expect(sut.DroneCanReach(2, nextPoints...)).To(BeFalse())
		})
	})

	Describe("DroneNumbers", func() {
		var sut = info{&itinerary{
			droneNumbersMap: map[DroneNumber]vehicle.IDrone{
				1: nil,
				2: nil,
			},
		}}

		It("should return all drone numbers", func() {
			receivedDroneNumbers := sut.DroneNumbers()
			Expect(receivedDroneNumbers).To(HaveLen(2))
			Expect(receivedDroneNumbers).To(ContainElements(DroneNumber(1), DroneNumber(2)))
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

			sut = info{&itinerary{
				droneNumbersMap: map[DroneNumber]vehicle.IDrone{
					1: mockedDrone1,
					2: mockedDrone2,
				},
			}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if the drone is flying", func() {
			mockedDrone1.EXPECT().IsFlying().Return(true)
			Expect(sut.DroneIsFlying(1)).To(BeTrue())
		})

		It("should return false if the drone is not flying", func() {
			mockedDrone2.EXPECT().IsFlying().Return(false)
			Expect(sut.DroneIsFlying(2)).To(BeFalse())
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

			sut = info{&itinerary{
				droneNumbersMap: map[DroneNumber]vehicle.IDrone{
					1: mockedDrone,
				},
			}}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		Context("when can delivery point and land at the next", func() {
			It("should return true", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(true)
				mockedDrone.EXPECT().CanReach(deliveryPoint, landingPoint).Return(true)
				Expect(sut.DroneSupport(1, deliveryPoint, landingPoint)).To(BeTrue())
			})
		})

		Context("when can delivery point but can not reach at the next", func() {
			It("should return false", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(true)
				mockedDrone.EXPECT().CanReach(deliveryPoint, landingPoint).Return(false)
				Expect(sut.DroneSupport(1, deliveryPoint, landingPoint)).To(BeFalse())
			})
		})

		Context("when can not delivery point", func() {
			It("should return false", func() {
				mockedDrone.EXPECT().Support(deliveryPoint).Return(false)
				Expect(sut.DroneSupport(1, deliveryPoint, landingPoint)).To(BeFalse())
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
