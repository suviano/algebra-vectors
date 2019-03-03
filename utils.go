package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// SetPrecision round number to a specific precision
func SetPrecision(number float64, precision int) float64 {
	strNumber := fmt.Sprintf("%.3f", number)
	resultNumber, _ := strconv.ParseFloat(strNumber, 64)
	return resultNumber
}

func splitFlagCoordinates(coord string) (coordinates []float64, err error) {
	coordinatesString := strings.Split(coord, ",")
	for index, coordinateString := range coordinatesString {
		coordinate, err := strconv.ParseFloat(coordinateString, 64)
		if err != nil {
			coordErrMsg := fmt.Sprintf("invalid flag in %d position", index)
			err = errors.Wrap(err, coordErrMsg)
		} else {
			coordinates = append(coordinates, coordinate)
		}
	}

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("[splitFlagCoordinates] flag '%s' is invalid", coord))
	}

	return
}

// RadToDegree Convert radiant to degrees
func RadToDegree(rad float64) float64 {
	return (rad * 180) / math.Pi
}
