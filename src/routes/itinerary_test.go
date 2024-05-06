package routes

import (
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewItinerary", func() {
	var (
		mockCtrl  *gomock.Controller
		mockedCar *mockvehicles.MockICar
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
	})

	Context("when all params are valid", func() {
		var intitialPosition = gps.Point{}

		It("returns new itinerary", func() {
			ms := NewMainStop(intitialPosition).(*mainStop)
			expectedItinerary := Itinerary{
				Car:   mockedCar,
				Route: &mainRoute{[]*mainStop{ms}},
			}
			mockedCar.EXPECT().ActualPosition().Return(intitialPosition)
			receivedItinerary := NewItinerary(mockedCar)
			Expect(receivedItinerary).To(Equal(expectedItinerary))
		})
	})
})
