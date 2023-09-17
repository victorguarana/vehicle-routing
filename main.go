package main

import (
	"fmt"

	"github.com/victorguarana/go-vehicle-route/gps"
	"github.com/victorguarana/go-vehicle-route/vehicles"
)

func main() {
	p := gps.Point{}
	v1 := vehicles.NewVehicle("vehicle 1", &p)

	fmt.Println(v1)
}
