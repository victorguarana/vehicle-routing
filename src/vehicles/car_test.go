package vehicles

import (
	mockRoutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewCar", func() {
	Context("when car can be created", func() {
		var mockCtrl *gomock.Controller
		var mockedInitialStop *mockRoutes.MockIMainStop

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockedInitialStop = mockRoutes.NewMockIMainStop(mockCtrl)
		})

		AfterEach(func() {
			defer mockCtrl.Finish()
		})

		It("should create car with correct params", func() {
			carParams := CarParams{
				Name:          "car1",
				StartingPoint: mockedInitialStop,
			}

			receivedCar := NewCar(carParams)
			expectedCar := car{
				drones: []*drone{},
				name:   "car1",
				speed:  defaultCarSpeed,
			}

			Expect(receivedCar).To(Equal(&expectedCar))
		})
	})
})

var _ = Describe("car{}", func() {
	Describe("Drones", func() {
		var drone1 = &drone{}
		var drone2 = &drone{}
		var sut = &car{
			drones: []*drone{drone1, drone2},
		}

		It("should return all drones", func() {
			Expect(sut.Drones()).To(Equal([]IDrone{drone1, drone2}))
		})
	})

	Describe("Name", func() {
		var sut = &car{
			name: "car1",
		}

		It("should return car name", func() {
			Expect(sut.Name()).To(Equal("car1"))
		})
	})

	Describe("NewDrone", func() {
		var sut = &car{
			drones: []*drone{},
			name:   "car1",
		}

		It("should create new drone", func() {
			droneParams := DroneParams{
				Name: "drone1",
			}
			sut.NewDrone(droneParams)
			Expect(len(sut.Drones())).To(Equal(1))
		})
	})

	Describe("Speed", func() {
		var sut = &car{
			speed: defaultCarSpeed,
		}

		It("should return car speed", func() {
			Expect(sut.Speed()).To(Equal(defaultCarSpeed))
		})
	})
})
