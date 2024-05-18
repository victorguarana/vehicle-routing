package routes

import "github.com/victorguarana/go-vehicle-route/src/slc"

type ISubRoute interface {
	Append(ISubStop)
	First() ISubStop
	Iterator() slc.Iterator[ISubStop]
	Last() ISubStop
	Return(IMainStop)
	ReturningPoint() IMainStop
	StartingPoint() IMainStop
}

type subRoute struct {
	returningPoint *mainStop
	startingPoint  *mainStop
	stops          []*subStop
}

func NewSubRoute(iStartingPoint IMainStop) ISubRoute {
	startingPoint := iStartingPoint.(*mainStop)
	sr := &subRoute{
		startingPoint: startingPoint,
		stops:         []*subStop{},
	}
	startingPoint.startingSubRoutes = append(startingPoint.startingSubRoutes, sr)
	return sr
}

func (sr *subRoute) Append(iSubStop ISubStop) {
	sr.stops = append(sr.stops, iSubStop.(*subStop))
}

func (sr *subRoute) First() ISubStop {
	return sr.stops[0]
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

func (sr *subRoute) Return(iMainStop IMainStop) {
	sr.returningPoint = iMainStop.(*mainStop)
}

func (sr *subRoute) ReturningPoint() IMainStop {
	return sr.returningPoint
}

func (sr *subRoute) StartingPoint() IMainStop {
	return sr.startingPoint
}
