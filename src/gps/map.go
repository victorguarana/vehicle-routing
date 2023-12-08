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
			{Name: "Dep#1", Latitude: -1, Longitude: 27},
			{Name: "Dep#2", Latitude: -14, Longitude: -43},
		},
		Clients: []*Point{
			{Name: "Cli#1", Latitude: 36, Longitude: 17, PackageSize: 8},
			{Name: "Cli#2", Latitude: -45, Longitude: -19, PackageSize: 2},
			{Name: "Cli#3", Latitude: 46, Longitude: -32, PackageSize: 4},
			{Name: "Cli#4", Latitude: -14, Longitude: -33, PackageSize: 4},
			{Name: "Cli#5", Latitude: -37, Longitude: 5, PackageSize: 2},
			{Name: "Cli#6", Latitude: 12, Longitude: 44, PackageSize: 5},
			{Name: "Cli#7", Latitude: 37, Longitude: -34, PackageSize: 7},
			{Name: "Cli#8", Latitude: -10, Longitude: 47, PackageSize: 6},
			{Name: "Cli#9", Latitude: -28, Longitude: 50, PackageSize: 1},
			{Name: "Cli#10", Latitude: 49, Longitude: -48, PackageSize: 8},
			{Name: "Cli#11", Latitude: -19, Longitude: -12, PackageSize: 4},
			{Name: "Cli#12", Latitude: -47, Longitude: 30, PackageSize: 6},
			{Name: "Cli#13", Latitude: -23, Longitude: -37, PackageSize: 5},
			{Name: "Cli#14", Latitude: -11, Longitude: -44, PackageSize: 4},
			{Name: "Cli#15", Latitude: -16, Longitude: 3, PackageSize: 6},
		},
	}
}

func GetMap() Map {
	return singletonMap
}
