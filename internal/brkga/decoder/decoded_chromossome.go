package decoder

import (
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

type decodedChromossome struct {
	customer    gps.Point
	car         vehicle.ICar
	drone       vehicle.IDrone
	itn         itinerary.Itinerary
	chromossome *brkga.Chromossome
}

func (d *decodedChromossome) isDroneChromossome() bool {
	return d.drone != nil
}
