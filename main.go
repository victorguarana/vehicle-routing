package main

import (
	"fmt"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/greedy"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func main() {
	p := gps.GetMap().Deposits[0]
	car1 := vehicles.NewCar("vehicle 1", p)
	car1.NewDrone("drone 1")
	car1.NewDrone("drone 2")
	route1, _ := routes.NewRoute(car1)

	err := greedy.ClosestNeighbor(route1, gps.GetMap())
	if err != nil {
		panic(err)
	}

	fmt.Printf("route1: %s\n", route1)

	greedy.DroneSimpleInsertion(route1)

	fmt.Print(route1.String())
}
