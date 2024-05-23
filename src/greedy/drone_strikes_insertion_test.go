package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mock"
	"github.com/victorguarana/vehicle-routing/src/route"
	mockroute "github.com/victorguarana/vehicle-routing/src/route/mock"
	"github.com/victorguarana/vehicle-routing/src/slc"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("initDroneStrikes", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor *mockitinerary.MockConstructor
	var mockedDrone1 = itinerary.DroneNumber(1)
	var mockedDrone2 = itinerary.DroneNumber(2)

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should initialize drone strikes", func() {
		expectedDronesStrikes := []droneStrikes{
			{droneNumber: mockedDrone1},
			{droneNumber: mockedDrone2},
		}
		mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
		receivedDroneStrikes := initDroneStrikes(mockedConstructor)
		Expect(receivedDroneStrikes).To(Equal(expectedDronesStrikes))
	})
})

var _ = Describe("anyDroneWasStriked", func() {
	Context("when any drone was striked", func() {
		var mockedCtrl *gomock.Controller
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedDroneStrikes []droneStrikes

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedDroneStrikes = []droneStrikes{
				{droneNumber: mockedDrone1, strikes: 0},
				{droneNumber: mockedDrone2, strikes: maxStrikes},
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
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedDroneStrikes = []droneStrikes{
			{droneNumber: mockedDrone1, strikes: 0},
			{droneNumber: mockedDrone2, strikes: 0},
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
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedDrone3 = itinerary.DroneNumber(3)
		var mockedDroneStrikes = []droneStrikes{
			{droneNumber: mockedDrone1},
			{droneNumber: mockedDrone2},
			{droneNumber: mockedDrone3},
		}
		var mockedConstructorStopPoint = gps.Point{}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructorStop = mockroute.NewMockIMainStop(mockedCtrl)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return true if any drone need to land", func() {
			mockedConstructorStop.EXPECT().Point().Return(mockedConstructorStopPoint)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone2, mockedConstructorStopPoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone3).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone3, mockedConstructorStopPoint).Return(false)
			Expect(anyDroneNeedToLand(mockedConstructor, mockedDroneStrikes, mockedConstructorStop)).To(BeTrue())
		})
	})

	Context("when no drone need to land", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		var mockedConstructorStop = mockroute.NewMockIMainStop(mockedCtrl)
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedDroneStrikes = []droneStrikes{
			{droneNumber: mockedDrone1},
			{droneNumber: mockedDrone2},
		}
		var mockedConstructorStopPoint = gps.Point{}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructorStop = mockroute.NewMockIMainStop(mockedCtrl)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should return false if no drone need to land", func() {
			mockedConstructorStop.EXPECT().Point().Return(mockedConstructorStopPoint)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone2, mockedConstructorStopPoint).Return(true)
			Expect(anyDroneNeedToLand(mockedConstructor, mockedDroneStrikes, mockedConstructorStop)).To(BeFalse())
		})
	})
})

var _ = Describe("updateDroneStrikes", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
	var mockedDeliveryStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedLandingStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 = itinerary.DroneNumber(1)
	var mockedDrone2 = itinerary.DroneNumber(2)
	var mockedDrone3 = itinerary.DroneNumber(3)
	var mockedDroneStrikes = []droneStrikes{
		{droneNumber: mockedDrone1, strikes: 0},
		{droneNumber: mockedDrone2, strikes: 0},
		{droneNumber: mockedDrone3, strikes: 0},
	}
	var deliveryPoint = gps.Point{Name: "Delivery Point"}
	var landingPoint = gps.Point{Name: "Landing Point"}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedDeliveryStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedLandingStop = mockroute.NewMockIMainStop(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	It("should update drone strikes", func() {
		mockedDeliveryStop.EXPECT().Point().Return(deliveryPoint)
		mockedLandingStop.EXPECT().Point().Return(landingPoint)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
		mockedConstructor.EXPECT().DroneSupport(mockedDrone1, deliveryPoint, landingPoint).Return(true)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
		mockedConstructor.EXPECT().DroneSupport(mockedDrone2, deliveryPoint, landingPoint).Return(false)
		mockedConstructor.EXPECT().DroneIsFlying(mockedDrone3).Return(false)
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
	var mockedDrone1 = itinerary.DroneNumber(1)
	var mockedDrone2 = itinerary.DroneNumber(2)
	var mockedDrone3 = itinerary.DroneNumber(3)
	var mockedDrone4 = itinerary.DroneNumber(4)
	var mockedDroneStrikes = []droneStrikes{
		{droneNumber: mockedDrone1},
		{droneNumber: mockedDrone2},
		{droneNumber: mockedDrone3},
		{droneNumber: mockedDrone4},
	}
	var mockedActualStopPoint = gps.Point{}
	var mockedNextStopPoint = gps.Point{}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedActualCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedNextCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	Context("when there is no flying drone that can support", func() {
		It("should return false", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone3).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone3, mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone4).Return(false)
			receivedDroneNumber, receivedExists := flyingDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(Equal(itinerary.DroneNumber(0)))
			Expect(receivedExists).To(BeFalse())
		})
	})

	Context("when there is a flying drone that can support", func() {
		It("should return first flying drone that can support", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone3).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone3, mockedActualStopPoint, mockedNextStopPoint).Return(true)
			receivedDroneNumber, receivedExists := flyingDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(Equal(itinerary.DroneNumber(3)))
			Expect(receivedExists).To(BeTrue())
		})
	})
})

