package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	mockRoutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"
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
