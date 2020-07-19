package main

import (
	"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVectorStr(t *testing.T) {
	Convey("Expecting the right coordinates", t, func() {
		mVector := Vector{
			Coordinates: []float64{12, 33, 666},
		}

		receivedVectorStr := mVector.Str()
		Convey("Received unexpected string message: "+receivedVectorStr, func() {
			So(receivedVectorStr, ShouldEqual, "Vector: [12 33 666]")
		})
	})

	Convey("Empty coordinates are acceptable", t, func() {
		mVector := Vector{
			Coordinates: []float64{},
		}

		receivedVectorStr := mVector.Str()
		Convey("Received empty string for this vector", func() {
			So(receivedVectorStr, ShouldEqual, "Vector: []")
		})
	})
}

func TestVectorEquals(t *testing.T) {
	Convey("Different Vector Sizes With Different Values", t, func() {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{1, 3}}
		So(v1.Equals(v2), ShouldBeFalse)
	})

	Convey("Simply equal vectors", t, func() {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{1, 2}}
		So(v1.Equals(v2), ShouldBeTrue)
	})

	Convey("Equal vectors with different sizes", t, func() {
		Convey("Equal vectors small dimensions difference", func() {
			v1 := Vector{Coordinates: []float64{1, 2, 0}}
			v2 := Vector{Coordinates: []float64{1, 2}}
			So(v1.Equals(v2), ShouldBeTrue)
		})

		Convey("Not equal first vector having more dimensions than the first", func() {
			v1 := Vector{Coordinates: []float64{1, 2, 0, 0}}
			v2 := Vector{Coordinates: []float64{2, 3}}
			So(v1.Equals(v2), ShouldBeFalse)
		})

		Convey("Equal vectors huge dimensions difference", func() {
			v1 := Vector{Coordinates: []float64{1, 2, 0, 0, 0}}
			v2 := Vector{Coordinates: []float64{1, 2}}
			So(v1.Equals(v2), ShouldBeTrue)
		})

		Convey("Equal vectors second vector has more dimensions than the first", func() {
			v1 := Vector{Coordinates: []float64{1, 2}}
			v2 := Vector{Coordinates: []float64{1, 2, 0, 0}}
			So(v1.Equals(v2), ShouldBeTrue)
		})
	})
}

func TestVectorSum(t *testing.T) {
	Convey("Addition of vectors with same dimentions", t, func() {
		v1 := Vector{Coordinates: []float64{8.218, -9.341}}
		v2 := Vector{Coordinates: []float64{-1.129, 2.111}}

		v1.Sum(v2)

		expectedSumVector := Vector{Coordinates: []float64{7.089, -7.229999999999999}}
		So(v1.Equals(expectedSumVector), ShouldBeTrue)
	})

	Convey("Addition of vectors with different sizes", t, func() {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{4, 5, 6}}

		v1.Sum(v2)

		expectedSumVector := Vector{Coordinates: []float64{5, 7, 6}}
		So(v1.Equals(expectedSumVector), ShouldBeTrue)
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
		angleRadFormatted := fmt.Sprintf("%.2f", angleRad)
		if angleRadFormatted != "0.87" {
			t.Errorf("expecting angle to be 0.87 rad found %s", angleRadFormatted)
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

func TestParallelTo(t *testing.T) {
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

func TestDecomposeVector(t *testing.T) {
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
	baseScenario := func(t *testing.T, v1, v2 Vector, expectedValues []string) {
		product := v1.CrossProduct(v2)

		xCoord := product.Coordinates[0]
		if fmt.Sprintf("%.3f", xCoord) != expectedValues[0] {
			t.Errorf("x coordinate should be %.3f got %s", xCoord, expectedValues[0])
		}

		yCoord := product.Coordinates[1]
		if fmt.Sprintf("%.3f", yCoord) != expectedValues[1] {
			t.Errorf("x coordinate should be %.3f got %s", yCoord, expectedValues[1])
		}

		zCoord := product.Coordinates[2]
		if fmt.Sprintf("%.3f", zCoord) != expectedValues[2] {
			t.Errorf("x coordinate should be %.3f got %s", zCoord, expectedValues[2])
		}

		if !product.IsOrthogonalTo(v1) {
			t.Errorf("cross product result should be orthogonal to v1")
		}

		if !product.IsOrthogonalTo(v2) {
			t.Errorf("cross product result should be orthogonal to v2")
		}
	}

	t.Run("Scenario1", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{5, 3, -2}}
		v2 := Vector{Coordinates: []float64{-1, 0, 3}}
		baseScenario(t, v1, v2, []string{"9.000", "-13.000", "3.000"})
	})

	t.Run("Scenario2", func(t *testing.T) {
		v1 := Vector{Coordinates: []float64{8.462, 7.893, -8.187}}
		v2 := Vector{Coordinates: []float64{6.984, -5.975, 4.778}}
		baseScenario(t, v1, v2, []string{"-11.205", "-97.609", "-105.685"})
	})
}

func TestParallelogramArea(t *testing.T) {
	v1 := Vector{Coordinates: []float64{-8.987, -9.838, 5.031}}
	v2 := Vector{Coordinates: []float64{-4.268, -1.861, -8.866}}
	area := v1.ParallelogramArea(v2)

	if fmt.Sprintf("%.3f", area) != "142.122" {
		t.Errorf("Expecting 1.0 found %.3f", area)
	}
}

func TestTriangleArea(t *testing.T) {
	v1 := Vector{Coordinates: []float64{1.5, 9.547, 3.691}}
	v2 := Vector{Coordinates: []float64{-6.007, 0.124, 5.772}}
	area := v1.TriangleArea(v2)

	if fmt.Sprintf("%.3f", area) != "42.565" {
		t.Errorf("Expecting 1.0 found %.3f", area)
	}
}
