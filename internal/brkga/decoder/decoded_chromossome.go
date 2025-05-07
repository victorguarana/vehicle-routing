package decoder

import (
	"slices"
	"sort"

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

func orderDecodedChromossomes(decodedChromossomeList []*decodedChromossome) []*decodedChromossome {
	orderedDecodedchromossomeList := slices.Clone(decodedChromossomeList)
	sort.Slice(orderedDecodedchromossomeList, func(i, j int) bool {
		return *orderedDecodedchromossomeList[i].chromossome < *orderedDecodedchromossomeList[j].chromossome
	})
	return orderedDecodedchromossomeList
}
