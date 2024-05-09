package routes

type ISubRoute interface {
	Append(ISubStop)
	Return(IMainStop)
	ReturningPoint() IMainStop
	StartingPoint() IMainStop
}

type subRoute struct {
	returningPoint *mainStop
	startingPoint  *mainStop
	stops          []*subStop
}

func NewSubRoute(iStartingPoint, iReturningPoint IMainStop) ISubRoute {
	startingPoint := iStartingPoint.(*mainStop)
	returningPoint := iReturningPoint.(*mainStop)
	sr := &subRoute{
		startingPoint:  startingPoint,
		returningPoint: returningPoint,
		stops:          []*subStop{},
	}
	startingPoint.subRoutes = append(startingPoint.subRoutes, sr)
	returningPoint.subRoutes = append(returningPoint.subRoutes, sr)
	return sr
}

func (sr *subRoute) Append(iSubStop ISubStop) {
	sr.stops = append(sr.stops, iSubStop.(*subStop))
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
