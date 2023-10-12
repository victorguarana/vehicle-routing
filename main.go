package main

import (
	"fmt"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/greedy"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func main() {
	p := gps.Point{}
	car1 := vehicles.NewCar("vehicle 1", &p)
	car1.NewDrone("drone 1")

	route, err := greedy.ClosestNeighbor(car1, gps.GetMap())
	if err != nil {
		panic(err)
	}

	fmt.Print(route.String())
}
