package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	mockRoutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"
	"github.com/victorguarana/go-vehicle-route/src/slc"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
	mockVehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("initDroneStrikes", func() {
	var mockedCtrl *gomock.Controller
	var mockedCar = mockVehicles.NewMockICar(mockedCtrl)
	var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockVehicles.NewMockICar(mockedCtrl)
		mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should initialize drone strikes", func() {
		expectedDronesStrikes := []droneStrikes{
			{drone: mockedDrone1},
			{drone: mockedDrone2},
		}
		mockedCar.EXPECT().Drones().Return([]vehicles.IDrone{mockedDrone1, mockedDrone2})
		receivedDroneStrikes := initDroneStrikes(mockedCar)
		Expect(receivedDroneStrikes).To(Equal(expectedDronesStrikes))
	})
})

var _ = Describe("landAllFlyingDrones", func() {
	var mockedCtrl *gomock.Controller
	var mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDroneStrikes []droneStrikes

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDroneStrikes = []droneStrikes{
			{drone: mockedDrone1},
			{drone: mockedDrone2},
		}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should land all flying drones", func() {
		mockedDrone1.EXPECT().IsFlying().Return(true)
		mockedDrone1.EXPECT().Land(mockedCarStop)
		mockedDrone2.EXPECT().IsFlying().Return(false)
		landAllFlyingDrones(mockedDroneStrikes, mockedCarStop)
	})
})

var _ = Describe("anyDroneWasStriked", func() {
	Context("when any drone was striked", func() {
		var mockedCtrl *gomock.Controller
		var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDroneStrikes []droneStrikes

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDroneStrikes = []droneStrikes{
				{drone: mockedDrone1, strikes: 0},
				{drone: mockedDrone2, strikes: maxStrikes},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if any drone was striked", func() {
			Expect(anyDroneWasStriked(mockedDroneStrikes)).To(BeTrue())
		})
	})

	Context("when no drone was striked", func() {
		var mockedCtrl *gomock.Controller
		var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDroneStrikes []droneStrikes

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDroneStrikes = []droneStrikes{
				{drone: mockedDrone1, strikes: 0},
				{drone: mockedDrone2, strikes: 0},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return false if no drone was striked", func() {
			Expect(anyDroneWasStriked(mockedDroneStrikes)).To(BeFalse())
		})
	})
})

var _ = Describe("anyDroneNeedToLand", func() {
	Context("when any drone need to land", func() {
		var mockedCtrl *gomock.Controller
		var mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDroneStrikes []droneStrikes
		var mockedCarStopPoint = gps.Point{}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
			mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDroneStrikes = []droneStrikes{
				{drone: mockedDrone1},
				{drone: mockedDrone2},
				{drone: mockedDrone3},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if any drone need to land", func() {
			mockedCarStop.EXPECT().Point().Return(mockedCarStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(mockedCarStopPoint).Return(true)
			mockedDrone3.EXPECT().IsFlying().Return(true)
			mockedDrone3.EXPECT().CanReach(mockedCarStopPoint).Return(false)
			Expect(anyDroneNeedToLand(mockedDroneStrikes, mockedCarStop)).To(BeTrue())
		})
	})

	Context("when no drone need to land", func() {
		var mockedCtrl *gomock.Controller
		var mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		var mockedDroneStrikes []droneStrikes
		var mockedCarStopPoint = gps.Point{}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
			mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
			mockedDroneStrikes = []droneStrikes{
				{drone: mockedDrone1},
				{drone: mockedDrone2},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return false if no drone need to land", func() {
			mockedCarStop.EXPECT().Point().Return(mockedCarStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(mockedCarStopPoint).Return(true)
			Expect(anyDroneNeedToLand(mockedDroneStrikes, mockedCarStop)).To(BeFalse())
		})
	})
})

var _ = Describe("updateDroneStrikes", func() {
	var mockedCtrl *gomock.Controller
	var mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDroneStrikes []droneStrikes
	var mockedStopPoint = gps.Point{}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDroneStrikes = []droneStrikes{
			{drone: mockedDrone1, strikes: 0},
			{drone: mockedDrone2, strikes: 0},
			{drone: mockedDrone3, strikes: 0},
		}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should update drone strikes", func() {
		mockedCarStop.EXPECT().Point().Return(mockedStopPoint)
		mockedDrone1.EXPECT().IsFlying().Return(true)
		mockedDrone1.EXPECT().Support(mockedStopPoint).Return(true)
		mockedDrone2.EXPECT().IsFlying().Return(true)
		mockedDrone2.EXPECT().Support(mockedStopPoint).Return(false)
		mockedDrone3.EXPECT().IsFlying().Return(false)
		updateDroneStrikes(mockedDroneStrikes, mockedCarStop)
		Expect(mockedDroneStrikes[0].strikes).To(Equal(0))
		Expect(mockedDroneStrikes[1].strikes).To(Equal(1))
		Expect(mockedDroneStrikes[2].strikes).To(Equal(0))
	})
})

