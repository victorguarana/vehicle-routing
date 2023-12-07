package routes

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	mockVehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"
)

var _ = Describe("NewFlight", Ordered, func() {
	var (
		mockCtrl    *gomock.Controller
		mockedDrone *mockVehicles.MockIDrone

		validCarStop *carStop
	)

	BeforeAll(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedDrone = mockVehicles.NewMockIDrone(mockCtrl)

		validCarStop = &carStop{point: &gps.Point{}}
	})

	Describe("valid params", func() {
		Context("when all params are valid", func() {
			It("takes off drone and returns correct struct", func() {
				expectedFlight := &flight{
					takeoffPoint: validCarStop,
					landingPoint: validCarStop,
					drone:        mockedDrone,
					stops:        []*droneStop{},
				}

				receivedFlight, receivedErr := NewFlight(mockedDrone, validCarStop, validCarStop)

				Expect(receivedErr).NotTo(HaveOccurred())
				Expect(receivedFlight).To(Equal(expectedFlight))
			})
		})

		Context("when only landing point is nil", func() {
			It("takes off drone and returns correct struct", func() {
				expectedFlight := &flight{
					takeoffPoint: validCarStop,
					landingPoint: nil,
					drone:        mockedDrone,
					stops:        []*droneStop{},
				}

				receivedFlight, receivedErr := NewFlight(mockedDrone, validCarStop, nil)

				Expect(receivedErr).NotTo(HaveOccurred())
				Expect(receivedFlight).To(Equal(expectedFlight))
			})
		})
	})

	Describe("invalid params", func() {
		Context("when drone is nil", func() {
			It("returns nil and error", func() {
				receivedFlight, receivedErr := NewFlight(nil, validCarStop, nil)

				Expect(receivedErr).To(MatchError(ErrNilDrone))
				Expect(receivedFlight).To(BeNil())
			})
		})

		Context("when takeoff point is invalid", func() {
			It("returns nil and error", func() {
				receivedFlight, receivedErr := NewFlight(mockedDrone, nil, nil)

				Expect(receivedErr).To(MatchError(ErrInvalidTakeoffPoint))
				Expect(receivedFlight).To(BeNil())
			})
		})
	})

})

var _ = Describe("Append", Ordered, func() {
	var (
		sut flight

		validPoint *gps.Point
	)

	BeforeAll(func() {
		validPoint = &gps.Point{}
		sut = flight{
			takeoffPoint: &carStop{point: validPoint},
			landingPoint: &carStop{point: validPoint},
		}
	})

	Context("when drone stop is valid", func() {
		It("appends dronestop to flight", func() {
			expectedDroneStop := newDroneStop(sut.drone, validPoint, &sut)

			Expect(sut.Append(validPoint)).To(Succeed())
			Expect(sut.stops).To(Equal([]*droneStop{expectedDroneStop}))
		})
	})
})

var _ = Describe("Land", Ordered, func() {
	var (
		sut flight
	)

	BeforeAll(func() {
		sut = flight{
			takeoffPoint: &carStop{point: &gps.Point{}},
			landingPoint: &carStop{point: &gps.Point{}},
		}
	})

	Context("when landing point is valid", func() {
		It("lands drone and sets landing point", func() {
			paramLandingPoint := &carStop{point: &gps.Point{}}
			receivedErr := sut.Land(paramLandingPoint)

			Expect(receivedErr).NotTo(HaveOccurred())
			Expect(sut.landingPoint).To(Equal(paramLandingPoint))
		})
	})

	Context("when landing point is invalid", func() {
		It("returns invalid car stop error", func() {
			receivedErr := sut.Land(nil)

			Expect(receivedErr).To(MatchError(ErrInvalidCarStop))
		})
	})
})
