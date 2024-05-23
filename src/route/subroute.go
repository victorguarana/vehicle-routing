package route

import (
	"log"

	"github.com/victorguarana/vehicle-routing/src/slc"
)

type ISubRoute interface {
	Append(ISubStop)
	First() ISubStop
	InsertAt(index int, iSubStop ISubStop)
	Iterator() slc.Iterator[ISubStop]
	Last() ISubStop
	Length() int
	Return(IMainStop)
	ReturningStop() IMainStop
	StartingStop() IMainStop
	RemoveSubStop(index int)
}

type subRoute struct {
	returningStop *mainStop
	startingStop  *mainStop
	stops         []*subStop
}

func NewSubRoute(iStartingStop IMainStop) ISubRoute {
	startingStop := iStartingStop.(*mainStop)
	sr := &subRoute{
		startingStop: startingStop,
		stops:        []*subStop{},
	}
	startingStop.startingSubRoutes = append(startingStop.startingSubRoutes, sr)
	return sr
}

func (sr *subRoute) Append(iSubStop ISubStop) {
	sr.stops = append(sr.stops, iSubStop.(*subStop))
}

func (sr *subRoute) First() ISubStop {
	return sr.stops[0]
}

func (sr *subRoute) InsertAt(index int, iSubStop ISubStop) {
	if index < 0 || index > len(sr.stops) {
		log.Printf("InserAt: index (%d) out of range\n", index)
		return
	}
	sr.stops = slc.InsertAt(sr.stops, iSubStop.(*subStop), index)
}

func (sr *subRoute) Iterator() slc.Iterator[ISubStop] {
	iSubStops := make([]ISubStop, len(sr.stops))
	for i, stop := range sr.stops {
		iSubStops[i] = stop
	}

	return slc.NewIterator(iSubStops)
}

func (sr *subRoute) Last() ISubStop {
	return sr.stops[len(sr.stops)-1]
}

func (sr *subRoute) Length() int {
	return len(sr.stops)
}

func (sr *subRoute) Return(iMainStop IMainStop) {
	sr.returningStop = iMainStop.(*mainStop)
}

func (sr *subRoute) ReturningStop() IMainStop {
	return sr.returningStop
}

func (sr *subRoute) StartingStop() IMainStop {
	return sr.startingStop
}

func (sr *subRoute) RemoveSubStop(index int) {
	if index < 0 || index >= len(sr.stops) {
		log.Printf("RemoveSubStop: index (%d) out of range\n", index)
		return
	}
	sr.stops = append(sr.stops[:index], sr.stops[index+1:]...)
}
