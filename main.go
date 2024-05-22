package main

import (
	"github.com/victorguarana/vehicle-routing/src/csp"
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/greedy"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/output"
	"github.com/victorguarana/vehicle-routing/src/vehicles"
)

func main() {
	BestInsertion()
	BestInsertionWithDrones()

	ClosestNeighbor()
	ClosestNeighborWithDrones()

	Covering()
}

func ClosestNeighbor() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	itn := itinerary.New(car)
	greedy.ClosestNeighbor([]itinerary.Itinerary{itn}, gps.GetMap())

	filename := "closest-neighbor.png"
	output.ToImage(filename, itn)
}

func ClosestNeighborWithDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	greedy.ClosestNeighbor([]itinerary.Itinerary{itn}, gps.GetMap())
	greedy.DroneStrikesInsertion(itn)

	filename := "closest-neighbor-with-drones.png"
	output.ToImage(filename, itn)
}

func BestInsertion() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	greedy.BestInsertion([]itinerary.Itinerary{itn}, gps.GetMap())

	filename := "best-insertion.png"
	output.ToImage(filename, itn)
}

func BestInsertionWithDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	greedy.BestInsertion([]itinerary.Itinerary{itn}, gps.GetMap())
	greedy.DroneStrikesInsertion(itn)

	filename := "best-insertion-with-drones.png"
	output.ToImage(filename, itn)
}

func Covering() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicles.DroneRange / 4
	csp.CoveringWithDrones([]itinerary.Itinerary{itn}, gps.GetMap(), neighorhoodDistance)

	filename := "covering.png"
	output.ToImage(filename, itn)
}
