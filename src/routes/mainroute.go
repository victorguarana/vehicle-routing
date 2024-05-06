package routes

import (
	"errors"
	"fmt"
	"strings"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type IMainRoute interface {
	Append(mainStop IMainStop)
	AtIndex(index int) (IMainStop, error)
	First() IMainStop
	Last() IMainStop
	Len() int
	RemoveMainStop(int) error
	String() string
}

type mainRoute struct {
	mainStops []*mainStop
}

func (r *mainRoute) Append(iMainStop IMainStop) {
	ms := iMainStop.(*mainStop)
	r.mainStops = append(r.mainStops, ms)
}

func (r *mainRoute) AtIndex(index int) (IMainStop, error) {
	if index < 0 || index >= len(r.mainStops) {
		return nil, ErrIndexOutOfRange
	}
	return r.mainStops[index], nil
}

func (r *mainRoute) First() IMainStop {
	return r.mainStops[0]
}

func (r *mainRoute) Last() IMainStop {
	return r.mainStops[len(r.mainStops)-1]
}

func (r *mainRoute) Len() int {
	return len(r.mainStops)
}

func (r *mainRoute) RemoveMainStop(index int) error {
	if index < 0 || index >= len(r.mainStops) {
		return ErrIndexOutOfRange
	}
	r.mainStops = append(r.mainStops[:index], r.mainStops[index+1:]...)
	return nil
}

func (r *mainRoute) String() string {
	str := []string{"Route:"}
	for i, mainStop := range r.mainStops {
		str = append(str, fmt.Sprintf("  MainStop #%d (%s)", i, mainStop.point))
		for j, flight := range mainStop.subRoutes {
			if flight.startingPoint != mainStop {
				continue
			}
			str = append(str, fmt.Sprintf("    Flight #%d.%d:", i, j))
			str = append(str, fmt.Sprintf("     Takeoff #%d.%d (%s)", i, j, flight.startingPoint.point))
			for k, droneStop := range flight.stops {
				str = append(str, fmt.Sprintf("      DroneStop #%d.%d.%d (%s)", i, j, k, droneStop.point))
			}
			str = append(str, fmt.Sprintf("     Landing #%d.%d (%s)", i, j, flight.returningPoint.point))
		}
		str = append(str, "")
	}
	return strings.Join(str, "\n")
}