var _ = Describe("dockedDroneThatCanSupport", func() {
	var mockedCtrl *gomock.Controller
	var mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
	var mockedActualCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedNextCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	var mockedDrone1 = itinerary.DroneNumber(1)
	var mockedDrone2 = itinerary.DroneNumber(2)
	var mockedDrone3 = itinerary.DroneNumber(3)
	var mockedDrone4 = itinerary.DroneNumber(4)
	var mockedDroneStrikes = []droneStrikes{
		{droneNumber: mockedDrone1},
		{droneNumber: mockedDrone2},
		{droneNumber: mockedDrone3},
		{droneNumber: mockedDrone4},
	}
	var mockedActualStopPoint = gps.Point{}
	var mockedNextStopPoint = gps.Point{}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
		mockedActualCarStop = mockroute.NewMockIMainStop(mockedCtrl)
		mockedNextCarStop = mockroute.NewMockIMainStop(mockedCtrl)
	})

	AfterEach(func() {
		mockedCtrl.Finish()
	})

	Context("when there is no docked drone that can support", func() {
		It("should return false", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone3).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone3, mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone4).Return(true)
			receivedDroneNumber, receivedExists := dockedDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(Equal(itinerary.DroneNumber(0)))
			Expect(receivedExists).To(BeFalse())
		})
	})

	Context("when there is a docked drone that can support", func() {
		It("should return first docked drone that can support", func() {
			mockedActualCarStop.EXPECT().Point().Return(mockedActualStopPoint)
			mockedNextCarStop.EXPECT().Point().Return(mockedNextStopPoint)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, mockedActualStopPoint, mockedNextStopPoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone3).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone3, mockedActualStopPoint, mockedNextStopPoint).Return(true)
			receivedDroneNumber, receivedExists := dockedDroneThatCanSupport(mockedConstructor, mockedDroneStrikes, mockedActualCarStop, mockedNextCarStop)
			Expect(receivedDroneNumber).To(Equal(itinerary.DroneNumber(3)))
			Expect(receivedExists).To(BeTrue())
		})
	})
})

