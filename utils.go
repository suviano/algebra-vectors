package main

import (
	"fmt"
	"strconv"
)

// SetPrecision round number to a specific precision
func SetPrecision(number float64, precision int) float64 {
	strNumber := fmt.Sprintf("%.3f", number)
	resultNumber, _ := strconv.ParseFloat(strNumber, 64)
	return resultNumber
}
