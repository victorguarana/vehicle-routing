package greedy

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes/mockroutes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
	"github.com/victorguarana/go-vehicle-route/src/vehicles/mockvehicles"
)

var _ = Describe("ClosestNeighbor", func() {
	var mockCtrl *gomock.Controller
	var mockedCar *mockvehicles.MockICar

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when car supports entire route", func() {
		It("return a route without deposits between clients", func() {
			initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
			client1 := &gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
			client2 := &gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
			deposit1 := &gps.Point{Latitude: 3, Longitude: 3}
			deposit2 := &gps.Point{Latitude: 4, Longitude: 4}

			m := gps.Map{
				Clients:  []*gps.Point{client1, client2},
				Deposits: []*gps.Point{deposit1, deposit2},
			}

			mockedCar.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar.EXPECT().Support(*client1, *deposit1).Return(true)
			mockedCar.EXPECT().Move(client1).Return(nil)

			mockedCar.EXPECT().ActualPosition().Return(client1)
			mockedCar.EXPECT().Support(*client2, *deposit2).Return(true)
			mockedCar.EXPECT().Move(client2).Return(nil)
			mockedCar.EXPECT().Move(deposit2).Return(nil)

			receivedRoute, receivedErr := ClosestNeighbor(mockedCar, m)

			expectedRoute := []*gps.Point{
				initialPoint,
				client1,
				client2,
				deposit2,
			}

			Expect(receivedErr).NotTo(HaveOccurred())
			Expect(receivedRoute.CompleteRoute()).To(Equal(expectedRoute))
		})
	})

	Context("when car needs to stop between clients", func() {
		It("return a route with deposits between clients", func() {
			initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
			client1 := &gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
			client2 := &gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
			deposit1 := &gps.Point{Latitude: 3, Longitude: 3}
			deposit2 := &gps.Point{Latitude: 4, Longitude: 4}

			m := gps.Map{
				Clients:  []*gps.Point{client1, client2},
				Deposits: []*gps.Point{deposit1, deposit2},
			}

			mockedCar.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar.EXPECT().Support(*client1, *deposit1).Return(true)
			mockedCar.EXPECT().Move(client1).Return(nil)

			mockedCar.EXPECT().ActualPosition().Return(client1)
			mockedCar.EXPECT().Support(*client2, *deposit2).Return(false)
			mockedCar.EXPECT().Move(deposit1).Return(nil)

			mockedCar.EXPECT().ActualPosition().Return(deposit1)
			mockedCar.EXPECT().Support(*client2, *deposit2).Return(true)
			mockedCar.EXPECT().Move(client2).Return(nil)
			mockedCar.EXPECT().Move(deposit2).Return(nil)

			receivedRoute, receivedErr := ClosestNeighbor(mockedCar, m)

			expectedRoute := []*gps.Point{
				initialPoint,
				client1,
				deposit1,
				client2,
				deposit2,
			}

			Expect(receivedErr).NotTo(HaveOccurred())
			Expect(receivedRoute.CompleteRoute()).To(Equal(expectedRoute))
		})
	})
})

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

var _ = Describe("removePoint", func() {
	var points []*gps.Point

	BeforeEach(func() {
		points = []*gps.Point{
			{Latitude: 1, Longitude: 1},
			{Latitude: 2, Longitude: 2},
			{Latitude: 3, Longitude: 3},
		}
	})

	Context("when point is in the slice", func() {
		Context("when point is the first element", func() {
			It("remove the point", func() {
				point := points[0]
				expectedPoints := []*gps.Point{{Latitude: 2, Longitude: 2}, {Latitude: 3, Longitude: 3}}
				receivedPoints := removePoint(points, point)

				Expect(receivedPoints).To(Equal(expectedPoints))
			})
		})

		Context("when point is in the middle", func() {
			It("remove the point", func() {
				point := points[1]
				expectedPoints := []*gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 3, Longitude: 3}}
				receivedPoints := removePoint(points, point)

				Expect(receivedPoints).To(Equal(expectedPoints))
			})
		})

		Context("when point is the last element", func() {
			It("remove the point", func() {
				point := points[2]
				expectedPoints := []*gps.Point{{Latitude: 1, Longitude: 1}, {Latitude: 2, Longitude: 2}}
				receivedPoints := removePoint(points, point)

				Expect(receivedPoints).To(Equal(expectedPoints))
			})
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
			mockedCar.EXPECT().Move(point).Return(vehicles.ErrWithoutRange)
			mockedRoute.EXPECT().Car().Return(mockedCar)

			err := moveAndAppend(mockedRoute, point)

			Expect(err).To(MatchError(vehicles.ErrWithoutRange))
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
