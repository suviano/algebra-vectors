package main

import (
	"fmt"
	"math"
)

// IVector vectors common algebra operations interface
type IVector interface {
	Str() string
	Sum(addendVector Vector) Vector
	Minus(vector Vector) Vector
	Multiply(scalar float64) Vector
	Equals(vector Vector) bool
	Magnitude(point1, point2 []float64) float64
	Direction() []float64
	DotProduct(v1, v2 Vector)
}

// Vector is the structure data to represent a algebra vector
type Vector struct {
	Coordinates []float64
}

// Dimensions of a existing vector
func (v *Vector) Dimensions() int {
	return len(v.Coordinates)
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
func (v *Vector) Sum(addendVector Vector) Vector {
	operation := func(augendCoordinate float64, addendCoordinate float64) float64 {
		return augendCoordinate + addendCoordinate
	}
	multiDimensionVectorIterator(v, &addendVector, operation)
	return Vector{Coordinates: v.Coordinates}
}

// Minus vectors algebra subtraction operation
func (v *Vector) Minus(subtrahendVector Vector) Vector {
	operation := func(minuendCoordinate float64, subtrahendCoordinate float64) float64 {
		return minuendCoordinate - subtrahendCoordinate
	}
	multiDimensionVectorIterator(v, &subtrahendVector, operation)
	return Vector{Coordinates: v.Coordinates}
}

// Multiply vector multiply algebra operation
func (v *Vector) Multiply(scalar float64) Vector {
	newCoordinates := []float64{}
	for _, coordinate := range v.Coordinates {
		newCoordinates = append(newCoordinates, coordinate*scalar)
	}
	v.Coordinates = newCoordinates
	return Vector{Coordinates: newCoordinates}
}

// Magnitude distance between vectors
func (v *Vector) Magnitude() float64 {
	coordinatePowSum := 0.0
	for _, coordinate := range v.Coordinates {
		coordinatePowSum += math.Pow(coordinate, 2)
	}
	return math.Sqrt(coordinatePowSum)
}

// Normalization operation to get the unit vector (A unit vector has length equal to 1)
// and still keep the direction to the original vector,
// a magnitude of a unit vector is always equal to 1
func (v *Vector) Normalization() []float64 {
	magnitude := v.Magnitude()
	unitVector := []float64{}
	for _, coordinate := range v.Coordinates {
		unitCoord := (1 / magnitude) * coordinate
		unitVector = append(unitVector, unitCoord)
	}
	return unitVector
}

// Direction canonical representation of a vector direction,
// is the same as getting the unit vector
func (v *Vector) Direction() []float64 {
	return v.Normalization()
}

func normalizeVectors(v1, v2 Vector) int {
	v1Dimension := v1.Dimensions()
	v2Dimension := v2.Dimensions()
	if v1Dimension > v2Dimension {
		v2Diff := int(math.Copysign(float64(v2Dimension-v1Dimension), 1))
		for index := 0; index < v2Diff; index++ {
			v2.Coordinates = append(v2.Coordinates, 0)
		}
	} else if v1Dimension < v2Dimension {
		v1Diff := int(math.Copysign(float64(v1Dimension-v2Dimension), 1))
		for index := 0; index < v1Diff; index++ {
			v1.Coordinates = append(v1.Coordinates, 0)
		}
	}

	return v1.Dimensions()
}

// DotProduct product of two vector multiplication
func DotProduct(v1, v2 Vector) float64 {
	dimensions := normalizeVectors(v1, v2)

	var response float64
	for index := 0; index < dimensions; index++ {
		response += (v1.Coordinates[index] * v2.Coordinates[index])
	}

	return response
}

// AngleBetweenVectors calculates the angle two vectors via dot product
func AngleBetweenVectors(v1, v2 Vector) float64 {
	dotProductValue := DotProduct(v1, v2)
	v1Magnitude := math.Copysign(v1.Magnitude(), 1)
	v2Magnitude := math.Copysign(v2.Magnitude(), 1)
	return math.Acos(dotProductValue / (v1Magnitude * v2Magnitude))
}
