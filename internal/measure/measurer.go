package measure

import "github.com/victorguarana/vehicle-routing/internal/itinerary"

type Measurer struct {
	measureFunc func(itinerary.Info) float64
}

func NewMeasurer(measureFunc func(itinerary.Info) float64) Measurer {
	return Measurer{
		measureFunc: measureFunc,
	}
}

func (m Measurer) Measure(list itinerary.ItineraryList) float64 {
	totalMeasure := 0.0
	for _, t := range list {
		totalMeasure += m.measureFunc(t.Info())
	}

	return totalMeasure
}
