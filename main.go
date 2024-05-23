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
	constructor := itn.Constructor()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gps.GetMap())

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("ClosestNeighbor: Total Distance:", totalDistance)
	log.Println("ClosestNeighbor: Total Time:", totalTime)
	log.Println("ClosestNeighbor: Total Fuel Spent:", totalFuelSpent)

	filename := "closest-neighbor.png"
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func ClosestNeighborWithDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	greedy.ClosestNeighbor([]itinerary.Constructor{constructor}, gps.GetMap())
	greedy.DroneStrikesInsertion(constructor, modifier)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("ClosestNeighborWithDrones: Total Distance:", totalDistance)
	log.Println("ClosestNeighborWithDrones: Total Time:", totalTime)
	log.Println("ClosestNeighborWithDrones: Total Fuel Spent:", totalFuelSpent)

	filename := "closest-neighbor-with-drones.png"
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func BestInsertion() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gps.GetMap())

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("BestInsertion: Total Distance:", totalDistance)
	log.Println("BestInsertion: Total Time:", totalTime)
	log.Println("BestInsertion: Total Fuel Spent:", totalFuelSpent)

	filename := "best-insertion.png"
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func BestInsertionWithDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	constructor := itn.Constructor()
	modifier := itn.Modifier()
	greedy.BestInsertion([]itinerary.Constructor{constructor}, gps.GetMap())
	greedy.DroneStrikesInsertion(constructor, modifier)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("BestInsertionWithDrones: Total Distance:", totalDistance)
	log.Println("BestInsertionWithDrones: Total Time:", totalTime)
	log.Println("BestInsertionWithDrones: Total Fuel Spent:", totalFuelSpent)

	filename := "best-insertion-with-drones.png"
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func Covering() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicle.DroneRange / 4
	constructor := itn.Constructor()
	csp.CoveringWithDrones([]itinerary.Constructor{constructor}, gps.GetMap(), neighorhoodDistance)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("Covering: Total Distance:", totalDistance)
	log.Println("Covering: Total Time:", totalTime)
	log.Println("Covering: Total Fuel Spent:", totalFuelSpent)

	filename := "covering.png"
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}

func CoveringMaxDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicle.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicle.DroneRange / 2
	constructor := itn.Constructor()
	csp.CoveringWithDrones([]itinerary.Constructor{constructor}, gps.GetMap(), neighorhoodDistance)

	itnInfo := itn.Info()
	totalDistance := measure.TotalDistance(itnInfo)
	totalTime := measure.TimeSpent(itnInfo)
	totalFuelSpent := measure.SpentFuel(itnInfo)
	log.Println("CoveringMaxDrones: Total Distance:", totalDistance)
	log.Println("CoveringMaxDrones: Total Time:", totalTime)
	log.Println("CoveringMaxDrones: Total Fuel Spent:", totalFuelSpent)

	filename := "covering-max-drones.png"
	output.ToImage(filename, itnInfo, totalDistance, totalTime)
}
