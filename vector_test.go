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

func TestVectorScalar(t *testing.T) {
	t.Run("VectorsMultiplyPositiveNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1, 2}}

		vector.Scalar(2)

		expectedProductVector := Vector{Coordinates: []float64{2, 4}}
		if !vector.Equals(expectedProductVector) {
			t.Errorf("method 'Scalar' has returned %s while the expected is %s", vector.Str(), expectedProductVector.Str())
		}
	})

	t.Run("VectorsMultiplyNegativeNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1, 2}}

		vector.Scalar(-2)

		expectedProductVector := Vector{Coordinates: []float64{-2, -4}}
		if !vector.Equals(expectedProductVector) {
			t.Errorf("method 'Scalar' has returned %s while the expected is %s", vector.Str(), expectedProductVector.Str())
		}
	})

	t.Run("VectorsMultiplyDecimalNumber", func(t *testing.T) {
		vector := Vector{Coordinates: []float64{1.671, -1.012, -.318}}

		vector.Scalar(7.41)

		expectedProductVector := Vector{Coordinates: []float64{12.38211, -7.49892, -2.35638}}
		if !vector.Equals(expectedProductVector) {
			t.Errorf("method 'Scalar' has returned %s while the expected is %s", vector.Str(), expectedProductVector.Str())
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

func TestVectorNormalization(t *testing.T) {
	unitVectorErrMsg := "Direction should have a vector of 1 has %.0f"
	t.Run("VectorNormalizationScenarioOne", func(t *testing.T) {
		unitVector := Vector{Coordinates: []float64{5.581, -2.136}}
		unitVectorNormalization := unitVector.Normalization()
		magnitude := unitVectorNormalization.Magnitude()
		if fmt.Sprintf("%.0f", magnitude) != "1" {
			t.Errorf(unitVectorErrMsg, magnitude)
		}
	})

	t.Run("VectorNormalizationScenarioTwo", func(t *testing.T) {
		unitVector := Vector{Coordinates: []float64{1.996, 3.108, -4.554}}
		unitVectorNormalization := unitVector.Normalization()
		magnitude := unitVectorNormalization.Magnitude()
		if fmt.Sprintf("%.0f", magnitude) != "1" {
			t.Errorf(unitVectorErrMsg, magnitude)
		}
	})
}

func TestDotProduct(t *testing.T) {
	t.Run("WithSameDimension", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{1, 2, -1}}
		v2 := Vector{Coordinates: []float64{3, 1, 0}}
		result := v1.Dot(v2)
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

		angleRad := v1.AngleWith(v2, false)
		angleRadFormated := fmt.Sprintf("%.2f", angleRad)
		if angleRadFormated != "0.87" {
			t.Errorf("expecting angle to be 0.87 rad found %s", angleRadFormated)
		}
	})
}

func TestOrthogonalTo(t *testing.T) {
	t.Run("Scenario1", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{-7.579, -7.88}}
		vector2 := Vector{Coordinates: []float64{22.737, 23.64}}
		isOrthogonal := vector1.IsOrthogonalTo(vector2)
		if isOrthogonal {
			t.Error("expecting not to be orthogonal")
		}
	})
	t.Run("Scenario2", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{-2.029, 9.97, 4.172}}
		vector2 := Vector{Coordinates: []float64{-9.231, -6.639, -7.245}}
		isOrthogonal := vector1.IsOrthogonalTo(vector2)
		if isOrthogonal {
			t.Error("expecting not to be orthogonal")
		}
	})
	t.Run("Scenario3", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{-2.328, -7.284, -1.214}}
		vector2 := Vector{Coordinates: []float64{-1.821, 1.072, -2.94}}
		isOrthogonal := vector1.IsOrthogonalTo(vector2)
		if !isOrthogonal {
			t.Error("expecting to be orthogonal")
		}
	})
	t.Run("Scenario4", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{2.118, 4.827}}
		vector2 := Vector{Coordinates: []float64{0, 0}}
		isOrthogonal := vector1.IsOrthogonalTo(vector2)
		if !isOrthogonal {
			t.Error("expecting to be orthogonal")
		}
	})
}

func TestParalelTo(t *testing.T) {
	t.Run("Scenario1", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{-7.579, -7.88}}
		vector2 := Vector{Coordinates: []float64{22.737, 23.64}}

		isParallel := vector1.IsParallelTo(vector2)

		if !isParallel {
			t.Error("expecting to be parallel")
		}
	})
	t.Run("Scenario2", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{-2.029, 9.97, 4.172}}
		vector2 := Vector{Coordinates: []float64{-9.231, -6.639, -7.245}}

		isParallel := vector1.IsParallelTo(vector2)

		if isParallel {
			t.Error("expecting not to be parallel")
		}
	})
	t.Run("Scenario3", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{-2.328, -7.284, -1.214}}
		vector2 := Vector{Coordinates: []float64{-1.821, 1.072, -2.94}}

		isParallel := vector1.IsParallelTo(vector2)

		if isParallel {
			t.Error("expecting not to be parallel")
		}
	})
	t.Run("Scenario4", func(t *testing.T) {
		vector1 := Vector{Coordinates: []float64{2.118, 4.827}}
		vector2 := Vector{Coordinates: []float64{0, 0}}

		isParallel := vector1.IsParallelTo(vector2)

		if !isParallel {
			t.Error("expecting to be parallel")
		}
	})
}

