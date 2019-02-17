package main

import "fmt"

/*
try:
    if not coordinates:
        raise ValueError
    self.coordinates = tuple(coordinates)
    self.dimension = len(coordinates)

except ValueError:
    raise ValueError('The coordinates must be nonempty')
*/

// IVector vectors common algebra operations interface
type IVector interface {
	Str() string
	Sum(addendVector Vector)
	Minus(vector Vector)
	Multiply(scalar float64)
	Equals(vector Vector) bool
}

// Vector is the structure data to represent a algebra vector
type Vector struct {
	Coordinates []float64
}

// Equals vectors algebra equality comparison
func (v *Vector) Equals(vector Vector) bool {
	// return self.coordinates == v.coordinates
	if len(vector.Coordinates) != len(v.Coordinates) {
		return false
	}
	for index, coordinate := range v.Coordinates {
		if coordinate != vector.Coordinates[index] {
			return false
		}
	}
	return true
}

// Str vector coordinates string format visualization
func (v *Vector) Str() string {
	return fmt.Sprintf("Vector: %v", v.Coordinates)
}

func multiDimensionVectorIterator(v1, v2 *Vector, operation func(float64, float64) float64) {
	secondVectorLen := len(v2.Coordinates)
	newCoordinates := []float64{}
	for index, coordinate := range v1.Coordinates {
		var secondCoordinate float64
		if secondVectorLen-1 >= index {
			secondCoordinate = v2.Coordinates[index]
		}
		newCoordinates = append(newCoordinates, operation(coordinate, secondCoordinate))
	}

	firstVectorLen := len(v1.Coordinates)
	if secondVectorLen > firstVectorLen {
		newCoordinates = append(newCoordinates, v2.Coordinates[firstVectorLen:]...)
	}

	v1.Coordinates = newCoordinates
}

// Sum vector algebra sum operation
func (v *Vector) Sum(addendVector Vector) {
	operation := func(augendCoordinate float64, addendCoordinate float64) float64 {
		return augendCoordinate + addendCoordinate
	}
	multiDimensionVectorIterator(v, &addendVector, operation)
}

// Minus vectors algebra subtraction operation
func (v *Vector) Minus(subtrahendVector Vector) {
	operation := func(minuendCoordinate float64, subtrahendCoordinate float64) float64 {
		return minuendCoordinate - subtrahendCoordinate
	}
	multiDimensionVectorIterator(v, &subtrahendVector, operation)
}

// Multiply vector multiply algebra operation
func (v *Vector) Multiply(scalar float64) {
	newCoordinates := []float64{}
	for _, coordinate := range v.Coordinates {
		newCoordinates = append(newCoordinates, coordinate*scalar)
	}
	v.Coordinates = newCoordinates
}
