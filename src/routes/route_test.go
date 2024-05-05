package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewRoute", func() {
	var (
		mockCtrl  *gomock.Controller
		mockedCar *mockvehicles.MockICar
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
	})

	Context("when all params are valid", func() {
		It("returns new route", func() {
			intitialPosition := gps.Point{}
			expectedRoute := &route{
				car:   mockedCar,
				stops: []*carStop{newCarStop(mockedCar, intitialPosition)},
			}

			mockedCar.EXPECT().ActualPosition().Return(intitialPosition)

			receivedRoute, receivedErr := NewRoute(mockedCar)

			Expect(receivedErr).NotTo(HaveOccurred())
			Expect(receivedRoute).To(Equal(expectedRoute))
		})
	})

	Context("when car is nil", func() {
		It("returns nil car error", func() {
			receivedRoute, receivedErr := NewRoute(nil)

			Expect(receivedErr).To(MatchError(ErrNilCar))
			Expect(receivedRoute).To(BeNil())
		})
	})
})

var _ = Describe("Append", func() {
	var (
		mockCtrl  *gomock.Controller
		mockedCar *mockvehicles.MockICar

		sut route

		validPoint gps.Point
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)

		validPoint = gps.Point{}
		sut = route{
			car:   mockedCar,
			stops: []*carStop{},
		}
	})

	Context("when all params are valid", func() {
		It("appends new car stop", func() {
			expectedStop := newCarStop(mockedCar, validPoint)

			sut.Append(validPoint)
			Expect(sut.stops).To(Equal([]*carStop{expectedStop}))
		})
	})
})

var _ = Describe("AtIndex", func() {
	var (
		mockCtrl  *gomock.Controller
		mockedCar *mockvehicles.MockICar

		sut route
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)

		sut = route{
			car: mockedCar,
			stops: []*carStop{
				{point: gps.Point{Latitude: 0, Longitude: 0}, car: mockedCar},
				{point: gps.Point{Latitude: 1, Longitude: 1}, car: mockedCar},
			},
		}
	})

	Context("when index is valid", func() {
		It("returns car stop at index", func() {
			expectedStop := sut.stops[1]

			receivedStop, receivedErr := sut.AtIndex(1)

			Expect(receivedErr).NotTo(HaveOccurred())
			Expect(receivedStop).To(Equal(expectedStop))
		})
	})

	Context("when index is invalid", func() {
		It("returns index out of range error", func() {
			receivedStop, receivedErr := sut.AtIndex(2)

			Expect(receivedErr).To(MatchError(ErrIndexOutOfRange))
			Expect(receivedStop).To(BeNil())
		})
	})
})

var _ = Describe("RemoveCarStop", func() {
	var (
		mockCtrl  *gomock.Controller
		mockedCar *mockvehicles.MockICar

		sut route
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)

		sut = route{
			car: mockedCar,
			stops: []*carStop{
				{point: gps.Point{Latitude: 0, Longitude: 0}, car: mockedCar},
				{point: gps.Point{Latitude: 1, Longitude: 1}, car: mockedCar},
			},
		}
	})

	Context("when index is valid", func() {
		It("removes car stop at index", func() {
			expectedStops := []*carStop{
				{point: gps.Point{Latitude: 0, Longitude: 0}, car: mockedCar},
			}

			Expect(sut.RemoveCarStop(1)).To(Succeed())
			Expect(sut.stops).To(Equal(expectedStops))
		})
	})

	Context("when index is invalid", func() {
		It("returns index out of range error", func() {
			Expect(sut.RemoveCarStop(2)).To(MatchError(ErrIndexOutOfRange))
		})
	})
})
