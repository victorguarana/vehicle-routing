package itinerary

import (
	"go.uber.org/mock/gomock"

	mockroute "github.com/victorguarana/vehicle-routing/src/route/mock"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("modifier{}", func() {
	Describe("RemoveMainStopFromRoute", func() {
		var sut modifier
		var mockedCtrl *gomock.Controller
		var mockedRoute *mockroute.MockIMainRoute
		var index = 1

		BeforeEach(func() {
			mockedCtrl = gomock.NewController(GinkgoT())
			mockedRoute = mockroute.NewMockIMainRoute(mockedCtrl)

			sut = modifier{
				&info{
					&itinerary{route: mockedRoute},
				},
			}
		})

		AfterEach(func() {
			mockedCtrl.Finish()
		})

		It("should remove main stop from route", func() {
			mockedRoute.EXPECT().RemoveMainStop(index)
			sut.RemoveMainStopFromRoute(index)
		})
	})
})
