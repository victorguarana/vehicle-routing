package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	mockitinerary "github.com/victorguarana/vehicle-routing/src/itinerary/mock"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("finishRoutesOnClosestWarehouses", func() {
	var mockCtrl *gomock.Controller
	var mockedConstructor *mockitinerary.MockConstructor
	var constructorList []itinerary.Constructor
	var closestWarehouse = gps.Point{Latitude: 1}
	var actualCarPoint = gps.Point{Latitude: 0}
	var gpsMap = gps.Map{Warehouses: []gps.Point{closestWarehouse}}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedConstructor = mockitinerary.NewMockConstructor(mockCtrl)
		constructorList = []itinerary.Constructor{mockedConstructor}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest warehouse and append it to the route", func() {
			mockedConstructor.EXPECT().ActualCarPoint().Return(actualCarPoint)
			mockedConstructor.EXPECT().MoveCar(closestWarehouse)
			finishOnClosestWarehouses(constructorList, gpsMap)
		})
	})
})
