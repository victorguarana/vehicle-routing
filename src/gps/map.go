package gps

var singletonMap Map

type Map struct {
	Deposits []*Point
	Clients  []*Point
}

func init() {
	singletonMap = loadMap()
}

func loadMap() Map {
	return Map{
		Deposits: []*Point{
			{Latitude: -1, Longitude: 27},
			{Latitude: -14, Longitude: -43},
		},
		Clients: []*Point{
			{Latitude: 36, Longitude: 17, PackageSize: 8},
			{Latitude: -45, Longitude: -19, PackageSize: 2},
			{Latitude: 46, Longitude: -32, PackageSize: 4},
			{Latitude: -14, Longitude: -33, PackageSize: 4},
			{Latitude: -37, Longitude: 5, PackageSize: 0},
			{Latitude: 12, Longitude: 44, PackageSize: 5},
			{Latitude: 37, Longitude: -34, PackageSize: 7},
			{Latitude: -10, Longitude: 47, PackageSize: 6},
			{Latitude: -28, Longitude: 50, PackageSize: 1},
			{Latitude: 49, Longitude: -48, PackageSize: 8},
			{Latitude: -19, Longitude: -12, PackageSize: 0},
			{Latitude: -47, Longitude: 30, PackageSize: 6},
			{Latitude: -23, Longitude: -37, PackageSize: 5},
			{Latitude: -11, Longitude: -44, PackageSize: 4},
			{Latitude: -16, Longitude: 3, PackageSize: 6},
		},
	}
}

func GetMap() Map {
	return singletonMap
}
