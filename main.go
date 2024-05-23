package main

import (
	"log"

	"github.com/victorguarana/vehicle-routing/src/csp"
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/greedy"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/measure"
	"github.com/victorguarana/vehicle-routing/src/output"
	"github.com/victorguarana/vehicle-routing/src/vehicle"
)

func main() {
	BestInsertion()
	BestInsertionWithDrones()

	ClosestNeighbor()
	ClosestNeighborWithDrones()

	Covering()
	CoveringMaxDrones()
}

func ClosestNeighbor() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	itn := itinerary.New(car)
	greedy.ClosestNeighbor([]itinerary.Itinerary{itn}, gps.GetMap())

	totalDistance := measure.TotalDistance(itn)
	totalTime := measure.TimeSpent(itn)
	totalFuelSpent := measure.SpentFuel(itn)
	log.Println("ClosestNeighbor: Total Distance:", totalDistance)
	log.Println("ClosestNeighbor: Total Time:", totalTime)
	log.Println("ClosestNeighbor: Total Fuel Spent:", totalFuelSpent)

	filename := "closest-neighbor.png"
	output.ToImage(filename, itn, totalDistance, totalTime)
}

func ClosestNeighborWithDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	greedy.ClosestNeighbor([]itinerary.Itinerary{itn}, gps.GetMap())
	greedy.DroneStrikesInsertion(itn)

	totalDistance := measure.TotalDistance(itn)
	totalTime := measure.TimeSpent(itn)
	totalFuelSpent := measure.SpentFuel(itn)
	log.Println("ClosestNeighborWithDrones: Total Distance:", totalDistance)
	log.Println("ClosestNeighborWithDrones: Total Time:", totalTime)
	log.Println("ClosestNeighborWithDrones: Total Fuel Spent:", totalFuelSpent)

	filename := "closest-neighbor-with-drones.png"
	output.ToImage(filename, itn, totalDistance, totalTime)
}

func BestInsertion() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	greedy.BestInsertion([]itinerary.Itinerary{itn}, gps.GetMap())

	totalDistance := measure.TotalDistance(itn)
	totalTime := measure.TimeSpent(itn)
	totalFuelSpent := measure.SpentFuel(itn)
	log.Println("BestiInsertion: Total Distance:", totalDistance)
	log.Println("BestiInsertion: Total Time:", totalTime)
	log.Println("BestiInsertion: Total Fuel Spent:", totalFuelSpent)

	filename := "best-insertion.png"
	output.ToImage(filename, itn, totalDistance, totalTime)
}

func BestInsertionWithDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	greedy.BestInsertion([]itinerary.Itinerary{itn}, gps.GetMap())
	greedy.DroneStrikesInsertion(itn)

	totalDistance := measure.TotalDistance(itn)
	totalTime := measure.TimeSpent(itn)
	totalFuelSpent := measure.SpentFuel(itn)
	log.Println("BestiInsertionWithDrones: Total Distance:", totalDistance)
	log.Println("BestiInsertionWithDrones: Total Time:", totalTime)
	log.Println("BestiInsertionWithDrones: Total Fuel Spent:", totalFuelSpent)

	filename := "best-insertion-with-drones.png"
	output.ToImage(filename, itn, totalDistance, totalTime)
}

func Covering() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicle.DroneRange / 4
	csp.CoveringWithDrones([]itinerary.Itinerary{itn}, gps.GetMap(), neighorhoodDistance)

	totalDistance := measure.TotalDistance(itn)
	totalTime := measure.TimeSpent(itn)
	totalFuelSpent := measure.SpentFuel(itn)
	log.Println("Covering: Total Distance:", totalDistance)
	log.Println("Covering: Total Time:", totalTime)
	log.Println("Covering: Total Fuel Spent:", totalFuelSpent)

	filename := "covering.png"
	output.ToImage(filename, itn, totalDistance, totalTime)
}

func CoveringMaxDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicle.DroneRange / 2
	csp.CoveringWithDrones([]itinerary.Itinerary{itn}, gps.GetMap(), neighorhoodDistance)

	totalDistance := measure.TotalDistance(itn)
	totalTime := measure.TimeSpent(itn)
	totalFuelSpent := measure.SpentFuel(itn)
	log.Println("CoveringMaxDrones: Total Distance:", totalDistance)
	log.Println("CoveringMaxDrones: Total Time:", totalTime)
	log.Println("CoveringMaxDrones: Total Fuel Spent:", totalFuelSpent)

	filename := "covering-max-drones.png"
	output.ToImage(filename, itn, totalDistance, totalTime)
}
