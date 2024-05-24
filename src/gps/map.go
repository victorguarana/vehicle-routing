package gps

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	Warehouses []Point
	Clients    []Point
}

func LoadMap(filename string) Map {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	completeFile := string(content)
	lines := strings.Split(completeFile, "\n")
	return linesToMap(lines)
}

func fileLineToPoint(line string) Point {
	fields := strings.Split(line, ";")
	latitude, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	longitude, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	packageSize, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		log.Fatal(err)
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
			loadedMap.Clients = append(loadedMap.Clients, newPoint)
		}
	}
	return loadedMap
}
