package decoder

import (
	"github.com/victorguarana/vehicle-routing/internal/brkga"
	decoderstrategy "github.com/victorguarana/vehicle-routing/internal/brkga/decoder/strategy"
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

func NewPositionalDecoderWithVehicleByStorage(carList []vehicle.ICar, gpsMap gps.Map) brkga.IDecoder[itinerary.ItineraryList] {
	return &positionDecoder{
		masterCarList: carList,
		gpsMap:        gpsMap,
		strategy:      decoderstrategy.NewVehicleChooserByStorage(gpsMap),
		name:          "PositionalDecoderWithVehicleByStorage",
	}
}

func NewPositionalDecoderWithVehicleByPercentage(carList []vehicle.ICar, gpsMap gps.Map, dronePercentage float64) brkga.IDecoder[itinerary.ItineraryList] {
	return &positionDecoder{
		masterCarList: carList,
		gpsMap:        gpsMap,
		strategy:      decoderstrategy.NewVehicleChooserByPercentage(gpsMap, dronePercentage),
		name:          "PositionalDecoderWithVehicleByPercentage",
	}
}

func NewTimeDecoderWithVehicleByStorage(carList []vehicle.ICar, gpsMap gps.Map) brkga.IDecoder[itinerary.ItineraryList] {
	return &timeWindowDecoder{
		masterCarList: carList,
		gpsMap:        gpsMap,
		strategy:      decoderstrategy.NewVehicleChooserByStorage(gpsMap),
		name:          "TimeDecoderWithVehicleByStorage",
	}
}

func NewTimeDecoderWithVehicleByPercentage(carList []vehicle.ICar, gpsMap gps.Map, dronePercentage float64) brkga.IDecoder[itinerary.ItineraryList] {
	return &timeWindowDecoder{
		masterCarList: carList,
		gpsMap:        gpsMap,
		strategy:      decoderstrategy.NewVehicleChooserByPercentage(gpsMap, dronePercentage),
		name:          "TimeDecoderWithVehicleByPercentage",
	}
}

func NewPositionalDecoderWithVehicleByPercentageV2(carList []vehicle.ICar, gpsMap gps.Map, dronePercentage float64) brkga.IDecoder[itinerary.ItineraryList] {
	return &positionDecoderAgatz{
		masterCarList: carList,
		gpsMap:        gpsMap,
		strategy:      decoderstrategy.NewVehicleChooserByPercentage(gpsMap, dronePercentage),
		name:          "PositionalAgatzDecoderWithVehicleByPercentage",
	}
}
