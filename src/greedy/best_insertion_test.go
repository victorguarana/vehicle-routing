package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	mockroutes "github.com/victorguarana/go-vehicle-route/src/routes/mocks"
	mockvehicles "github.com/victorguarana/go-vehicle-route/src/vehicles/mocks"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BestInsertion", func() {
	var (
		mockCtrl     *gomock.Controller
		mockedCar1   *mockvehicles.MockICar
		mockedCar2   *mockvehicles.MockICar
		mockedRoute1 *mockroutes.MockIRoute
		mockedRoute2 *mockroutes.MockIRoute

		routesList []routes.IRoute
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicles.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicles.NewMockICar(mockCtrl)
		mockedRoute1 = mockroutes.NewMockIRoute(mockCtrl)
		mockedRoute2 = mockroutes.NewMockIRoute(mockCtrl)

		routesList = []routes.IRoute{mockedRoute1, mockedRoute2}
	})

	Context("when car supports entire route", func() {
		It("return a route without deposits between clients", func() {
			initialPoint := &gps.Point{Latitude: 0, Longitude: 0}
			client1 := &gps.Point{Latitude: 1, Longitude: 1, PackageSize: 1}
			client2 := &gps.Point{Latitude: 2, Longitude: 2, PackageSize: 1}
			client3 := &gps.Point{Latitude: 3, Longitude: 3, PackageSize: 1}
			client4 := &gps.Point{Latitude: 4, Longitude: 4, PackageSize: 1}
			client5 := &gps.Point{Latitude: 5, Longitude: 5, PackageSize: 1}
			client6 := &gps.Point{Latitude: 6, Longitude: 6, PackageSize: 1}
			deposit1 := &gps.Point{Latitude: 0, Longitude: 0}
			deposit2 := &gps.Point{Latitude: 7, Longitude: 7}

			m := gps.Map{
				Clients:  []*gps.Point{client4, client2, client5, client1, client3, client6},
				Deposits: []*gps.Point{deposit1, deposit2},
			}

			mockedRoute1.EXPECT().Car().Return(mockedCar1).AnyTimes()
			mockedRoute2.EXPECT().Car().Return(mockedCar2).AnyTimes()

			mockedCar1.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar1.EXPECT().Support(client3, deposit1).Return(true)
			mockedCar1.EXPECT().Move(client3)
			mockedRoute1.EXPECT().Append(client3)

			mockedCar1.EXPECT().Support(client4, deposit2).Return(true)
			mockedCar1.EXPECT().Move(client4)
			mockedRoute1.EXPECT().Append(client4)

			mockedCar1.EXPECT().Support(client5, deposit2).Return(true)
			mockedCar1.EXPECT().Move(client5)
			mockedRoute1.EXPECT().Append(client5)

			mockedCar1.EXPECT().Move(deposit2)
			mockedRoute1.EXPECT().Append(deposit2)

			mockedCar2.EXPECT().ActualPosition().Return(initialPoint)
			mockedCar2.EXPECT().Support(client1, deposit1).Return(true)
			mockedCar2.EXPECT().Move(client1)
			mockedRoute2.EXPECT().Append(client1)

			mockedCar2.EXPECT().Support(client2, deposit1).Return(true)
			mockedCar2.EXPECT().Move(client2)
			mockedRoute2.EXPECT().Append(client2)

			mockedCar2.EXPECT().Support(client6, deposit2).Return(true)
			mockedCar2.EXPECT().Move(client6)
			mockedRoute2.EXPECT().Append(client6)

			mockedCar2.EXPECT().Move(deposit2)
			mockedRoute2.EXPECT().Append(deposit2)

			Expect(BestInsertion(routesList, m)).To(Succeed())
		})
	})
})

var _ = Describe("orderedClients", func() {
	var (
		mockCtrl     *gomock.Controller
		mockedCar1   *mockvehicles.MockICar
		mockedCar2   *mockvehicles.MockICar
		mockedRoute1 *mockroutes.MockIRoute
		mockedRoute2 *mockroutes.MockIRoute

		routesList []routes.IRoute
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockedCar1 = mockvehicles.NewMockICar(mockCtrl)
		mockedCar2 = mockvehicles.NewMockICar(mockCtrl)
		mockedRoute1 = mockroutes.NewMockIRoute(mockCtrl)
		mockedRoute2 = mockroutes.NewMockIRoute(mockCtrl)

		routesList = []routes.IRoute{mockedRoute1, mockedRoute2}

		mockedRoute1.EXPECT().Car().Return(mockedCar1)
		mockedRoute2.EXPECT().Car().Return(mockedCar2)

		mockedCar1.EXPECT().ActualPosition().Return(&gps.Point{Latitude: 0, Longitude: 0})
		mockedCar2.EXPECT().ActualPosition().Return(&gps.Point{Latitude: 0, Longitude: 0})
	})

	Context("when clients is empty", func() {
		It("return empty array", func() {
			orderedClients := orderedClientsByRoutes(routesList, []*gps.Point{})

			Expect(orderedClients).To(Equal([][]*gps.Point{nil, nil}))
		})
	})

	Context("when clients has more than one element", func() {
		It("return ordered clients", func() {
			client1 := &gps.Point{Latitude: 1, Longitude: 1}
			client2 := &gps.Point{Latitude: 2, Longitude: 2}
			client3 := &gps.Point{Latitude: 3, Longitude: 3}
			client4 := &gps.Point{Latitude: 4, Longitude: 4}
			client5 := &gps.Point{Latitude: 5, Longitude: 5}
			client6 := &gps.Point{Latitude: 6, Longitude: 6}

			clients := []*gps.Point{
				client5,
				client2,
				client4,
				client6,
				client1,
				client3,
			}

			expectedOrderedClients := [][]*gps.Point{
				{
					client1,
					client4,
					client5,
				},
				{
					client2,
					client3,
					client6,
				},
			}

			orderedClients := orderedClientsByRoutes(routesList, clients)

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
