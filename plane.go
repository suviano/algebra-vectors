package main

import (
	"fmt"
	"log"
)

// IPlane plane common algebra operations interface
type IPlane interface {
	FirstNonzeroIndex([]float64) float64
	Initialize(Vector, float64) error
	Str() string
	SetBasePoint()
	IsParallel(line Plane) bool
	IsEqual(line Plane) bool
}

// Plane plane common algebra operations interface
type Plane struct {
	Dimensions   int
	NormalVector Vector
	ConstantTerm float64
	BasePoint    Vector
})

// Initialize a plane struct
func (p *Plane) Initialize(normalVector Vector, constantTerm float64) error {
	p.ConstantTerm = constantTerm
	p.NormalVector = normalVector

	p.Dimensions = 3

	if len(p.NormalVector.Coordinates) == 0 {
		var allZeros []float64

		for index := 0; index < p.Dimensions; index++ {
			allZeros = append(allZeros, 0)
		}

		p.NormalVector = Vector{
			Coordinates: allZeros,
		}
	}
	return p.SetBasePoint()
}

// SetBasePoint define
func (p *Plane) SetBasePoint() error {
	var basePointCoords []float64
	for index := 0; index < p.Dimensions; index++ {
		basePointCoords = append(basePointCoords, 0)
	}
	initialIndex, err := p.FirstNonZeroIndex(p.NormalVector.Coordinates)
	if err != nil {
		p.BasePoint = Vector{}
		log.Printf("setting base point returned an error: %+v", err)
	}
	initialCoefficient := p.NormalVector.Coordinates[int(initialIndex)]
	basePointCoords[int(initialIndex)] = p.ConstantTerm / initialCoefficient
	p.BasePoint = Vector{
		Coordinates: basePointCoords,
	}
	return nil
}

// FirstNonZeroIndex get the first non zero index
func (p *Plane) FirstNonZeroIndex(vector []float64) (int, error) {
	for index, item := range vector {
		if item != 0 {
			return index, nil
		}
	}
	return 0, fmt.Errorf(NoNonzeroEltsFoundMsg)
}

/*
Check if the following are equal, parallel but unequal, not parallel
-0.412,3.806,0.728=-3.46
1.03,-9.515,-1.82=8.65
--------
2.611,5.528,0.283=4.6
7.715,8.306,5.342=3.76
--------
-7.926,8.625,-7.217=-7.952
-2.642,2.875,-2.404=-2.443
*/
