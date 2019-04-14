package main

import (
	"fmt"
	"testing"
)

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

	t.Run("IsParallel", func(t *testing.T) {
		t.Run("ParallelLinesIdentified", func(t *testing.T) {
			line1 := Line{}
			err := line1.Initialize(Vector{Coordinates: []float64{1.182, 5.562}}, 6.744)
			if err != nil {
				t.Errorf("vector initialization for the line1 throwed an error")
			}

			line2 := Line{}
			err = line2.Initialize(Vector{Coordinates: []float64{1.773, 8.343}}, 9.525)
			if err != nil {
				t.Errorf("vector initialization for the line2 throwed an error")
			}

			if !line1.IsParallel(line2) {
				t.Error("the lines are parallel but was not identified as such")
			}

		})

		t.Run("UnParallelLinesIdentified", func(t *testing.T) {
			line1 := Line{}
			err := line1.Initialize(Vector{Coordinates: []float64{7.204, 3.182}}, 86.8)
			if err != nil {
				t.Errorf("vector initialization for the line1 throwed an error")
			}

			line2 := Line{}
			err = line2.Initialize(Vector{Coordinates: []float64{8.172, 4.114}}, 9.883)
			if err != nil {
				t.Errorf("vector initialization for the line2 throwed an error")
			}

			if line1.IsParallel(line2) {
				t.Error("the lines are not parallel but was not identified as such")
			}
		})
	})

	t.Run("IsEqual", func(t *testing.T) {
		t.Run("UnequalLines", func(t *testing.T) {
			line1 := Line{}
			err := line1.Initialize(Vector{Coordinates: []float64{7.204, 3.182}}, 86.8)
			if err != nil {
				t.Errorf("vector initialization for the line1 throwed an error")
			}

			line2 := Line{}
			err = line2.Initialize(Vector{Coordinates: []float64{8.172, 4.114}}, 9.883)
			if err != nil {
				t.Errorf("vector initialization for the line2 throwed an error")
			}

			if line1.IsEqual(line2) {
				t.Error("the lines are not equal, but where identified as such")
			}
		})

		t.Run("EqualLines", func(t *testing.T) {
			l1 := Line{}
			err := l1.Initialize(Vector{Coordinates: []float64{4.046, 2.836}}, 1.21)
			if err != nil {
				t.Errorf("vector initialization for the line1 throwed an error")
			}

			l2 := Line{}
			err = l2.Initialize(Vector{Coordinates: []float64{10.115, 7.09}}, 3.025)
			if err != nil {
				t.Errorf("vector initialization for the line2 throwed an error")
			}

			if !l1.IsEqual(l2) {
				t.Error("the lines are equal, but they where identified as such")
			}
		})
	})

	t.Run("Intersection", func(t *testing.T) {
		l1 := Line{}
		err := l1.Initialize(Vector{Coordinates: []float64{7.204, 3.182}}, 8.68)
		if err != nil {
			t.Errorf("vector initialization for the l1 throwed an error")
		}

		l2 := Line{}
		err = l2.Initialize(Vector{Coordinates: []float64{8.172, 4.114}}, 9.883)
		if err != nil {
			t.Errorf("vector initialization for the l2 throwed an error")
		}

		intersection, err := l1.IntersectWith(l2)

		if err != nil {
			t.Errorf("Should not be throw a error %+v", err)
		}

		parsedPointOne := fmt.Sprintf("%.3f", intersection.Coordinates[0])
		if parsedPointOne != "1.173" {
			t.Errorf("first point of intersection should be 1.173 returned %s", parsedPointOne)
		}

		parsedPointTwo := fmt.Sprintf("%.3f", intersection.Coordinates[1])
		if parsedPointTwo != "0.073" {
			t.Errorf("first point of intersection should be 0.073 returned %s", parsedPointTwo)
		}
	})
}
