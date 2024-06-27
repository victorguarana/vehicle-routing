package itinerary

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
	mockvehicle "github.com/victorguarana/vehicle-routing/internal/vehicle/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("New", func() {
	var mockedCtrl *gomock.Controller
	var mockedCar *mockvehicle.MockICar
	var mockedDrone1 *mockvehicle.MockIDrone
	var mockedDrone2 *mockvehicle.MockIDrone
	var initialPoint = gps.Point{Latitude: 1, Longitude: 2, PackageSize: 3, Name: "initialPoint"}

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicle.NewMockICar(mockedCtrl)
		mockedDrone1 = mockvehicle.NewMockIDrone(mockedCtrl)
		mockedDrone2 = mockvehicle.NewMockIDrone(mockedCtrl)
	})

	It("should return an itinerary", func() {
		mockedCar.EXPECT().Drones().Return([]vehicle.IDrone{mockedDrone1, mockedDrone2})
		mockedCar.EXPECT().ActualPoint().Return(initialPoint)
		expectedItinerary := &itinerary{
			activeFlights:             map[DroneNumber]route.ISubRoute{},
			droneNumbersMap:           map[DroneNumber]vehicle.IDrone{1: mockedDrone1, 2: mockedDrone2},
			car:                       mockedCar,
			completedSubItineraryList: []SubItinerary{},
			route:                     route.NewMainRoute(route.NewMainStop(initialPoint)),
		}
		receivedItinerary := New(mockedCar)
		Expect(receivedItinerary).To(Equal(expectedItinerary))
	})
})
