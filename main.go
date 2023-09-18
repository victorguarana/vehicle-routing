package main

import (
	"fmt"

	"github.com/victorguarana/go-vehicle-route/gps"
	"github.com/victorguarana/go-vehicle-route/vehicles"
)

func main() {
	p := gps.Point{}
	v1 := vehicles.NewCar("vehicle 1", &p)
	v1.NewDrone("drone 1")

	fmt.Println(v1)
}