var _ = Describe("flyingDroneThatCanSupport", func() {
	var mockedCtrl *gomock.Controller
	var mockedActualCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
	var mockedNextCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone4 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDroneStrikes []droneStrikes
	var mockedActualStopPoint = gps.Point{}
	var mockedNextStopPoint = gps.Point{}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedActualCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		mockedNextCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone4 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDroneStrikes = []droneStrikes{
			{drone: mockedDrone1},
			{drone: mockedDrone2},
			{drone: mockedDrone3},
			{drone: mockedDrone4},
		}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should return first flying drone that can support", func() {
		mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
		mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
		mockedDrone1.EXPECT().IsFlying().Return(false)
		mockedDrone2.EXPECT().IsFlying().Return(true)
		mockedDrone2.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
		mockedDrone3.EXPECT().IsFlying().Return(true)
		mockedDrone3.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(true)
		Expect(flyingDroneThatCanSupport(mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)).To(Equal(mockedDrone3))
	})
})

var _ = Describe("dockedDroneThatCanSupport", func() {
	var mockedCtrl *gomock.Controller
	var mockedActualCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
	var mockedNextCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDrone4 = mockVehicles.NewMockIDrone(mockedCtrl)
	var mockedDroneStrikes []droneStrikes
	var mockedActualStopPoint = gps.Point{}
	var mockedNextStopPoint = gps.Point{}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedActualCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		mockedNextCarStop = mockRoutes.NewMockIMainStop(mockedCtrl)
		mockedDrone1 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone3 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDrone4 = mockVehicles.NewMockIDrone(mockedCtrl)
		mockedDroneStrikes = []droneStrikes{
			{drone: mockedDrone1},
			{drone: mockedDrone2},
			{drone: mockedDrone3},
			{drone: mockedDrone4},
		}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should return first docked drone that can support", func() {
		mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
		mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
		mockedDrone1.EXPECT().IsFlying().Return(true)
		mockedDrone2.EXPECT().IsFlying().Return(false)
		mockedDrone2.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
		mockedDrone3.EXPECT().IsFlying().Return(false)
		mockedDrone3.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(true)
		Expect(dockedDroneThatCanSupport(mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)).To(Equal(mockedDrone2))
	})
})

