package greedy

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	mockitinerary "github.com/victorguarana/vehicle-routing/internal/itinerary/mock"
	"github.com/victorguarana/vehicle-routing/internal/route"
	mockroute "github.com/victorguarana/vehicle-routing/internal/route/mock"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("initDroneStrikes", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor *mockitinerary.MockConstructor
	var mockedCar *mockvehicle.MockICar
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedCar = mockvehicle.NewMockICar(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should initialize drone strikes", func() {
		expectedDronesStrikes := []droneStrikes{
			{drone: mockedDrone1},
			{drone: mockedDrone2},
		}
		mockedConstructor.EXPECT().Car().Return(mockedCar)
		mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
		receivedDroneStrikes := initDroneStrikes(mockedConstructor)
		Expect(receivedDroneStrikes).To(Equal(expectedDronesStrikes))
	})
})

var _ = Describe("anyDroneWasStriked", func() {
	Context("when any drone was striked", func() {
		var mockedCtrl *gomock.Controller
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedDroneStrikes []droneStrikes

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
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
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedDroneStrikes = []droneStrikes{
			{drone: mockedDrone1, strikes: 0},
			{drone: mockedDrone2, strikes: 0},
		}

		It("should return false if no drone was striked", func() {
			Expect(anyDroneWasStriked(mockedDroneStrikes)).To(BeFalse())
		})
	})
})

var _ = Describe("anyDroneNeedToLand", func() {
	Context("when any drone need to land", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		var mockedConstructorStop = mockroute.NewMockIMainStop(mockedCtrl)
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedDrone3 *mockvehicle.MockIDrone
		var mockedDroneStrikes []droneStrikes
		var mockedConstructorStopPoint = gps.Point{}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructorStop = mockroute.NewMockIMainStop(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone3 = mockvehicle.NewMockIDrone(mockedCtrl)
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
			mockedConstructorStop.EXPECT().Point().Return(mockedConstructorStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(mockedConstructorStopPoint).Return(true)
			mockedDrone3.EXPECT().IsFlying().Return(true)
			mockedDrone3.EXPECT().CanReach(mockedConstructorStopPoint).Return(false)
			Expect(anyDroneNeedToLand(mockedConstructor, mockedDroneStrikes, mockedConstructorStop)).To(BeTrue())
		})
	})

	Context("when no drone need to land", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		var mockedConstructorStop = mockroute.NewMockIMainStop(mockedCtrl)
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedDroneStrikes []droneStrikes
		var mockedConstructorStopPoint = gps.Point{}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructorStop = mockroute.NewMockIMainStop(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDroneStrikes = []droneStrikes{
				{drone: mockedDrone1},
				{drone: mockedDrone2},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return false if no drone need to land", func() {
			mockedConstructorStop.EXPECT().Point().Return(mockedConstructorStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(mockedConstructorStopPoint).Return(true)
			Expect(anyDroneNeedToLand(mockedConstructor, mockedDroneStrikes, mockedConstructorStop)).To(BeFalse())
		})
	})
})

var _ = Describe("updateDroneStrikes", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
	var mockedDeliveryStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedLandingStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedDrone3 *mockvehicle.MockIDrone
	var mockedDroneStrikes []droneStrikes
	var deliveryPoint = gps.Point{Name: "Delivery Point"}
	var landingPoint = gps.Point{Name: "Landing Point"}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedDeliveryStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedLandingStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone3 = mockvehicle.NewMockIDrone(mockedCtrl)
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
		mockedDeliveryStop.EXPECT().Point().Return(deliveryPoint)
		mockedLandingStop.EXPECT().Point().Return(landingPoint)
		mockedDrone1.EXPECT().IsFlying().Return(true)
		mockedDrone1.EXPECT().Support(deliveryPoint, landingPoint).Return(true)
		mockedDrone2.EXPECT().IsFlying().Return(true)
		mockedDrone2.EXPECT().Support(deliveryPoint, landingPoint).Return(false)
		mockedDrone3.EXPECT().IsFlying().Return(false)
		updateDroneStrikes(mockedConstructor, mockedDroneStrikes, mockedDeliveryStop, mockedLandingStop)
		Expect(mockedDroneStrikes[0].strikes).To(Equal(0))
		Expect(mockedDroneStrikes[1].strikes).To(Equal(1))
		Expect(mockedDroneStrikes[2].strikes).To(Equal(0))
	})
})

