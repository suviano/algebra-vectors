package main

import (
	"fmt"
	"math"
	"testing"
)

func TestVectorStr(t *testing.T) {
	t.Run("ExpectingTheRightCoordinates", func(t *testing.T) {
		mVector := Vector{
			Coordinates: []float64{12, 33, 666},
		}

		receivedVectorStr := mVector.Str()
		if receivedVectorStr != "Vector: [12 33 666]" {
			t.Errorf("Received unexpected string message: %s", receivedVectorStr)
		}
	})
}

func TestVectorEquals(t *testing.T) {
	t.Run("DifferentVectorSizesWithDifferentValues", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{1, 2, 3}}
		isEqual := v1.Equals(v2)
		if isEqual {
			t.Errorf("method 'Equals' has return true, but %s is different from %s", v1.Str(), v2.Str())
		}
	})
	t.Run("DifferentVectorWithDifferentValues", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{1, 3}}
		isEqual := v1.Equals(v2)
		if isEqual {
			t.Errorf("method 'Equals' has return true, but %s is different from %s", v1.Str(), v2.Str())
		}
	})

	t.Run("EqualVector", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{1, 2}}
		isEqual := v1.Equals(v2)
		if !isEqual {
			t.Errorf("method 'Equals' has return false, but %s is represent the same vector as %s", v1.Str(), v2.Str())
		}
	})
}

func TestVectorSum(t *testing.T) {
	t.Run("VectorsSum", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{8.218, -9.341}}
		v2 := Vector{Coordinates: []float64{-1.129, 2.111}}

		v1.Sum(v2)

		expectedSumVector := Vector{Coordinates: []float64{7.089, -7.229999999999999}}
		if !v1.Equals(expectedSumVector) {
			t.Errorf("method 'Sum' has returned %s while the expected is %s", v1.Str(), expectedSumVector.Str())
		}
	})

	t.Run("VectorsSumWithDifferentSizes", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{4, 5, 6}}

		v1.Sum(v2)

		expectedSumVector := Vector{Coordinates: []float64{5, 7, 6}}
		if !v1.Equals(expectedSumVector) {
			t.Errorf("method 'Sum' has returned %s while the expected is %s", v1.Str(), expectedSumVector.Str())
		}
	})
}

func TestVectorMinus(t *testing.T) {
	t.Run("VectorsMinus", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{7.119, 8.215}}
		v2 := Vector{Coordinates: []float64{-8.223, .878}}

		v1.Minus(v2)

		expectedDifferenceVector := Vector{Coordinates: []float64{15.342, 7.337}}
		if !v1.Equals(expectedDifferenceVector) {
			t.Errorf("method 'Minus' has returned %s while the expected is %s", v1.Str(), expectedDifferenceVector.Str())
		}
	})

	t.Run("VectorsMinusWithDifferentSizes", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{4, 5, 6}}

		v1.Minus(v2)

		expectedDifferenceVector := Vector{Coordinates: []float64{-3, -3, 6}}
		if !v1.Equals(expectedDifferenceVector) {
			t.Errorf("method 'Minus' has returned %s while the expected is %s", v1.Str(), expectedDifferenceVector.Str())
		}
	})
}

func TestVectorMultiply(t *testing.T) {
	t.Run("VectorsMultiplyPositiveNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1, 2}}

		vector.Multiply(2)

		expectedProductVector := Vector{Coordinates: []float64{2, 4}}
		if !vector.Equals(expectedProductVector) {
			t.Errorf("method 'Multiply' has returned %s while the expected is %s", vector.Str(), expectedProductVector.Str())
		}
	})

	t.Run("VectorsMultiplyNegativeNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1, 2}}

		vector.Multiply(-2)

		expectedProductVector := Vector{Coordinates: []float64{-2, -4}}
		if !vector.Equals(expectedProductVector) {
			t.Errorf("method 'Multiply' has returned %s while the expected is %s", vector.Str(), expectedProductVector.Str())
		}
	})

	t.Run("VectorsMultiplyDecimalNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1.671, -1.012, -.318}}

		vector.Multiply(7.41)

		expectedProductVector := Vector{Coordinates: []float64{12.38211, -7.49892, -2.35638}}
		if !vector.Equals(expectedProductVector) {
			t.Errorf("method 'Multiply' has returned %s while the expected is %s", vector.Str(), expectedProductVector.Str())
		}
	})
}

func TestVectorMagnitude(t *testing.T) {
	t.Run("VectorMagnitudeScenarioOne", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{-.221, 7.437}}
		magnitude := vector.Magnitude()
		if "7.440" != fmt.Sprintf("%.3f", magnitude) {
			t.Errorf("method 'Magnitude' has returned %f expecting %f", math.Abs(magnitude), math.Abs(7.440283))
		}
	})

	t.Run("VectorMagnitudeScenarioTwo", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{8.813, -1.331, -6.247}}
		magnitude := vector.Magnitude()
		if "10.884" != fmt.Sprintf("%.3f", magnitude) {
			t.Errorf("method 'Magnitude' has returned %f expected %f", math.Abs(magnitude), math.Abs(10.884188))
		}
	})
}

func TestVectorDirection(t *testing.T) {
	unitVectorErrMsg := "Direction should have a vector of 1 has %.0f"
	t.Run("VectorDirectionScenarioOne", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{5.581, -2.136}}
		unitVector := Vector{Coordinates: vector.Direction()}
		unitVectorMagnitude := unitVector.Magnitude()
		if fmt.Sprintf("%.0f", unitVectorMagnitude) != "1" {
			t.Errorf(unitVectorErrMsg, unitVectorMagnitude)
		}
	})

	t.Run("VectorDirectionScenarioTwo", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1.996, 3.108, -4.554}}
		unitVector := Vector{Coordinates: vector.Direction()}
		unitVectorMagnitude := unitVector.Magnitude()
		if fmt.Sprintf("%.0f", unitVectorMagnitude) != "1" {
			t.Errorf(unitVectorErrMsg, unitVectorMagnitude)
		}
	})
}

func TestDotProduct(t *testing.T) {
	t.Run("WithSameDimension", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2, -1}}
		v2 := Vector{Coordinates: []float64{3, 1, 0}}
		result := DotProduct(v1, v2)
		fResult := fmt.Sprintf("%.3f", result)
		if fResult != "5.000" {
			t.Errorf("expecting 5 got %s", fResult)
		}
	})
}

func TestAngleBetweenVectors(t *testing.T) {
	t.Run("Scenario1", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2, -1}}
		v2 := Vector{Coordinates: []float64{3, 1, 0}}

		angleRad := AngleBetweenVectors(v1, v2)
		angleRadFormated := fmt.Sprintf("%.2f", angleRad)
		if angleRadFormated != "0.87" {
			t.Errorf("expecting angle to be 0.87 rad found %s", angleRadFormated)
		}
	})
}
