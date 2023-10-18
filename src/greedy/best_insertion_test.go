package greedy

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles/mockvehicles"
)

var _ = Describe("BestInsertion", func() {
	var mockCtrl *gomock.Controller
	var mockedCar *mockvehicles.MockICar

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar = mockvehicles.NewMockICar(mockCtrl)
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
			mockedCar.EXPECT().Support(*client1, *deposit1).Return(true)
			mockedCar.EXPECT().Move(client1).Return(nil)

			mockedCar.EXPECT().Support(*client2, *deposit2).Return(true)
			mockedCar.EXPECT().Move(client2).Return(nil)

			mockedCar.EXPECT().Move(deposit2).Return(nil)

			receivedRoute, receivedErr := BestInsertion(mockedCar, m)

			expectedRoute := []*gps.Point{
				initialPoint,
				client1,
				client2,
				deposit2,
			}

			Expect(receivedRoute.CompleteRoute()).To(Equal(expectedRoute))
			Expect(receivedErr).To(BeNil())
		})
	})
})

var _ = Describe("orderedClients", func() {
	Context("when clients is empty", func() {
		It("return empty array", func() {
			orderedClients := orderedClients(&gps.Point{}, []*gps.Point{})

			Expect(orderedClients).To(BeEmpty())
		})
	})

	Context("when clients has more than one element", func() {
		It("return ordered clients", func() {
			initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
			client1 := &gps.Point{Latitude: 1, Longitude: 1}
			client2 := &gps.Point{Latitude: 2, Longitude: 2}
			client3 := &gps.Point{Latitude: 3, Longitude: 3}
			client4 := &gps.Point{Latitude: 4, Longitude: 4}
			client5 := &gps.Point{Latitude: 5, Longitude: 5}

			clients := []*gps.Point{
				client5,
				client2,
				client4,
				client1,
				client3,
			}

			expectedOrderedClients := []*gps.Point{
				client1,
				client2,
				client3,
				client4,
				client5,
			}

			orderedClients := orderedClients(initialPoint, clients)

			Expect(orderedClients).To(Equal(expectedOrderedClients))
		})
	})
})

var _ = Describe("findBestInsertionIndex", func() {
	Context("when orderedClients is empty", func() {
		It("return 0", func() {
			initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
			client := &gps.Point{Latitude: 1, Longitude: 1}
			orderedClients := []*gps.Point{}

			bestIndex := findBestInsertionIndex(initialPoint, client, orderedClients)

			Expect(bestIndex).To(Equal(0))
		})
	})

	Context("when orderedClients has more than one element", func() {
		Context("when best insertion is the first position", func() {
			It("return 0", func() {
				initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
				client := &gps.Point{Latitude: 1, Longitude: 1}
				orderedClients := []*gps.Point{
					{Latitude: 2, Longitude: 2},
					{Latitude: 3, Longitude: 3},
					{Latitude: 4, Longitude: 4},
					{Latitude: 5, Longitude: 5},
				}

				bestIndex := findBestInsertionIndex(initialPoint, client, orderedClients)

				Expect(bestIndex).To(Equal(0))
			})
		})

		Context("when best insertion is in the middle", func() {
			It("return 2", func() {
				initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
				client := &gps.Point{Latitude: 3, Longitude: 3}
				orderedClients := []*gps.Point{
					{Latitude: 1, Longitude: 1},
					{Latitude: 2, Longitude: 2},
					{Latitude: 4, Longitude: 4},
					{Latitude: 5, Longitude: 5},
				}

				bestIndex := findBestInsertionIndex(initialPoint, client, orderedClients)

				Expect(bestIndex).To(Equal(2))
			})
		})

		Context("when best insertion is at the end", func() {
			It("return 4", func() {
				initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
				client := &gps.Point{Latitude: 5, Longitude: 5}
				orderedClients := []*gps.Point{
					{Latitude: 1, Longitude: 1},
					{Latitude: 2, Longitude: 2},
					{Latitude: 3, Longitude: 3},
					{Latitude: 4, Longitude: 4},
				}

				bestIndex := findBestInsertionIndex(initialPoint, client, orderedClients)

				Expect(bestIndex).To(Equal(4))
			})
		})

		Context("when new client is behind initial point", func() {
			It("return 0", func() {
				initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
				client := &gps.Point{Latitude: -2, Longitude: -2}
				orderedClients := []*gps.Point{
					{Latitude: 1, Longitude: 1},
					{Latitude: 2, Longitude: 2},
					{Latitude: 3, Longitude: 3},
					{Latitude: 4, Longitude: 4},
				}

				bestIndex := findBestInsertionIndex(initialPoint, client, orderedClients)

				Expect(bestIndex).To(Equal(0))
			})
		})

		Context("when new client is between initial point and first client", func() {
			It("return 0", func() {
				initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
				client := &gps.Point{Latitude: 4, Longitude: 4}
				orderedClients := []*gps.Point{
					{Latitude: 5, Longitude: 5},
				}

				bestIndex := findBestInsertionIndex(initialPoint, client, orderedClients)

				Expect(bestIndex).To(Equal(0))
			})
		})
	})
})

var _ = Describe("insertAt", func() {
	Context("when index is 0", func() {
		It("insert at the begining", func() {
			client := &gps.Point{Latitude: 1, Longitude: 1}
			orderedClients := []*gps.Point{
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
			}

			expectedOrderedClients := []*gps.Point{
				client,
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
			}

			receivedOrderedClients := insertAt(0, client, orderedClients)

			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when index is 1", func() {
		It("insert at the middle", func() {
			client := &gps.Point{Latitude: 1, Longitude: 1}
			orderedClients := []*gps.Point{
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
			}

			expectedOrderedClients := []*gps.Point{
				{Latitude: 2, Longitude: 2},
				client,
				{Latitude: 3, Longitude: 3},
			}

			receivedOrderedClients := insertAt(1, client, orderedClients)

			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})

	Context("when index is 2", func() {
		It("insert at the end", func() {
			client := &gps.Point{Latitude: 1, Longitude: 1}
			orderedClients := []*gps.Point{
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
			}

			expectedOrderedClients := []*gps.Point{
				{Latitude: 2, Longitude: 2},
				{Latitude: 3, Longitude: 3},
				client,
			}

			receivedOrderedClients := insertAt(2, client, orderedClients)

			Expect(receivedOrderedClients).To(Equal(expectedOrderedClients))
		})
	})
})