var _ = Describe("flyingDroneThatCanSupport", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
	var mockedActualCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedNextCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedDrone3 *mockvehicle.MockIDrone
	var mockedDrone4 *mockvehicle.MockIDrone
	var mockedDroneStrikes []droneStrikes
	var mockedActualStopPoint = gps.Point{}
	var mockedNextStopPoint = gps.Point{}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedActualCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedNextCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone3 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone4 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDroneStrikes = []droneStrikes{
			{drone: mockedDrone1, strikes: 0},
			{drone: mockedDrone2, strikes: 0},
			{drone: mockedDrone3, strikes: 0},
			{drone: mockedDrone4, strikes: 0},
		}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	Context("when there is no flying drone that can support", func() {
		It("should return false", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedDrone3.EXPECT().IsFlying().Return(true)
			mockedDrone3.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedDrone4.EXPECT().IsFlying().Return(false)
			receivedDroneNumber, receivedExists := flyingDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(BeNil())
			Expect(receivedExists).To(BeFalse())
		})
	})

	Context("when there is a flying drone that can support", func() {
		It("should return first flying drone that can support", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedDrone3.EXPECT().IsFlying().Return(true)
			mockedDrone3.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(true)
			receivedDroneNumber, receivedExists := flyingDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(Equal(mockedDrone3))
			Expect(receivedExists).To(BeTrue())
		})
	})
})

var _ = Describe("dockedDroneThatCanSupport", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
	var mockedActualCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedNextCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var mockedDrone3 *mockvehicle.MockIDrone
	var mockedDrone4 *mockvehicle.MockIDrone
	var mockedDroneStrikes []droneStrikes
	var mockedActualStopPoint = gps.Point{}
	var mockedNextStopPoint = gps.Point{}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedActualCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedNextCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone3 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone4 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDroneStrikes = []droneStrikes{
			{drone: mockedDrone1, strikes: 0},
			{drone: mockedDrone2, strikes: 0},
			{drone: mockedDrone3, strikes: 0},
			{drone: mockedDrone4, strikes: 0},
		}
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	Context("when there is no docked drone that can support", func() {
		It("should return false", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedDrone3.EXPECT().IsFlying().Return(false)
			mockedDrone3.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedDrone4.EXPECT().IsFlying().Return(true)
			receivedDroneNumber, receivedExists := dockedDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(BeNil())
			Expect(receivedExists).To(BeFalse())
		})
	})

	Context("when there is a docked drone that can support", func() {
		It("should return first docked drone that can support", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedDrone3.EXPECT().IsFlying().Return(false)
			mockedDrone3.EXPECT().Support(mockedActualStopPoint, mockedNextStopPoint).Return(true)
			receivedDroneNumber, receivedExists := dockedDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(Equal(mockedDrone3))
			Expect(receivedExists).To(BeTrue())
		})
	})
})

