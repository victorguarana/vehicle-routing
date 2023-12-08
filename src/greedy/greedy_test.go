package greedy

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	mockroutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"
	"go.uber.org/mock/gomock"
)

var _ = Describe("closestPoint", func() {
	DescribeTable("when receive cadidates", func(candidatePoints []*gps.Point, expectedPoint *gps.Point) {
		originPoint := &gps.Point{Latitude: 0, Longitude: 0}
		receivedPoint := closestPoint(originPoint, candidatePoints)

		Expect(receivedPoint).To(Equal(expectedPoint))
	},
		Entry("when latitude is equal, return closest point", []*gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 1, Longitude: 2}}, &gps.Point{Latitude: 1, Longitude: 1}),
		Entry("when longitude is equal, return closest point", []*gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 2, Longitude: 1}}, &gps.Point{Latitude: 1, Longitude: 1}),
	)

	Context("when there are no candidate points", func() {
		It("return nil", func() {
			originPoint := &gps.Point{Latitude: 0, Longitude: 0}
			receivedPoint := closestPoint(originPoint, []*gps.Point{})

			Expect(receivedPoint).To(BeNil())
		})
	})
})

var _ = Describe("finishRoutes", func() {
	var (
		mockCtrl      *gomock.Controller
		mockedCar     *mockvehicles.MockICar
		mockedRoute   *mockroutes.MockIRoute
		mockedCarStop *mockroutes.MockICarStop

		routesList     []routes.IRoute
		closestDeposit *gps.Point
		m              gps.Map
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
		mockedRoute = mockroutes.NewMockIRoute(mockCtrl)
		mockedCarStop = mockroutes.NewMockICarStop(mockCtrl)

		routesList = []routes.IRoute{mockedRoute}
		closestDeposit = &gps.Point{Latitude: 1, Longitude: 1}
		m = gps.Map{Deposits: []*gps.Point{closestDeposit}}

		mockedRoute.EXPECT().Last().Return(mockedCarStop)
		mockedCarStop.EXPECT().Point().Return(closestDeposit)
		mockedRoute.EXPECT().Car().Return(mockedCar)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can support the route", func() {
		It("move the car to the closest deposit and append it to the route", func() {
			mockedCar.EXPECT().Move(closestDeposit).Return(nil)
			mockedRoute.EXPECT().Append(closestDeposit).Return(nil)

			Expect(finishRoutes(routesList, m)).To(Succeed())
		})
	})

	Context("when car can not support the route", func() {
		It("return an error", func() {
			mockedErr := errors.New("mocked error")
			mockedCar.EXPECT().Move(closestDeposit).Return(mockedErr)

			receivedErr := finishRoutes(routesList, m)

			Expect(receivedErr).To(MatchError(mockedErr))
		})
	})
})

var _ = Describe("moveAndAppend", func() {
	var mockCtrl *gomock.Controller
	var mockedRoute *mockroutes.MockIRoute
	var mockedCar *mockvehicles.MockICar

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
		mockedRoute = mockroutes.NewMockIRoute(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car can move to the point", func() {
		It("move the car and append the point to the route", func() {
			point := &gps.Point{Latitude: 1, Longitude: 1}
			mockedCar.EXPECT().Move(point).Return(nil)
			mockedRoute.EXPECT().Car().Return(mockedCar)
			mockedRoute.EXPECT().Append(point).Return(nil)

			receivedErr := moveAndAppend(mockedRoute, point)

			Expect(receivedErr).NotTo(HaveOccurred())
		})
	})

	Context("when car can not move to the point", func() {
		It("return an error", func() {
			point := &gps.Point{Latitude: 1, Longitude: 1}
			mockedCar.EXPECT().Move(point).Return(vehicles.ErrDestinationNotSupported)
			mockedRoute.EXPECT().Car().Return(mockedCar)

			err := moveAndAppend(mockedRoute, point)

			Expect(err).To(MatchError(vehicles.ErrDestinationNotSupported))
		})
	})

	Context("when route can not append point", func() {
		It("return an error", func() {
			point := &gps.Point{Latitude: 1, Longitude: 1}
			mockedErr := errors.New("mocked error")
			mockedCar.EXPECT().Move(point).Return(nil)
			mockedRoute.EXPECT().Car().Return(mockedCar)
			mockedRoute.EXPECT().Append(point).Return(mockedErr)

			err := moveAndAppend(mockedRoute, point)

			Expect(err).To(MatchError(mockedErr))
		})
	})
})
