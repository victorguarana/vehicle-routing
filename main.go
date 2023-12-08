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
	car1 := vehicles.NewCar("car 1", p)
	car1.NewDrone("drone 1.1")
	car1.NewDrone("drone 1.2")
	route1, _ := routes.NewRoute(car1)

	car2 := vehicles.NewCar("car 2", p)
	car2.NewDrone("drone 2.1")
	car2.NewDrone("drone 2.2")
	route2, _ := routes.NewRoute(car2)

	routesList := []routes.IRoute{route1, route2}

	err := greedy.BestInsertion(routesList, gps.GetMap())
	if err != nil {
		panic(err)
	}

	for i, r := range routesList {
		fmt.Printf("#{%d} - %s", i, r.String())
	}

	greedy.DroneSimpleInsertion(route1)

	fmt.Print(route1.String())
}