var _ = Describe("DroneStrikesInsertion", func() {
	Context("when both drones are docked and can support actual customer", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 1 to first customer", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(false)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(true)
			// Start drone 1 flight and move to the first customer and remove the stop from the route
			mockedConstructor.EXPECT().StartDroneFlight(mockedDrone1, mockedInitialWarehouseStop)
			mockedConstructor.EXPECT().MoveDrone(mockedDrone1, customerPoint)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(1)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when drone 1 is flying and only drone 1 can support the actual customer", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 1 to actual customer", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(warehousePoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(true)
			// Move drone 2 to the first customer and remove the stop from the route
			mockedConstructor.EXPECT().MoveDrone(mockedDrone1, customerPoint)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(1)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when drone 1 is flying and both drones can support the actual customer", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 2 to actual customer", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(warehousePoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(customerPoint, warehousePoint).Return(true)
			// Start drone 2 flight and move to the first customer and remove the stop from the route
			mockedConstructor.EXPECT().StartDroneFlight(mockedDrone2, mockedInitialWarehouseStop)
			mockedConstructor.EXPECT().MoveDrone(mockedDrone2, customerPoint)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(1)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and drone 1 can not reach next stop", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(warehousePoint).Return(false)
			// Land all drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and none can support the actual customer", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should continue without move drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(warehousePoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(warehousePoint).Return(true)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and only drone 2 can support the actual customer", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 2 to actual customer", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(warehousePoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().CanReach(warehousePoint).Return(true)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(customerPoint, warehousePoint).Return(true)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().Support(customerPoint, warehousePoint).Return(true)
			// Move drone 2 to the first customer and remove the stop from the route
			mockedConstructor.EXPECT().MoveDrone(mockedDrone2, customerPoint)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(1)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and drone 1 need to land", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land both drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(warehousePoint).Return(false)
			// Land all drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when one drone is flying and other is docked but both can not support the actual customer", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedCustomerStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedCustomerStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should continue without move drones", func() {
			// Checking if any drone need to land
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().CanReach(warehousePoint).Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Updating drone strikes
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Search for a docked drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			mockedDrone2.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			// Search for a flying drone that can support
			mockedDrone1.EXPECT().IsFlying().Return(true)
			mockedDrone1.EXPECT().Support(customerPoint, warehousePoint).Return(false)
			mockedDrone2.EXPECT().IsFlying().Return(false)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedCustomerStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and actual stop is warehouse", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedCar *mockvehicle.MockICar
		var mockedDrone1 *mockvehicle.MockIDrone
		var mockedDrone2 *mockvehicle.MockIDrone
		var mockedInitialCustomerStop *mockroute.MockIMainStop
		var mockedFinalCustomerStop *mockroute.MockIMainStop
		var mockedWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var customerPoint = gps.Point{Name: "Customer", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedCar = mockvehicle.NewMockICar(mockedCtrl)
			mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
			mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedInitialCustomerStop = mockCustomerStop(mockedCtrl, initialPoint)
			mockedFinalCustomerStop = mockCustomerStop(mockedCtrl, customerPoint)
			mockedWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedConstructor.EXPECT().Car().Return(mockedCar).AnyTimes()
			fillItineraryStops(mockedConstructor, mockedInitialCustomerStop, mockedWarehouseStop, mockedFinalCustomerStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones", func() {
			// Land all drones
			mockedConstructor.EXPECT().LandAllDrones(mockedWarehouseStop)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalCustomerStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})
})

func mockCustomerStop(ctrl *gomock.Controller, point gps.Point) *mockroute.MockIMainStop {
	mockedStop := mockroute.NewMockIMainStop(ctrl)
	mockedStop.EXPECT().Point().Return(point).AnyTimes()
	mockedStop.EXPECT().IsWarehouse().Return(false).AnyTimes()
	mockedStop.EXPECT().IsCustomer().Return(true).AnyTimes()
	return mockedStop
}

func mockWarehouseStop(ctrl *gomock.Controller, point gps.Point) *mockroute.MockIMainStop {
	mockedStop := mockroute.NewMockIMainStop(ctrl)
	mockedStop.EXPECT().Point().Return(point).AnyTimes()
	mockedStop.EXPECT().IsWarehouse().Return(true).AnyTimes()
	mockedStop.EXPECT().IsCustomer().Return(false).AnyTimes()
	return mockedStop
}

func fillItineraryStops(mockedConstructor *mockitinerary.MockConstructor, stops ...*mockroute.MockIMainStop) {
	stopsList := []route.IMainStop{}
	for _, stop := range stops {
		stopsList = append(stopsList, stop)
	}
	routeIterator := slc.NewIterator(stopsList)
	routeIterator.GoToNext()
	mockedConstructor.EXPECT().RouteIterator().Return(routeIterator).AnyTimes()
}
