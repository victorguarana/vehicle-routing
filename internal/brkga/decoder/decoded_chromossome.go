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
	customer        gps.Point
	car             vehicle.ICar
	drone           vehicle.IDrone
	itn             itinerary.Itinerary
	chromossome     *brkga.Chromossome
	timeWindowIndex int
}

func (d *decodedChromossome) isDroneChromossome() bool {
	return d.drone != nil
}

func (d *decodedChromossome) isCarChromossome() bool {
	return d.drone == nil
}

func orderDecodedChromossomesByChromossome(decodedChromossomeList []*decodedChromossome) []*decodedChromossome {
	orderedDecodedchromossomeList := slices.Clone(decodedChromossomeList)
	sort.Slice(orderedDecodedchromossomeList, func(i, j int) bool {
		return *orderedDecodedchromossomeList[i].chromossome < *orderedDecodedchromossomeList[j].chromossome
	})
	return orderedDecodedchromossomeList
}

func collectDroneChromossomes(decodedChromossomeList []*decodedChromossome) []*decodedChromossome {
	droneChromossomeList := []*decodedChromossome{}
	for _, dc := range decodedChromossomeList {
		if dc.isDroneChromossome() {
			droneChromossomeList = append(droneChromossomeList, dc)
		}
	}
	return droneChromossomeList
}

func collectCarChromossomes(decodedChromossomeList []*decodedChromossome) []*decodedChromossome {
	carChromossomeList := []*decodedChromossome{}
	for _, dc := range decodedChromossomeList {
		if dc.isCarChromossome() {
			carChromossomeList = append(carChromossomeList, dc)
		}
	}
	return carChromossomeList
}
