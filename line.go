package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

// NoNonzeroEltsFoundMsg standard message for no zero element available
const NoNonzeroEltsFoundMsg = "no nonzero elements found"

// ILine line common algebra operations interface
type ILine interface {
	FirstNonzeroIndex([]float64) float64
	Initialize(Vector, float64) error
	Str() string
	SetBasePoint()
	IsParallel(line Line) bool
	IsEqual(line Line) bool
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
		var allZeros []float64

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
	var basePointCoords []float64
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

	var terms []string
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
	return l.NormalVector.IsParallelTo(line.NormalVector)
}

// IsEqual verify if two lines are equal
func (l *Line) IsEqual(line Line) bool {
	if l.NormalVector.IsZero() {
		if !line.NormalVector.IsZero() {
			return false
		} else {
			diff := l.ConstantTerm - line.ConstantTerm
			return diff == 0
		}
	} else if line.NormalVector.IsZero() {
		return false
	}

	if !l.IsParallel(line) {
		return false
	}

	b1 := l.BasePoint
	b2 := line.BasePoint

	basePointDiff := b1.Minus(b2)
	n := l.NormalVector
	return basePointDiff.IsOrthogonalTo(n)
}

func (l *Line) IntersectWith(line Line) (Vector, error) {
	if len(l.NormalVector.Coordinates) != 2 {
		return Vector{}, fmt.Errorf("line1 must be have two dimenstions")
	}
	A := l.NormalVector.Coordinates[0]
	B := l.NormalVector.Coordinates[1]

	if len(line.NormalVector.Coordinates) != 2 {
		return Vector{}, fmt.Errorf("line2 must have two dimenstions")
	}
	C := line.NormalVector.Coordinates[0]
	D := line.NormalVector.Coordinates[1]

	k1 := l.ConstantTerm
	k2 := line.ConstantTerm

	xNumerator := (D * k1) - (B * k2)
	yNumerator := (-C * k1) + (A * k2)
	oneOverDenominator := 1 / ((A * D) - (B * C))

	v := Vector{Coordinates: []float64{xNumerator, yNumerator}}
	return v.Scalar(oneOverDenominator), nil
}
