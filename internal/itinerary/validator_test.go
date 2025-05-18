package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("validator", func() {
	var sut validator
	var mockedCtrl *gomock.Controller

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
	})

	Describe("carCanSupportRoute", func() {
		var mockedRoute *mockroute.MockIMainRoute
		var mockedCar *mockvehicle.MockICar
		var mockedWarehouse1 *mockroute.MockIMainStop
		var mockedCustomer1 *mockroute.MockIMainStop
		var mockedCustomer2 *mockroute.MockIMainStop
		var mockedCustomer3 *mockroute.MockIMainStop
		var stubWarehousePoint1 = gps.Point{Latitude: 5}
		var stubCustomerPoint1 = gps.Point{Latitude: 0, PackageSize: 1}
		var stubCustomerPoint2 = gps.Point{Latitude: 10, PackageSize: 1}
		var stubCustomerPoint3 = gps.Point{Latitude: 30, PackageSize: 5}

		BeforeEach(func() {
			mockedRoute = mockroute.NewMockIMainRoute(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedWarehouse1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCustomer1 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCustomer2 = mockroute.NewMockIMainStop(mockedCtrl)
			mockedCustomer3 = mockroute.NewMockIMainStop(mockedCtrl)

			sut = validator{
				info: &info{
					itinerary: &itinerary{
						route: mockedRoute,
						car:   mockedCar,
					},
				},
			}
		})

		Context("when car can supports entire route", func() {
			It("should return true", func() {
				mockedRoute.EXPECT().MainStopList().Return([]route.IMainStop{
					mockedCustomer1, mockedCustomer2, mockedWarehouse1, mockedCustomer3,
				})
				mockedCustomer1.EXPECT().Point().Return(stubCustomerPoint1).AnyTimes()
				mockedCustomer1.EXPECT().IsWarehouse().Return(false)
				mockedCustomer1.EXPECT().IsCustomer().Return(true)
				mockedCustomer1.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCustomer2.EXPECT().Point().Return(stubCustomerPoint2).AnyTimes()
				mockedCustomer2.EXPECT().IsWarehouse().Return(false)
				mockedCustomer2.EXPECT().IsCustomer().Return(true)
				mockedCustomer2.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedWarehouse1.EXPECT().Point().Return(stubWarehousePoint1).AnyTimes()
				mockedWarehouse1.EXPECT().IsWarehouse().Return(true)
				mockedWarehouse1.EXPECT().IsCustomer().Return(false)

				mockedCustomer3.EXPECT().Point().Return(stubCustomerPoint3).AnyTimes()
				mockedCustomer3.EXPECT().IsWarehouse().Return(false)
				mockedCustomer3.EXPECT().IsCustomer().Return(true)
				mockedCustomer3.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCar.EXPECT().Range().Return(100.0).AnyTimes()
				mockedCar.EXPECT().Storage().Return(10.0).AnyTimes()

				Expect(sut.carCanSupportRoute()).To(BeTrue())
			})
		})

		Context("when car is out of range for first part of route", func() {
			It("should return false", func() {
				mockedRoute.EXPECT().MainStopList().Return([]route.IMainStop{
					mockedCustomer1, mockedCustomer2, mockedWarehouse1, mockedCustomer3,
				})
				mockedCustomer1.EXPECT().IsWarehouse().Return(false)
				mockedCustomer1.EXPECT().Point().Return(stubCustomerPoint1).AnyTimes()
				mockedCustomer1.EXPECT().IsCustomer().Return(true)
				mockedCustomer1.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCustomer2.EXPECT().IsWarehouse().Return(false)
				mockedCustomer2.EXPECT().IsCustomer().Return(true)
				mockedCustomer2.EXPECT().Point().Return(stubCustomerPoint2).AnyTimes()
				mockedCustomer2.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedWarehouse1.EXPECT().Point().Return(stubWarehousePoint1).AnyTimes()

				mockedCar.EXPECT().Range().Return(10.0).AnyTimes()
				mockedCar.EXPECT().Storage().Return(100.0).AnyTimes()

				Expect(sut.carCanSupportRoute()).To(BeFalse())
			})
		})

		Context("when car is out of storage for first part of route", func() {
			It("should return false", func() {
				mockedRoute.EXPECT().MainStopList().Return([]route.IMainStop{
					mockedCustomer1, mockedCustomer2, mockedWarehouse1, mockedCustomer3,
				})
				mockedCustomer1.EXPECT().IsCustomer().Return(true)
				mockedCustomer1.EXPECT().IsWarehouse().Return(false)
				mockedCustomer1.EXPECT().Point().Return(stubCustomerPoint1).AnyTimes()
				mockedCustomer1.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCustomer2.EXPECT().IsWarehouse().Return(false)
				mockedCustomer2.EXPECT().IsCustomer().Return(true)
				mockedCustomer2.EXPECT().Point().Return(stubCustomerPoint2).AnyTimes()
				mockedCustomer2.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCar.EXPECT().Range().Return(100.0).AnyTimes()
				mockedCar.EXPECT().Storage().Return(1.0).AnyTimes()

				Expect(sut.carCanSupportRoute()).To(BeFalse())
			})
		})

		Context("when car is out of range for second part of route", func() {
			It("should return false", func() {
				mockedRoute.EXPECT().MainStopList().Return([]route.IMainStop{
					mockedCustomer1, mockedCustomer2, mockedWarehouse1, mockedCustomer3,
				})
				mockedCustomer1.EXPECT().Point().Return(stubCustomerPoint1).AnyTimes()
				mockedCustomer1.EXPECT().IsWarehouse().Return(false)
				mockedCustomer1.EXPECT().IsCustomer().Return(true)
				mockedCustomer1.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCustomer2.EXPECT().Point().Return(stubCustomerPoint2).AnyTimes()
				mockedCustomer2.EXPECT().IsWarehouse().Return(false)
				mockedCustomer2.EXPECT().IsCustomer().Return(true)
				mockedCustomer2.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedWarehouse1.EXPECT().Point().Return(stubWarehousePoint1).AnyTimes()
				mockedWarehouse1.EXPECT().IsWarehouse().Return(true)
				mockedWarehouse1.EXPECT().IsCustomer().Return(false)

				mockedCustomer3.EXPECT().Point().Return(stubCustomerPoint3).AnyTimes()

				mockedCar.EXPECT().Range().Return(20.0).AnyTimes()
				mockedCar.EXPECT().Storage().Return(100.0).AnyTimes()

				Expect(sut.carCanSupportRoute()).To(BeFalse())
			})
		})

		Context("when car is out of storage for second part of route", func() {
			It("should return false", func() {
				mockedRoute.EXPECT().MainStopList().Return([]route.IMainStop{
					mockedCustomer1, mockedCustomer2, mockedWarehouse1, mockedCustomer3,
				})
				mockedCustomer1.EXPECT().Point().Return(stubCustomerPoint1).AnyTimes()
				mockedCustomer1.EXPECT().IsWarehouse().Return(false)
				mockedCustomer1.EXPECT().IsCustomer().Return(true)
				mockedCustomer1.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCustomer2.EXPECT().Point().Return(stubCustomerPoint2).AnyTimes()
				mockedCustomer2.EXPECT().IsWarehouse().Return(false)
				mockedCustomer2.EXPECT().IsCustomer().Return(true)
				mockedCustomer2.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedWarehouse1.EXPECT().Point().Return(stubWarehousePoint1).AnyTimes()
				mockedWarehouse1.EXPECT().IsWarehouse().Return(true)
				mockedWarehouse1.EXPECT().IsCustomer().Return(false)

				mockedCustomer3.EXPECT().Point().Return(stubCustomerPoint3).AnyTimes()
				mockedCustomer3.EXPECT().IsWarehouse().Return(false)
				mockedCustomer3.EXPECT().IsCustomer().Return(true)
				mockedCustomer3.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})

				mockedCar.EXPECT().Range().Return(100.0).AnyTimes()
				mockedCar.EXPECT().Storage().Return(4.0).AnyTimes()

				Expect(sut.carCanSupportRoute()).To(BeFalse())
			})
		})
	})

	Describe("calcMainStopRequiredStorage", func() {
		var mockedMainStop *mockroute.MockIMainStop
		var mockedFlight *mockroute.MockISubRoute
		var mockedStartingPoint *mockroute.MockIMainStop
		var mockedReturningPoint *mockroute.MockIMainStop
		var mockedSubStop1 *mockroute.MockISubStop
		var mockedSubStop2 *mockroute.MockISubStop
		var stubStartingPoint = gps.Point{PackageSize: 1}
		var stubReturningPoint = gps.Point{PackageSize: 2}
		var stubSubStopPoint1 = gps.Point{PackageSize: 4}
		var stubSubStopPoint2 = gps.Point{PackageSize: 8}
		var stubMainStopPoint = gps.Point{PackageSize: 16}

		BeforeEach(func() {
			mockedMainStop = mockroute.NewMockIMainStop(mockedCtrl)
			mockedFlight = mockroute.NewMockISubRoute(mockedCtrl)
			mockedStartingPoint = mockroute.NewMockIMainStop(mockedCtrl)
			mockedReturningPoint = mockroute.NewMockIMainStop(mockedCtrl)
			mockedSubStop1 = mockroute.NewMockISubStop(mockedCtrl)
			mockedSubStop2 = mockroute.NewMockISubStop(mockedCtrl)
		})

		Context("when stop is not starting point for flights", func() {
			It("should return just actual stop storage", func() {
				mockedMainStop.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{})
				mockedMainStop.EXPECT().Point().Return(stubMainStopPoint)

				Expect(sut.calcMainStopRequiredStorage(mockedMainStop)).To(Equal(16.0))
			})
		})

		Context("when stop is starting point for flights", func() {
			It("should return actual stop and flight storage", func() {
				mockedMainStop.EXPECT().StartingSubRoutes().Return([]route.ISubRoute{mockedFlight})
				mockedMainStop.EXPECT().Point().Return(stubMainStopPoint)

				mockedFlight.EXPECT().StartingStop().Return(mockedStartingPoint)
				mockedFlight.EXPECT().SubStopList().Return([]route.ISubStop{mockedSubStop1, mockedSubStop2})
				mockedFlight.EXPECT().ReturningStop().Return(mockedReturningPoint)
				mockedStartingPoint.EXPECT().Point().Return(stubStartingPoint)
				mockedSubStop1.EXPECT().Point().Return(stubSubStopPoint1)
				mockedSubStop2.EXPECT().Point().Return(stubSubStopPoint2)
				mockedReturningPoint.EXPECT().Point().Return(stubReturningPoint)

				Expect(sut.calcMainStopRequiredStorage(mockedMainStop)).To(Equal(28.0))
			})
		})
	})

	Describe("calcFlightRequirements", func() {
		var mockedFlight *mockroute.MockISubRoute
		var mockedStartingPoint *mockroute.MockIMainStop
		var mockedReturningPoint *mockroute.MockIMainStop
		var mockedSubStop1 *mockroute.MockISubStop
		var mockedSubStop2 *mockroute.MockISubStop
		var stubStartingPoint = gps.Point{Latitude: 0, Longitude: 0, PackageSize: 1}
		var stubReturningPoint = gps.Point{Latitude: 15, Longitude: 0, PackageSize: 2}
		var stubSubStopPoint1 = gps.Point{Latitude: 10, Longitude: 0, PackageSize: 4}
		var stubSubStopPoint2 = gps.Point{Latitude: 5, Longitude: 0, PackageSize: 8}

		BeforeEach(func() {
			mockedFlight = mockroute.NewMockISubRoute(mockedCtrl)
			mockedStartingPoint = mockroute.NewMockIMainStop(mockedCtrl)
			mockedReturningPoint = mockroute.NewMockIMainStop(mockedCtrl)
			mockedSubStop1 = mockroute.NewMockISubStop(mockedCtrl)
			mockedSubStop2 = mockroute.NewMockISubStop(mockedCtrl)

			sut = validator{}
		})

		It("should return correct requirements", func() {
			expectedRequirements := requirements{
				requiredStorage: 12.0,
				requiredRange:   25,
			}

			mockedFlight.EXPECT().StartingStop().Return(mockedStartingPoint)
			mockedFlight.EXPECT().SubStopList().Return([]route.ISubStop{mockedSubStop1, mockedSubStop2})
			mockedFlight.EXPECT().ReturningStop().Return(mockedReturningPoint)
			mockedStartingPoint.EXPECT().Point().Return(stubStartingPoint)
			mockedSubStop1.EXPECT().Point().Return(stubSubStopPoint1)
			mockedSubStop2.EXPECT().Point().Return(stubSubStopPoint2)
			mockedReturningPoint.EXPECT().Point().Return(stubReturningPoint)

			Expect(sut.calcFlightRequirements(mockedFlight)).To(Equal(expectedRequirements))
		})
	})

	Describe("dronesCanSupportFlighs", func() {
		var mockedFlight *mockroute.MockISubRoute
		var mockedStartingPoint *mockroute.MockIMainStop
		var mockedReturningPoint *mockroute.MockIMainStop
		var mockedSubStop1 *mockroute.MockISubStop
		var mockedSubStop2 *mockroute.MockISubStop
		var stubStartingPoint = gps.Point{Latitude: 0, Longitude: 0, PackageSize: 1}
		var stubReturningPoint = gps.Point{Latitude: 15, Longitude: 0, PackageSize: 2}
		var stubSubStopPoint1 = gps.Point{Latitude: 10, Longitude: 0, PackageSize: 4}
		var stubSubStopPoint2 = gps.Point{Latitude: 5, Longitude: 0, PackageSize: 8}
		var mockedDrone *mockvehicle.MockIDrone
		var subItinerary SubItinerary

		BeforeEach(func() {
			mockedFlight = mockroute.NewMockISubRoute(mockedCtrl)
			mockedStartingPoint = mockroute.NewMockIMainStop(mockedCtrl)
			mockedReturningPoint = mockroute.NewMockIMainStop(mockedCtrl)
			mockedSubStop1 = mockroute.NewMockISubStop(mockedCtrl)
			mockedSubStop2 = mockroute.NewMockISubStop(mockedCtrl)
			mockedDrone = mockvehicle.NewMockIDrone(mockedCtrl)

			subItinerary = SubItinerary{
				Drone:  mockedDrone,
				Flight: mockedFlight,
			}

			sut = validator{
				info: &info{
					itinerary: &itinerary{
						completedSubItineraryList: []SubItinerary{subItinerary},
					},
				},
			}

			mockedFlight.EXPECT().StartingStop().Return(mockedStartingPoint)
			mockedFlight.EXPECT().SubStopList().Return([]route.ISubStop{mockedSubStop1, mockedSubStop2})
			mockedFlight.EXPECT().ReturningStop().Return(mockedReturningPoint)
			mockedStartingPoint.EXPECT().Point().Return(stubStartingPoint)
			mockedSubStop1.EXPECT().Point().Return(stubSubStopPoint1)
			mockedSubStop2.EXPECT().Point().Return(stubSubStopPoint2)
			mockedReturningPoint.EXPECT().Point().Return(stubReturningPoint)
		})

		Context("when drone supports its flight", func() {
			It("should return correct requirements", func() {
				mockedDrone.EXPECT().Range().Return(25.0)
				mockedDrone.EXPECT().Storage().Return(12.0)
				Expect(sut.dronesCanSupportFlights()).To(BeTrue())
			})
		})

		Context("when drone does not supports its flight range", func() {
			It("should return correct requirements", func() {
				mockedDrone.EXPECT().Range().Return(24.0)
				Expect(sut.dronesCanSupportFlights()).To(BeFalse())
			})
		})

		Context("when drone does not supports its flight storage", func() {
			It("should return correct requirements", func() {
				mockedDrone.EXPECT().Range().Return(25.0)
				mockedDrone.EXPECT().Storage().Return(11.0)
				Expect(sut.dronesCanSupportFlights()).To(BeFalse())
			})
		})
	})
})
