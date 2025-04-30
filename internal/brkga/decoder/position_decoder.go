package decoder

import (
	"errors"
	"slices"
	"sort"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

var _ brkga.IDecoder[itinerary.ItineraryList] = (*positionDecoder)(nil)

type positionDecoder struct {
	masterCarList []vehicle.ICar
	masterGPSMap  gps.Map

	carList []vehicle.ICar
	gpsMap  gps.Map

	individual *brkga.Individual

	customerByChromossome map[*brkga.Chromossome]gps.Point
	orderedChromossomes   []*brkga.Chromossome
	itineraryByCar        map[vehicle.ICar]itinerary.Itinerary
	itineraryByDrone      map[vehicle.IDrone]itinerary.Itinerary
	carByChromossome      map[*brkga.Chromossome]vehicle.ICar
	droneByChromossome    map[*brkga.Chromossome]vehicle.IDrone

	vehicleChooser vehicleChooser
}

func NewPositionalDecoder(carList []vehicle.ICar, gpsMap gps.Map, vehicleChooser vehicleChooser) *positionDecoder {
	return &positionDecoder{
		masterCarList:  carList,
		masterGPSMap:   gpsMap,
		gpsMap:         gpsMap,
		vehicleChooser: vehicleChooser,
	}
}

func (d *positionDecoder) Decode(individual *brkga.Individual) (itinerary.ItineraryList, error) {
	d.initializeDecoding(individual)
	d.processChromossomes()
	d.finalizeItineraries()

	if !d.isValidSolution() {
		return nil, errors.New("Invalid Solution")
	}

	return d.collectItineraries(), nil
}

func (d *positionDecoder) initializeDecoding(individual *brkga.Individual) {
	d.individual = individual
	d.cloneCars()
	d.mapCustomerByChromossome()
	d.orderChromossomes()
	d.mapItineraryByVehicles()
	d.mapChromossomeByVehicle()
}

func (d *positionDecoder) processChromossomes() {
	for _, chromossome := range d.orderedChromossomes {
		if d.isDroneChromossome(chromossome) {
			d.decodeDroneChromossome(chromossome)
		} else {
			d.decodeCarChromossome(chromossome)
		}
	}
}

func (d *positionDecoder) finalizeItineraries() {
	for car, itinerary := range d.itineraryByCar {
		constructor := itinerary.Constructor()
		constructor.MoveCar(d.closestWarehouse(car))
		constructor.LandAllDrones(constructor.ActualCarStop())
	}
}

func (d *positionDecoder) collectItineraries() []itinerary.Itinerary {
	var itineraryList []itinerary.Itinerary
	for _, itn := range d.itineraryByCar {
		itineraryList = append(itineraryList, itn)
	}
	return itineraryList
}

func (d *positionDecoder) cloneCars() {
	d.carList = make([]vehicle.ICar, len(d.masterCarList))
	for i, c := range d.masterCarList {
		d.carList[i] = c.Clone()
	}
}

func (d *positionDecoder) mapItineraryByVehicles() {
	d.itineraryByCar = make(map[vehicle.ICar]itinerary.Itinerary)
	d.itineraryByDrone = make(map[vehicle.IDrone]itinerary.Itinerary)
	for _, car := range d.carList {
		itn := itinerary.New(car)
		d.itineraryByCar[car] = itn
		for _, drone := range car.Drones() {
			d.itineraryByDrone[drone] = itn
		}
	}
}

func (d *positionDecoder) mapCustomerByChromossome() {
	d.customerByChromossome = make(map[*brkga.Chromossome]gps.Point, len(d.gpsMap.Clients))

	for i, customer := range d.gpsMap.Clients {
		chromossome := d.individual.Chromosomes[i]
		d.customerByChromossome[chromossome] = customer
	}
}

func (d *positionDecoder) orderChromossomes() {
	chromossomeList := slices.Clone(d.individual.Chromosomes)
	sort.Slice(chromossomeList, func(i, j int) bool {
		return *chromossomeList[i] < *chromossomeList[j]
	})
	d.orderedChromossomes = chromossomeList
}

func (d *positionDecoder) mapChromossomeByVehicle() {
	d.carByChromossome = make(map[*brkga.Chromossome]vehicle.ICar)
	d.droneByChromossome = make(map[*brkga.Chromossome]vehicle.IDrone)

	for _, chromossome := range d.individual.Chromosomes {
		car, drone := d.vehicleChooser.DefineVehicle(d.carList, chromossome)
		if car != nil {
			d.carByChromossome[chromossome] = car
		}
		if drone != nil {
			d.droneByChromossome[chromossome] = drone
		}
	}
}

func (d *positionDecoder) closestWarehouse(car vehicle.ICar) gps.Point {
	return gps.ClosestPoint(car.ActualPoint(), d.gpsMap.Warehouses)
}

func (d *positionDecoder) isDroneChromossome(chromossome *brkga.Chromossome) bool {
	_, ok := d.droneByChromossome[chromossome]
	return ok
}

func (d *positionDecoder) decodeDroneChromossome(chromossome *brkga.Chromossome) {
	drone := d.droneByChromossome[chromossome]
	constructor := d.itineraryByDrone[drone].Constructor()
	if !drone.IsFlying() {
		constructor.StartDroneFlight(drone, constructor.ActualCarStop())
	}

	actualCustomerPoint := d.customerByChromossome[chromossome]
	constructor.MoveDrone(drone, actualCustomerPoint)
}

func (d *positionDecoder) decodeCarChromossome(chromossome *brkga.Chromossome) {
	car := d.carByChromossome[chromossome]
	constructor := d.itineraryByCar[car].Constructor()

	actualCustomerPoint := d.customerByChromossome[chromossome]
	constructor.MoveCar(actualCustomerPoint)
	constructor.LandAllDrones(constructor.ActualCarStop())
}

func (d *positionDecoder) isValidSolution() bool {
	for _, itn := range d.itineraryByCar {
		if !itn.Validator().IsValid() {
			return false
		}
	}

	return true
}
