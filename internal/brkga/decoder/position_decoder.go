package decoder

import (
	"errors"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/slc"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

var _ brkga.IDecoder[itinerary.ItineraryList] = (*positionDecoder)(nil)

type positionDecoder struct {
	masterCarList []vehicle.ICar
	gpsMap        gps.Map
	strategy      strategy
}

func (d *positionDecoder) Decode(individual *brkga.Individual) (itinerary.ItineraryList, error) {
	decodedChromossomeList := d.decodeChromossomeList(individual.Chromosomes)
	d.parseChromossomes(decodedChromossomeList)
	itineraryList := d.collectItineraries(decodedChromossomeList)

	finalizeItineraries(itineraryList, d.gpsMap)

	if !isValidSolution(itineraryList) {
		return nil, errors.New("Invalid Solution")
	}

	return itineraryList, nil
}

func (d *positionDecoder) decodeChromossomeList(chromossomeList []*brkga.Chromossome) []*decodedChromossome {
	clonedCarList := cloneCars(d.masterCarList)
	decodedChromossomeList := make([]*decodedChromossome, len(chromossomeList))
	itineraryByCar := mapItineraryByCar(clonedCarList)

	for i, chromossome := range chromossomeList {
		car, drone := d.strategy.DefineVehicle(clonedCarList, chromossome)
		decodedChromossome := &decodedChromossome{
			customer:    d.gpsMap.Clients[i],
			car:         car,
			drone:       drone,
			itn:         itineraryByCar[car],
			chromossome: chromossome,
		}
		decodedChromossomeList[i] = decodedChromossome

	}

	return orderDecodedChromossomes(decodedChromossomeList)
}

func (d *positionDecoder) parseChromossomes(decodedChromossomeList []*decodedChromossome) {
	for _, dc := range decodedChromossomeList {
		if dc.isDroneChromossome() {
			d.parseDecodedDroneChromossome(dc)
		} else {
			d.parseDecodedCarChromossome(dc)
		}
	}
}

func (*positionDecoder) parseDecodedDroneChromossome(dc *decodedChromossome) {
	drone := dc.drone
	constructor := dc.itn.Constructor()
	if !drone.IsFlying() {
		constructor.StartDroneFlight(drone, constructor.ActualCarStop())
	}

	actualCustomerPoint := dc.customer
	constructor.MoveDrone(drone, actualCustomerPoint)
}

func (*positionDecoder) parseDecodedCarChromossome(dc *decodedChromossome) {
	constructor := dc.itn.Constructor()

	actualCustomerPoint := dc.customer
	constructor.MoveCar(actualCustomerPoint)
	constructor.LandAllDrones(constructor.ActualCarStop())
}

func (*positionDecoder) collectItineraries(decodedChromossomeList []*decodedChromossome) []itinerary.Itinerary {
	itineraryList := []itinerary.Itinerary{}
	for _, dc := range decodedChromossomeList {
		itineraryList = slc.AppendIfNotExists(itineraryList, dc.itn)
	}
	return itineraryList
}
