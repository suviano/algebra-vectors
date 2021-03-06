package main

import (
	"fmt"
	"math"
)

// IVector vectors common algebra operations interface
type IVector interface {
	Str() string
	Equals(vector Vector) bool
	Sum(addendVector Vector) Vector
	Minus(vector Vector) Vector
	Scalar(scalar float64) Vector
	Magnitude(point1, point2 []float64) float64
	Normalization() []float64
	Dot(v1, v2 Vector)
	AngleWith(vector Vector, inDegrees bool) float64
	IsParallelTo(vector Vector) bool
	IsOrthogonalTo(vector Vector) bool
	Project(vector Vector) Vector
	Orthogonal(vector Vector) Vector
	CrossProduct(vector Vector) Vector
	AreaParallelogram(vector Vector)
	IsZero() bool
}

// Vector is the structure data to represent a algebra vector
type Vector struct {
	Coordinates []float64
}

// Str vector coordinates string format visualization
func (v *Vector) Str() string {
	return fmt.Sprintf("Vector: %v", v.Coordinates)
}

// Dimensions of a existing vector
func (v *Vector) Dimensions() int {
	return len(v.Coordinates)
}

// CoordinateLengthNorm increase a number of dimensions of an vector
func (v *Vector) CoordinateLengthNorm(lengthDifference float64) {
	diff := int(math.Abs(lengthDifference))
	for index := 0; index < diff; index++ {
		v.Coordinates = append(v.Coordinates, 0)
	}
}

// Equals vectors algebra equality comparison
func (v *Vector) Equals(vector Vector) bool {
	// return self.coordinates == v.coordinates
	if len(vector.Coordinates) != len(v.Coordinates) {
		difference := float64(len(vector.Coordinates) - len(v.Coordinates))
		if difference < 0 {
			vector.CoordinateLengthNorm(difference)
		}

		if difference > 0 {
			v.CoordinateLengthNorm(difference)
		}
	}

	for index, coordinate := range v.Coordinates {
		if coordinate != vector.Coordinates[index] {
			return false
		}
	}

	return true
}

func multiDimensionVectorIterator(v1, v2 *Vector, operation func(float64, float64) float64) {
	secondVectorLen := len(v2.Coordinates)
	var newCoordinates []float64
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

// Sum vector addition operation
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

// Scalar vector multiply algebra operation
func (v *Vector) Scalar(scalar float64) Vector {
	var newCoordinates []float64
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
// returns the "Direction"
func (v *Vector) Normalization() Vector {
	var unitVector []float64
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector{Coordinates: unitVector}
	}

	return Vector{Coordinates: v.Scalar(1. / magnitude).Coordinates}
}

func (v *Vector) adjustDimensions(vector Vector) int {
	v1Dimension := v.Dimensions()
	v2Dimension := vector.Dimensions()
	if v1Dimension > v2Dimension {
		v2Diff := int(math.Copysign(float64(v2Dimension-v1Dimension), 1))
		for index := 0; index < v2Diff; index++ {
			vector.Coordinates = append(vector.Coordinates, 0)
		}
	} else if v1Dimension < v2Dimension {
		v1Diff := int(math.Copysign(float64(v1Dimension-v2Dimension), 1))
		for index := 0; index < v1Diff; index++ {
			v.Coordinates = append(v.Coordinates, 0)
		}
	}

	return v.Dimensions()
}

// Dot product of two vector multiplication
func (v *Vector) Dot(vector Vector) float64 {
	dimensions := v.adjustDimensions(vector)

	var response float64
	for index := 0; index < dimensions; index++ {
		response += v.Coordinates[index] * vector.Coordinates[index]
	}

	return response
}

// AngleWith calculates the angle two vectors via dot product
func (v *Vector) AngleWith(vector Vector, inDegrees bool) float64 {
	dotProductValue := v.Dot(vector)
	v1Magnitude := math.Copysign(v.Magnitude(), 1)
	v2Magnitude := math.Copysign(vector.Magnitude(), 1)
	magnitudeProduct := v1Magnitude * v2Magnitude
	angleRad := math.Acos(dotProductValue / magnitudeProduct)
	if inDegrees {
		return RadToDegree(angleRad)
	}
	return angleRad
}

// IsOrthogonalTo verify if a vector is orthogonal to another
func (v *Vector) IsOrthogonalTo(vector Vector) bool {
	dotProduct := math.Copysign(v.Dot(vector), 1)
	tolerance := 1e-10
	return dotProduct < tolerance
}

// IsParallelTo verify if a vector is parallel to another
func (v *Vector) IsParallelTo(vector Vector) bool {
	tolerance := 1e-10
	if v.Magnitude() < tolerance || vector.Magnitude() < tolerance {
		return true
	}
	angleBetweenVectors := v.AngleWith(vector, false)
	// Nan is a valid return for an parallel comparison???
	return angleBetweenVectors == 0 || angleBetweenVectors == math.Pi || math.IsNaN(angleBetweenVectors)
}

func (v *Vector) componentParallelTo(vector Vector) Vector {
	v.adjustDimensions(vector)

	vectorNormalization := vector.Normalization()
	weight := v.Dot(vectorNormalization)
	return vectorNormalization.Scalar(weight)
}

// Project execute the projection of an vector to another
func (v *Vector) Project(vector Vector) Vector {
	// v.adjustDimensions(vector)
	// product := v.Dot(vector)
	// magnitude := product / vector.Magnitude()
	// normalization := vector.Normalization()
	// return normalization.Scalar(magnitude)
	return v.componentParallelTo(vector)
}

// componentOrthogonalTo
func (v *Vector) componentOrthogonalTo(vector Vector) Vector {
	projection := v.componentParallelTo(vector)
	return v.Minus(projection)
}

// Orthogonal return a vector orthogonal to vector argument in relation from the other vector
func (v *Vector) Orthogonal(vector Vector) Vector {
	return v.componentOrthogonalTo(vector)
}

// CrossProduct operation to calculate a vector orthogonal to two vectors
func (v *Vector) CrossProduct(vector Vector) Vector {
	if v.Dimensions() == 3 && v.Dimensions() == 3 {
		return Vector{Coordinates: []float64{
			(v.Coordinates[1] * vector.Coordinates[2]) - (vector.Coordinates[1] * v.Coordinates[2]),
			-1 * ((v.Coordinates[0] * vector.Coordinates[2]) - (vector.Coordinates[0] * v.Coordinates[2])),
			(v.Coordinates[0] * vector.Coordinates[1]) - (vector.Coordinates[0] * v.Coordinates[1])}}
	}
	return Vector{}
}

// ParallelogramArea the area from two vectors parallelogram
func (v *Vector) ParallelogramArea(vector Vector) float64 {
	product := v.CrossProduct(vector)
	return product.Magnitude()
}

// TriangleArea half size of the area from two vectors parallelogram
func (v *Vector) TriangleArea(vector Vector) float64 {
	product := v.CrossProduct(vector)
	return product.Magnitude() / 2
}

// IsZero assert if a vectors is the point zero
func (v *Vector) IsZero() bool {
	coordinatesSum := 0.0
	for _, coordinate := range v.Coordinates {
		coordinatesSum += coordinate
	}
	return coordinatesSum == 0
}