func TestProject(t *testing.T) {
	t.Run("CommonScenario1", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{3.039, 1.879}}
		v2 := Vector{Coordinates: []float64{.825, 2.036}}

		projection := v1.Project(v2)

		expectedSlice := []float64{1.083, 2.672}
		if len(projection.Coordinates) <= 0 {
			t.Errorf("empty projection coordinates")
		}

		for index, coordinate := range projection.Coordinates {
			rounderCoord := fmt.Sprintf("%.3f", coordinate)
			if rounderCoord != fmt.Sprintf("%.3f", expectedSlice[index]) {
				t.Errorf("expecting %+v to be equal %+v", projection.Coordinates, expectedSlice)
				break
			}
		}
	})
}

func TestOrthogonal(t *testing.T) {
	t.Run("CommonScenario1", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{-9.88, -3.264, -8.159}}
		v2 := Vector{Coordinates: []float64{-2.155, -9.353, -9.473}}

		projection := v1.Orthogonal(v2)

		if len(projection.Coordinates) <= 0 {
			t.Errorf("empty projection coordinates")
		}

		expectedSlice := []float64{-8.350, 3.376, -1.434}
		for index, coordinate := range projection.Coordinates {
			rounderCoord := fmt.Sprintf("%.3f", coordinate)
			if rounderCoord != fmt.Sprintf("%.3f", expectedSlice[index]) {
				t.Errorf("expecting %+v to be equal %+v", projection.Coordinates, expectedSlice)
				break
			}
		}

		if !projection.IsOrthogonalTo(v2) {
			t.Errorf("expecint %+v to be orthogonal to %+v", projection, v2)
		}
	})
}

func TestDecomponseVector(t *testing.T) {
	// scenario3 [] [1.040,-3.361,-5.190]
	v1 := Vector{Coordinates: []float64{3.009, -6.172, 3.692, -2.51}}
	v2 := Vector{Coordinates: []float64{6.404, -9.144, 2.759, 8.718}}

	projectionParallel := v1.Project(v2)

	expectedSliceParallel := []float64{1.969, -2.811, .848, 2.680}
	if len(projectionParallel.Coordinates) <= 0 {
		t.Errorf("empty projection coordinates")
	}

	for index, coordinate := range projectionParallel.Coordinates {
		rounderCoord := fmt.Sprintf("%.3f", coordinate)
		if rounderCoord != fmt.Sprintf("%.3f", expectedSliceParallel[index]) {
			t.Errorf("expecting %+v to be equal %+v", projectionParallel.Coordinates, expectedSliceParallel)
			break
		}
	}

	projectionOrthogonal := v1.Orthogonal(v2)

	if len(projectionOrthogonal.Coordinates) <= 0 {
		t.Errorf("empty projection coordinates")
	}

	expectedSliceOrthogonal := []float64{1.04, -3.361, 2.844, -5.19}
	for index, coordinate := range projectionOrthogonal.Coordinates {
		rounderCoord := fmt.Sprintf("%.3f", coordinate)
		if rounderCoord != fmt.Sprintf("%.3f", expectedSliceOrthogonal[index]) {
			t.Errorf("expecting %+v to be equal %+v", projectionOrthogonal.Coordinates, expectedSliceOrthogonal)
			break
		}
	}

	if !projectionOrthogonal.IsOrthogonalTo(v2) {
		t.Errorf("expecint %+v to be orthogonal to %+v", projectionOrthogonal, v2)
	}
}

func TestCrossProduct(t *testing.T) {
	t.Run("Scenario1", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{5, 3, -2}}
		v2 := Vector{Coordinates: []float64{-1, 0, 3}}

		product := v1.CrossProduct(v2)

		xCoord := product.Coordinates[0]
		if fmt.Sprintf("%.3f", xCoord) != "9.000" {
			t.Errorf("x coordinate should be %.3f got 9.000", xCoord)
		}

		yCoord := product.Coordinates[1]
		if fmt.Sprintf("%.3f", yCoord) != "-13.000" {
			t.Errorf("x coordinate should be %.3f got -13.000", yCoord)
		}

		zCoord := product.Coordinates[2]
		if fmt.Sprintf("%.3f", zCoord) != "3.000" {
			t.Errorf("x coordinate should be %.3f got 3.000", zCoord)
		}

		if !product.IsOrthogonalTo(v1) {
			t.Errorf("cross product result should be orthogonal to v1")
		}

		if !product.IsOrthogonalTo(v2) {
			t.Errorf("cross product result should be orthogonal to v2")
		}
	})
}

func TestParallelogramArea(t *testing.T) {

}
func TestTriangleArea(t *testing.T) {

}
