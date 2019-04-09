package main

import "testing"

func TestLine(t *testing.T) {
	t.Run("Initialize", func(t *testing.T) {
		t.Run("TheLineVectorDefiningDefaultsWithoutError", func(t *testing.T) {
			normalVector := Vector{Coordinates: []float64{6, 6, 6}}
			line := Line{}
			err := line.Initialize(normalVector, 69)
			if err != nil {
				t.Errorf("Unnexpected error: %+v initializing the line vector", err)
			}
		})

		t.Run("TheLineVectorWillNotPanicForEmptyValues", func(t *testing.T) {
			normalVector := Vector{}
			line := Line{}
			err := line.Initialize(normalVector, 69)
			if err != nil {
				t.Errorf("Unnexpected error: %+v initializing the line vector with not values", err)
			}
		})
	})

	t.Run("Str", func(t *testing.T) {
		t.Run("RetrievingTheRightStringRepresentationOfAnLine", func(t *testing.T) {
			normalVector := Vector{Coordinates: []float64{6, 6, 6}}
			line := Line{}
			_ = line.Initialize(normalVector, 69)
			lineString := line.Str()
			if len(lineString) == 0 {
				t.Error("failure to parse line to string message")
			}
		})
	})
}
