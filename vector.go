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
	Sum(vector Vector) Vector
	Minus(vector Vector) Vector
	Multiply(scalar float64) Vector
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

// Sum vector algebra sum operation
func (v *Vector) Sum(vector Vector) Vector {
	addendLen := len(vector.Coordinates)
	newCoordinates := []float64{}
	for index, coordinate := range v.Coordinates {
		var addition float64
		if addendLen-1 >= index {
			addition = vector.Coordinates[index]
		}
		// sumResult := SetPrecision(coordinate+addition, 3)
		newCoordinates = append(newCoordinates, coordinate+addition)
	}

	augendLen := len(v.Coordinates)
	if addendLen > augendLen {
		newCoordinates = append(newCoordinates, vector.Coordinates[augendLen:]...)
	}

	v.Coordinates = newCoordinates
	return Vector{Coordinates: newCoordinates}
}

// Minus vectors algebra subtraction operation
func (v *Vector) Minus(vector Vector) Vector {
	return Vector{}
}

// Multiply vector multiply algebra operation
func (v *Vector) Multiply(scalar float64) Vector {
	return Vector{}
}
