package decoder

import (
	"errors"
	"slices"
	"sort"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

var _ brkga.IDecoder[itinerary.ItineraryList] = (*timeWindowDecoder)(nil)

type timeWindowDecoder struct {
	masterCarList []vehicle.ICar
	gpsMap        gps.Map
	strategy      strategy
}

func (d *timeWindowDecoder) Decode(individual *brkga.Individual) (itinerary.ItineraryList, error) {
	decodedChromossomeList := d.decodeChromossomeList(individual.Chromosomes)
	d.parseChromossomes(decodedChromossomeList)
	itineraryList := collectItineraries(decodedChromossomeList)

	finalizeItineraries(itineraryList, d.gpsMap)

	if !isValidSolution(itineraryList) {
		return nil, errors.New("Invalid Solution")
	}

	return itineraryList, nil
}

func (d *timeWindowDecoder) decodeChromossomeList(chromossomeList []*brkga.Chromossome) []*decodedChromossome {
	clonedCarList := cloneCars(d.masterCarList)
	decodedChromossomeList := make([]*decodedChromossome, len(chromossomeList))
	itineraryByCar := mapItineraryByCar(clonedCarList)

	for i, chromossome := range chromossomeList {
		car, drone := d.strategy.DefineVehicle(clonedCarList, chromossome)
		timeWindowIndex := d.strategy.DefineWindowTime(clonedCarList, chromossome)
		decodedChromossome := &decodedChromossome{
			customer:        d.gpsMap.Customers[i],
			car:             car,
			drone:           drone,
			itn:             itineraryByCar[car],
			chromossome:     chromossome,
			timeWindowIndex: timeWindowIndex,
		}
		decodedChromossomeList[i] = decodedChromossome
	}

	return decodedChromossomeList
}

func (d *timeWindowDecoder) parseChromossomes(decodedChromossomeList []*decodedChromossome) {
	decodedchromossomesByTimeWindow := d.mapDecodedChromossomeByTimeWindow(decodedChromossomeList)
	timeWindows := slc.Keys(decodedchromossomesByTimeWindow)
	orderedTimeWindows := d.orderTimeWindow(timeWindows)

	for _, timeWindowIndex := range orderedTimeWindows {
		decodedChromossomeList := decodedchromossomesByTimeWindow[timeWindowIndex]
		orderedDecodedChromossomeList := orderDecodedChromossomesByChromossome(decodedChromossomeList)
		decodedDroneChromossomeList := collectDroneChromossomes(orderedDecodedChromossomeList)
		decodedCarChromossomeList := collectCarChromossomes(orderedDecodedChromossomeList)

		for _, dc := range decodedDroneChromossomeList {
			d.parseDecodedDroneChromossome(dc)
		}

		for _, dc := range decodedCarChromossomeList {
			d.parseDecodedCarChromossome(dc)
		}

		itnList := collectItineraries(decodedCarChromossomeList)
		for _, itn := range itnList {
			itineraryConstuctor := itn.Constructor()
			itineraryConstuctor.LandAllDrones(itineraryConstuctor.ActualCarStop())
		}
	}
}

func (*timeWindowDecoder) parseDecodedDroneChromossome(dc *decodedChromossome) {
	drone := dc.drone
	constructor := dc.itn.Constructor()
	if !drone.IsFlying() {
		constructor.StartDroneFlight(drone, constructor.ActualCarStop())
	}

	actualCustomerPoint := dc.customer
	constructor.MoveDrone(drone, actualCustomerPoint)
}

func (*timeWindowDecoder) parseDecodedCarChromossome(dc *decodedChromossome) {
	constructor := dc.itn.Constructor()

	actualCustomerPoint := dc.customer
	constructor.MoveCar(actualCustomerPoint)
}

func (*timeWindowDecoder) orderTimeWindow(timeWindowList []int) []int {
	orderedTimeWindowList := slices.Clone(timeWindowList)
	sort.Slice(orderedTimeWindowList, func(i, j int) bool {
		return orderedTimeWindowList[i] < orderedTimeWindowList[j]
	})
	return orderedTimeWindowList
}

func (d *timeWindowDecoder) mapDecodedChromossomeByTimeWindow(decodedChromossomeList []*decodedChromossome) map[int][]*decodedChromossome {
	orderedDecodedchromossomeByTimeWindow := make(map[int][]*decodedChromossome)
	for _, dc := range decodedChromossomeList {
		orderedDecodedchromossomeByTimeWindow[dc.timeWindowIndex] = append(orderedDecodedchromossomeByTimeWindow[dc.timeWindowIndex], dc)
	}
	return orderedDecodedchromossomeByTimeWindow
}