var _ = Describe("DroneStrikesInsertion", func() {
	Context("when both drones are docked and can support actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedRoute *mockRoutes.MockIMainRoute
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			mockedRoute = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 1 to first client", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone1.EXPECT().Support(clientPoint, depositPoint).Return(true)
			// Move drone 1 to the first client and remove the stop from the route
			mockedDrone1.EXPECT().Move(routes.NewSubStop(clientPoint))
			mockedRoute.EXPECT().RemoveMainStop(0)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedDepositStop)
			mockedDrone2.EXPECT().IsFlying().Return(false)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when drone 1 is flying and only drone 1 can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedRoute *mockRoutes.MockIMainRoute
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			mockedRoute = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 1 to actual client", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(depositPoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(clientPoint, depositPoint).Return(false)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint, depositPoint).Return(true)
			// Move drone 2 to the first client and remove the stop from the route
			mockedDrone1.EXPECT().Move(routes.NewSubStop(clientPoint))
			mockedRoute.EXPECT().RemoveMainStop(0)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedDepositStop)
			mockedDrone2.EXPECT().IsFlying().Return(false)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when drone 1 is flying and both drones can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedRoute *mockRoutes.MockIMainRoute
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			mockedRoute = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 2 to actual client", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(depositPoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(clientPoint, depositPoint).Return(true)
			// Move drone 2 to the first client and remove the stop from the route
			mockedDrone2.EXPECT().Move(routes.NewSubStop(clientPoint))
			mockedRoute.EXPECT().RemoveMainStop(0)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedDepositStop)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Land(mockedDepositStop)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when both drones are flying and drone 1 can not reach next stop", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			_ = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(depositPoint).Return(false)
			// Land all drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedClientStop)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Land(mockedDepositStop)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when both drones are flying and none can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			_ = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should continue without move drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(depositPoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(depositPoint).Return(true)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(clientPoint).Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(clientPoint).Return(false)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedDepositStop)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Land(mockedDepositStop)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when both drones are flying and only drone 2 can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedRoute *mockRoutes.MockIMainRoute
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			mockedRoute = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 2 to actual client", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(depositPoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(depositPoint).Return(true)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(clientPoint).Return(true)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(clientPoint).Return(true)
			// Move drone 2 to the first client and remove the stop from the route
			mockedDrone2.EXPECT().Move(routes.NewSubStop(clientPoint))
			mockedRoute.EXPECT().RemoveMainStop(0)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedDepositStop)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Land(mockedDepositStop)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when both drones are flying and drone 1 need to land", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			_ = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land both drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(depositPoint).Return(false)
			// Land all drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedClientStop)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Land(mockedDepositStop)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when one drone is flying and other is docked but both can not support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			_ = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedClientStop, mockedDepositStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should continue without move drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(depositPoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(clientPoint).Return(false)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(clientPoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedClientStop)
			mockedDrone2.EXPECT().IsFlying().Return(false)

			DroneStrikesInsertion(mockedCar)
		})
	})

	Context("when both drones are flying and actual stop is deposit", func() {
		var mockedCtrl *gomock.Controller
		var mockedCar *mockVehicles.MockICar
		var mockedDrone1 *mockVehicles.MockIDrone
		var mockedDrone2 *mockVehicles.MockIDrone
		var mockedClientStop *mockRoutes.MockIMainStop
		var mockedDepositStop *mockRoutes.MockIMainStop
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var depositPoint = gps.Point{Name: "Deposit"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedCar, mockedDrone1, mockedDrone2 = mockCarWithDrones(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedDepositStop = mockDepositStop(mockedCtrl, depositPoint)
			_ = mockCarRouteWithStops(mockedCtrl, mockedCar, mockedDepositStop, mockedClientStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones", func() {
			// Land all drones
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Land(mockedDepositStop)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Land(mockedDepositStop)
			// Finish landing all flying drones
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)

			DroneStrikesInsertion(mockedCar)
		})
	})
})

func mockClientStop(ctrl *gomock.Controller, point gps.Point) *mockRoutes.MockIMainStop {
	mockedStop := mockRoutes.NewMockIMainStop(ctrl)
	mockedStop.EXPECT().Point().Return(point).AnyTimes()
	mockedStop.EXPECT().IsDeposit().Return(false).AnyTimes()
	mockedStop.EXPECT().IsClient().Return(true).AnyTimes()
	return mockedStop
}

func mockDepositStop(ctrl *gomock.Controller, point gps.Point) *mockRoutes.MockIMainStop {
	mockedStop := mockRoutes.NewMockIMainStop(ctrl)
	mockedStop.EXPECT().Point().Return(point).AnyTimes()
	mockedStop.EXPECT().IsDeposit().Return(true).AnyTimes()
	mockedStop.EXPECT().IsClient().Return(false).AnyTimes()
	return mockedStop
}

func mockCarWithDrones(ctrl *gomock.Controller) (*mockVehicles.MockICar, *mockVehicles.MockIDrone, *mockVehicles.MockIDrone) {
	mockedCar := mockVehicles.NewMockICar(ctrl)
	mockedDrone1 := mockVehicles.NewMockIDrone(ctrl)
	mockedDrone2 := mockVehicles.NewMockIDrone(ctrl)
	mockedCar.EXPECT().Drones().Return([]vehicles.IDrone{mockedDrone1, mockedDrone2})
	return mockedCar, mockedDrone1, mockedDrone2
}

func mockCarRouteWithStops(ctrl *gomock.Controller, mockedCar *mockVehicles.MockICar, stops ...*mockRoutes.MockIMainStop) *mockRoutes.MockIMainRoute {
	mockedRoute := mockRoutes.NewMockIMainRoute(ctrl)
	stopsList := []routes.IMainStop{}
	for _, stop := range stops {
		stopsList = append(stopsList, stop)
	}
	mockedCar.EXPECT().Route().Return(mockedRoute).AnyTimes()
	mockedIterator := slc.NewIterator(stopsList)
	mockedRoute.EXPECT().Iterator().Return(mockedIterator)
	return mockedRoute
}
