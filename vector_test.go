package main

import "testing"

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

		sumVector := v1.Sum(v2)

		expectedSumVector := Vector{Coordinates: []float64{7.089, -7.229999999999999}}
		if !sumVector.Equals(expectedSumVector) {
			t.Errorf("method 'Sum' has returned %s while the expected is %s", sumVector.Str(), expectedSumVector.Str())
		}
	})

	t.Run("VectorsSumWithDifferentSizes", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{4, 5, 6}}

		sumVector := v1.Sum(v2)

		expectedSumVector := Vector{Coordinates: []float64{5, 7, 6}}
		if !sumVector.Equals(expectedSumVector) {
			t.Errorf("method 'Sum' has returned %s while the expected is %s", sumVector.Str(), expectedSumVector.Str())
		}
	})
}

func TestVectorMinus(t *testing.T) {
	t.Run("VectorsMinus", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{7.119, 8.215}}
		v2 := Vector{Coordinates: []float64{-8.223, .878}}

		differenceVector := v1.Minus(v2)

		expectedDifferenceVector := Vector{Coordinates: []float64{-1.104, 7.337}}
		if !differenceVector.Equals(expectedDifferenceVector) {
			t.Errorf("method 'Minus' has returned %s while the expected is %s", differenceVector.Str(), expectedDifferenceVector.Str())
		}
	})

	t.Run("VectorsMinusWithDifferentSizes", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{4, 5, 6}}

		differenceVector := v1.Minus(v2)

		expectedDifferenceVector := Vector{Coordinates: []float64{-3, -3, -6}}
		if !differenceVector.Equals(expectedDifferenceVector) {
			t.Errorf("method 'Minus' has returned %s while the expected is %s", differenceVector.Str(), expectedDifferenceVector.Str())
		}
	})
}

func TestVectorMultiply(t *testing.T) {
	t.Run("VectorsMultiplyPositiveNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1, 2}}

		productVector := vector.Multiply(2)

		expectedProductVector := Vector{Coordinates: []float64{2, 4}}
		if !productVector.Equals(expectedProductVector) {
			t.Errorf("method 'Multiply' has returned %s while the expected is %s", productVector.Str(), expectedProductVector.Str())
		}
	})

	t.Run("VectorsMultiplyNegativeNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1, 2}}

		productVector := vector.Multiply(-2)

		expectedProductVector := Vector{Coordinates: []float64{-2, -4}}
		if !productVector.Equals(expectedProductVector) {
			t.Errorf("method 'Multiply' has returned %s while the expected is %s", productVector.Str(), expectedProductVector.Str())
		}
	})

	t.Run("VectorsMultiplyDecimalNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1.671, -1.012, -0.318}}

		productVector := vector.Multiply(7.41)

		expectedProductVector := Vector{Coordinates: []float64{12.38211, -7.49892, -2.35638}}
		if !productVector.Equals(expectedProductVector) {
			t.Errorf("method 'Multiply' has returned %s while the expected is %s", productVector.Str(), expectedProductVector.Str())
		}
	})
}
