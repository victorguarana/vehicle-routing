package routes

import (
	"log"

	"github.com/victorguarana/vehicle-routing/src/slc"
)

type IMainRoute interface {
	Append(mainStop IMainStop)
	AtIndex(index int) IMainStop
	First() IMainStop
	Iterator() slc.Iterator[IMainStop]
	Last() IMainStop
	RemoveMainStop(index int)
}

type mainRoute struct {
	mainStops []*mainStop
}

func NewMainRoute(iMainStop IMainStop) IMainRoute {
	ms := iMainStop.(*mainStop)
	return &mainRoute{
		mainStops: []*mainStop{ms},
	}
}

func (r *mainRoute) Append(iMainStop IMainStop) {
	ms := iMainStop.(*mainStop)
	r.mainStops = append(r.mainStops, ms)
}

func (r *mainRoute) AtIndex(index int) IMainStop {
	if index < 0 || index >= len(r.mainStops) {
		log.Printf("AtIndex: index (%d) out of range\n", index)
		return nil
	}
	return r.mainStops[index]
}

func (r *mainRoute) First() IMainStop {
	return r.mainStops[0]
}

func (r *mainRoute) Iterator() slc.Iterator[IMainStop] {
	iMainStops := make([]IMainStop, len(r.mainStops))
	for i, ms := range r.mainStops {
		iMainStops[i] = ms
	}

	return slc.NewIterator(iMainStops)
}

func (r *mainRoute) Last() IMainStop {
	return r.mainStops[len(r.mainStops)-1]
}

func (r *mainRoute) RemoveMainStop(index int) {
	if index < 0 || index >= len(r.mainStops) {
		log.Printf("RemoveMainStop: index (%d) out of range\n", index)
		return
	}
	r.mainStops = append(r.mainStops[:index], r.mainStops[index+1:]...)
}
