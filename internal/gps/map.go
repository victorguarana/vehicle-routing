package gps

import (
	"os"
	"strconv"
	"strings"
)

type Map struct {
	Warehouses []Point
	Customers  []Point
}

func LoadMap(filename string) Map {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	completeFile := string(content)
	lines := strings.Split(completeFile, "\n")
	return linesToMap(lines)
}

func fileLineToPoint(line string) Point {
	fields := strings.Split(line, ";")
	latitude, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		panic(err)
	}

	longitude, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		panic(err)
	}

	packageSize, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		panic(err)
	}

	return Point{
		Name:        fields[0],
		Latitude:    latitude,
		Longitude:   longitude,
		PackageSize: packageSize,
	}
}

func linesToMap(lines []string) Map {
	var loadedMap Map
	for _, line := range lines {
		newPoint := fileLineToPoint(line)
		if newPoint.PackageSize == 0 {
			loadedMap.Warehouses = append(loadedMap.Warehouses, newPoint)
		} else {
			loadedMap.Customers = append(loadedMap.Customers, newPoint)
		}
	}
	return loadedMap
}