var _ = Describe("DroneStrikesInsertion", func() {
	Context("when both drones are docked and can support actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 1 to first client", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Updating drone strikes
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Search for a docked drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(true)
			// Start drone 1 flight and move to the first client and remove the stop from the route
			mockedConstructor.EXPECT().StartDroneFlight(mockedDrone1, mockedInitialWarehouseStop)
			mockedConstructor.EXPECT().MoveDrone(mockedDrone1, clientPoint)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(1)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when drone 1 is flying and only drone 1 can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 1 to actual client", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone1, warehousePoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Updating drone strikes
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Search for a docked drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, clientPoint, warehousePoint).Return(false)
			// Search for a flying drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(true)
			// Move drone 2 to the first client and remove the stop from the route
			mockedConstructor.EXPECT().MoveDrone(mockedDrone1, clientPoint)
			mockedModifier.EXPECT().RemoveMainStopFromRoute(1)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when drone 1 is flying and both drones can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 2 to actual client", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone1, warehousePoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Updating drone strikes
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Search for a docked drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, clientPoint, warehousePoint).Return(true)
			// Start drone 2 flight and move to the first client and remove the stop from the route
			mockedConstructor.EXPECT().StartDroneFlight(mockedDrone2, mockedInitialWarehouseStop)
			mockedConstructor.EXPECT().MoveDrone(mockedDrone2, clientPoint)
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
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone1, warehousePoint).Return(false)
			// Land all drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and none can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should continue without move drones", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone1, warehousePoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone2, warehousePoint).Return(true)
			// Updating drone strikes
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, clientPoint, warehousePoint).Return(false)
			// Search for a docked drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			// Search for a flying drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, clientPoint, warehousePoint).Return(false)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and only drone 2 can support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should move drone 2 to actual client", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone1, warehousePoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone2, warehousePoint).Return(true)
			// Updating drone strikes
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, clientPoint, warehousePoint).Return(true)
			// Search for a docked drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			// Search for a flying drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, clientPoint, warehousePoint).Return(true)
			// Move drone 2 to the first client and remove the stop from the route
			mockedConstructor.EXPECT().MoveDrone(mockedDrone2, clientPoint)
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
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land both drones", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone1, warehousePoint).Return(false)
			// Land all drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalWarehouseStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when one drone is flying and other is docked but both can not support the actual client", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedClientStop *mockroute.MockIMainStop
		var mockedInitialWarehouseStop *mockroute.MockIMainStop
		var mockedFinalWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedInitialWarehouseStop = mockWarehouseStop(mockedCtrl, initialPoint)
			mockedFinalWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialWarehouseStop, mockedClientStop, mockedFinalWarehouseStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should continue without move drones", func() {
			// Checking if any drone need to land
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneCanReach(mockedDrone1, warehousePoint).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Updating drone strikes
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Search for a docked drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone2, clientPoint, warehousePoint).Return(false)
			// Search for a flying drone that can support
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone1).Return(true)
			mockedConstructor.EXPECT().DroneSupport(mockedDrone1, clientPoint, warehousePoint).Return(false)
			mockedConstructor.EXPECT().DroneIsFlying(mockedDrone2).Return(false)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedClientStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})

	Context("when both drones are flying and actual stop is warehouse", func() {
		var mockedCtrl *gomock.Controller
		var mockedConstructor *mockitinerary.MockConstructor
		var mockedModifier *mockitinerary.MockModifier
		var mockedDrone1 = itinerary.DroneNumber(1)
		var mockedDrone2 = itinerary.DroneNumber(2)
		var mockedInitialClientStop *mockroute.MockIMainStop
		var mockedFinalClientStop *mockroute.MockIMainStop
		var mockedWarehouseStop *mockroute.MockIMainStop
		var initialPoint = gps.Point{Name: "Initial Point"}
		var clientPoint = gps.Point{Name: "Client", Latitude: 1}
		var warehousePoint = gps.Point{Name: "Warehouse"}

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedConstructor = mockitinerary.NewMockConstructor(mockedCtrl)
			mockedConstructor.EXPECT().DroneNumbers().Return([]itinerary.DroneNumber{mockedDrone1, mockedDrone2})
			mockedModifier = mockitinerary.NewMockModifier(mockedCtrl)
			mockedInitialClientStop = mockClientStop(mockedCtrl, initialPoint)
			mockedFinalClientStop = mockClientStop(mockedCtrl, clientPoint)
			mockedWarehouseStop = mockWarehouseStop(mockedCtrl, warehousePoint)
			fillItineraryStops(mockedConstructor, mockedInitialClientStop, mockedWarehouseStop, mockedFinalClientStop)
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should land all drones", func() {
			// Land all drones
			mockedConstructor.EXPECT().LandAllDrones(mockedWarehouseStop)
			// Finish landing all flying drones
			mockedConstructor.EXPECT().LandAllDrones(mockedFinalClientStop)

			DroneStrikesInsertion(mockedConstructor, mockedModifier)
		})
	})
})

func mockClientStop(ctrl *gomock.Controller, point gps.Point) *mockroute.MockIMainStop {
	mockedStop := mockroute.NewMockIMainStop(ctrl)
	mockedStop.EXPECT().Point().Return(point).AnyTimes()
	mockedStop.EXPECT().IsWarehouse().Return(false).AnyTimes()
	mockedStop.EXPECT().IsClient().Return(true).AnyTimes()
	return mockedStop
}

func mockWarehouseStop(ctrl *gomock.Controller, point gps.Point) *mockroute.MockIMainStop {
	mockedStop := mockroute.NewMockIMainStop(ctrl)
	mockedStop.EXPECT().Point().Return(point).AnyTimes()
	mockedStop.EXPECT().IsWarehouse().Return(true).AnyTimes()
	mockedStop.EXPECT().IsClient().Return(false).AnyTimes()
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
