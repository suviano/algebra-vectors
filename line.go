package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

// TODO determine if two lines are parallel
// TODO determine if two lines are equal
// TODO determine the intersection of two lines

// NoNonzeroEltsFoundMsg standard message for no zero element available
const NoNonzeroEltsFoundMsg = "no nonzero elements found"

// ILine line common algebra operations interface
type ILine interface {
	FirstNonzeroIndex([]float64) float64
	Initialize(Vector, float64) error
	Str() string
	SetBasePoint()
}

// Line line algebra structure representation
type Line struct {
	Dimensions   int
	NormalVector Vector
	ConstantTerm float64
	BasePoint    Vector
}

// Initialize the line vectors from the normal vector and a constant
func (l *Line) Initialize(normalVector Vector, constantTerm float64) error {
	l.ConstantTerm = constantTerm
	l.NormalVector = normalVector

	l.Dimensions = 2

	if len(l.NormalVector.Coordinates) == 0 {
		allZeros := make([]float64, 2)

		for index := 0; index < l.Dimensions; index++ {
			allZeros = append(allZeros, 0)
		}

		l.NormalVector = Vector{
			Coordinates: allZeros,
		}
	}

	return l.SetBasePoint()
}

// SetBasePoint define the base point from a vector
func (l *Line) SetBasePoint() error {
	basePointCoords := make([]float64, 2)
	for index := 0; index < l.Dimensions; index++ {
		basePointCoords = append(basePointCoords, 0)
	}

	initialIndex, err := l.FirstNonzeroIndex(l.NormalVector.Coordinates)

	if err != nil {
		l.BasePoint = Vector{}
		log.Printf("setting base point returned an error: %+v", err)
	}

	initialCoefficient := l.NormalVector.Coordinates[int(initialIndex)]

	basePointCoords[int(initialIndex)] = l.ConstantTerm / initialCoefficient

	l.BasePoint = Vector{
		Coordinates: basePointCoords,
	}

	return nil
}

// FirstNonzeroIndex returns the first non zero index items from the slice
func (l *Line) FirstNonzeroIndex(vector []float64) (int, error) {
	for index, item := range vector {
		if item != 0 {
			return index, nil
		}
	}
	return 0, fmt.Errorf(NoNonzeroEltsFoundMsg)
}

// Str return a string representation of the line
func (l *Line) Str() string {
	numDecimalPlaces := 3
	writeCoefficient := func(coefficient float64, isInitialTerm bool) string {
		coefficient = SetPrecision(coefficient, numDecimalPlaces)

		output := ""

		if coefficient < 0 {
			output += "-"
		}

		if coefficient > 0 && !isInitialTerm {
			output += "+"
		}

		if !isInitialTerm {
			output += " "
		}

		if math.Abs(coefficient) != 1 {
			output += fmt.Sprintf("%f", math.Abs(coefficient))
		}

		return output
	}

	n := l.NormalVector

	initialIndex, err := l.FirstNonzeroIndex(n.Coordinates)
	if err != nil {
		log.Fatalf("creating the string representation of an line causes the error %+v", err)
	}

	terms := make([]string, l.Dimensions)
	for index := 0; index < l.Dimensions; index++ {
		if SetPrecision(n.Coordinates[index], numDecimalPlaces) > 0 {
			coefficientResult := writeCoefficient(n.Coordinates[index], index == int(initialIndex))
			terms = append(terms, coefficientResult)
		}
	}

	output := strings.Join(terms, " ")

	constant := SetPrecision(l.ConstantTerm, numDecimalPlaces)
	output += fmt.Sprintf(" %f", constant)

	return output
}

// IsParallel verify if two lines are parallel
func (l *Line) IsParallel(line Line) bool {
	// pick two coordinates from the first line other two from the other
	return false
}
