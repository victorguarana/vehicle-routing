package decoder

import (
	"math"
	"slices"
	"sort"

	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

var _ brkga.IDecoder[itinerary.Itinerary] = (*positionDecoder)(nil)

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

	cachedGeneAmplifier float64
	cachedGeneModule    float64
}

func NewSimpleDecoder(carList []vehicle.ICar, gpsMap gps.Map) *positionDecoder {
	return &positionDecoder{
		masterCarList: carList,
		masterGPSMap:  gpsMap,
		gpsMap:        gpsMap,
	}
}

func (d *positionDecoder) Decode(individual *brkga.Individual) []itinerary.Itinerary {
	d.initializeDecoding(individual)
	d.processChromossomes()
	d.finalizeItineraries()

	return d.collectItineraries()
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
		car, drone := d.defineVehicle(chromossome)
		if car != nil {
			d.carByChromossome[chromossome] = car
		}
		if drone != nil {
			d.droneByChromossome[chromossome] = drone
		}
	}
}

func (d *positionDecoder) defineVehicle(chromossome *brkga.Chromossome) (vehicle.ICar, vehicle.IDrone) {
	modSum := 0.0
	amplifiedGene := chromossome.Gene() * d.geneAmplifier()
	moduledGene := math.Mod(amplifiedGene, d.geneModule())

	for _, car := range d.carList {
		modSum += car.Storage()
		if moduledGene < modSum {
			return car, nil
		}
	}

	for _, car := range d.carList {
		for _, drone := range car.Drones() {
			modSum += drone.Storage()
			if moduledGene < modSum {
				return car, drone
			}
		}
	}

	return nil, nil
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
	constructor.LandAllDrones(constructor.ActualCarStop())

	actualCustomerPoint := d.customerByChromossome[chromossome]
	constructor.MoveCar(actualCustomerPoint)
}

func (d *positionDecoder) geneAmplifier() float64 {
	if d.cachedGeneAmplifier == 0 {
		d.cachedGeneAmplifier = float64(len(d.gpsMap.Clients)) * d.calcTotalStorage()
	}

	return d.cachedGeneAmplifier
}

func (d *positionDecoder) geneModule() float64 {
	if d.cachedGeneModule == 0 {
		d.cachedGeneModule = d.calcTotalStorage()
	}

	return d.cachedGeneModule
}

func (d *positionDecoder) calcTotalStorage() float64 {
	totalStorage := 0.0
	for _, car := range d.carList {
		totalStorage += car.Storage()
		for _, drone := range car.Drones() {
			totalStorage += drone.Storage()
		}
	}
	return totalStorage
}
