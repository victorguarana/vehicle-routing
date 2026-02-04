package decoder

import (
	"errors"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

var _ brkga.IDecoder[itinerary.ItineraryList] = (*positionDecoderAgatz)(nil)

// This decoder is based on Position Decoder,
// but with the restriction that drones can only do one delivery per flight.
// Also allows that drones can be launched and landed at the same car stop,
// following the Mother ship strategy.
type positionDecoderAgatz struct {
	masterCarList []vehicle.ICar
	gpsMap        gps.Map
	strategy      strategy
	name          string
}

func (d *positionDecoderAgatz) Name() string {
	return d.name
}

func (d *positionDecoderAgatz) Decode(individual *brkga.Individual) (itinerary.ItineraryList, error) {
	decodedChromossomeList := d.decodeChromossomeList(individual.Chromosomes)
	orderedDecodedChromossomeList := orderDecodedChromossomesByChromossome(decodedChromossomeList)
	d.parseChromossomes(orderedDecodedChromossomeList)

	itineraryList := collectItineraries(orderedDecodedChromossomeList)

	finalizeItineraries(itineraryList, d.gpsMap)

	if !isValidSolution(itineraryList) {
		return nil, errors.New("Invalid Solution")
	}

	return itineraryList, nil
}

func (d *positionDecoderAgatz) decodeChromossomeList(chromossomeList []*brkga.Chromossome) []*decodedChromossome {
	clonedCarList := cloneCars(d.masterCarList)
	decodedChromossomeList := make([]*decodedChromossome, len(chromossomeList))
	itineraryByCar := mapItineraryByCar(clonedCarList)

	for i, chromossome := range chromossomeList {
		car, drone := d.strategy.DefineVehicle(clonedCarList, chromossome)
		decodedChromossome := &decodedChromossome{
			customer:    d.gpsMap.Customers[i],
			car:         car,
			drone:       drone,
			itn:         itineraryByCar[car],
			chromossome: chromossome,
		}
		decodedChromossomeList[i] = decodedChromossome

	}

	return decodedChromossomeList
}

func (d *positionDecoderAgatz) parseChromossomes(decodedChromossomeList []*decodedChromossome) {
	for _, dc := range decodedChromossomeList {
		if dc.isDroneChromossome() {
			d.parseDecodedDroneChromossome(dc)
		} else {
			d.parseDecodedCarChromossome(dc)
		}
	}
}

func (*positionDecoderAgatz) parseDecodedDroneChromossome(dc *decodedChromossome) {
	drone := dc.drone
	constructor := dc.itn.Constructor()
	if drone.IsFlying() {
		constructor.LandDrone(drone, constructor.ActualCarStop())
	}

	constructor.StartDroneFlight(drone, constructor.ActualCarStop())
	actualCustomerPoint := dc.customer
	constructor.MoveDrone(drone, actualCustomerPoint)
	constructor.LandDrone(drone, constructor.ActualCarStop())
}

func (*positionDecoderAgatz) parseDecodedCarChromossome(dc *decodedChromossome) {
	constructor := dc.itn.Constructor()

	actualCustomerPoint := dc.customer
	constructor.MoveCar(actualCustomerPoint)
	constructor.LandAllDrones(constructor.ActualCarStop())
}
