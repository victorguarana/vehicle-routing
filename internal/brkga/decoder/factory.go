package decoder

import (
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	"github.com/victorguarana/vehicle-routing/internal/brkga/decoder/chooser"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

func NewPositionalDecoderWithVehicleByStorage(carList []vehicle.ICar, gpsMap gps.Map) brkga.IDecoder[itinerary.ItineraryList] {
	return &positionDecoder{
		masterCarList:  carList,
		gpsMap:         gpsMap,
		vehicleChooser: chooser.NewVehicleChooserByStorage(gpsMap),
	}
}

func NewPositionalDecoderWithVehicleByPercentage(carList []vehicle.ICar, gpsMap gps.Map, dronePercentage float64) brkga.IDecoder[itinerary.ItineraryList] {
	return &positionDecoder{
		masterCarList:  carList,
		gpsMap:         gpsMap,
		vehicleChooser: chooser.NewVehicleChooserByPercentage(gpsMap, dronePercentage),
	}
}
