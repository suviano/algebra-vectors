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
	Convey("VectorsMinus", t, func() {
		v1 := Vector{Coordinates: []float64{7.119, 8.215}}
		v2 := Vector{Coordinates: []float64{-8.223, .878}}

		v1.Minus(v2)

		expectedDifferenceVector := Vector{Coordinates: []float64{15.342, 7.337}}
		So(v1.Equals(expectedDifferenceVector), ShouldBeTrue)
	})

	Convey("VectorsMinusWithDifferentSizes", t, func() {
		v1 := Vector{Coordinates: []float64{1, 2}}
		v2 := Vector{Coordinates: []float64{4, 5, 6}}

		v1.Minus(v2)

		expectedDifferenceVector := Vector{Coordinates: []float64{-3, -3, 6}}
		So(v1.Equals(expectedDifferenceVector), ShouldBeTrue)
	})
}

func TestVectorScalar(t *testing.T) {
	Convey("VectorsMultiplyPositiveNumber", t, func() {
		vector := Vector{Coordinates: []float64{1, 2}}

		vector.Scalar(2)

		expectedProductVector := Vector{Coordinates: []float64{2, 4}}
		So(vector.Equals(expectedProductVector), ShouldBeTrue)
	})

	Convey("VectorsMultiplyNegativeNumber", t, func() {
		vector := Vector{Coordinates: []float64{1, 2}}

		vector.Scalar(-2)

		expectedProductVector := Vector{Coordinates: []float64{-2, -4}}
		So(vector.Equals(expectedProductVector), ShouldBeTrue)
	})

	Convey("VectorsMultiplyDecimalNumber", t, func() {
		vector := Vector{Coordinates: []float64{1.671, -1.012, -.318}}

		vector.Scalar(7.41)

		expectedProductVector := Vector{Coordinates: []float64{12.38211, -7.49892, -2.35638}}
		So(vector.Equals(expectedProductVector), ShouldBeTrue)
	})
}

func TestVectorMagnitude(t *testing.T) {
	Convey(fmt.Sprintf("VectorMagnitudeScenarioOne 'Magnitude' expected %f", math.Abs(7.440283)), t, func() {
		vector := Vector{Coordinates: []float64{-.221, 7.437}}
		magnitude := vector.Magnitude()
		So(fmt.Sprintf("%.3f", magnitude), ShouldEqual, "7.440")
	})

	Convey(fmt.Sprintf("VectorMagnitudeScenarioTwo 'Magnitude' has expected %f", math.Abs(10.884188)), t, func() {
		vector := Vector{Coordinates: []float64{8.813, -1.331, -6.247}}
		magnitude := vector.Magnitude()
		So(fmt.Sprintf("%.3f", magnitude), ShouldEqual, "10.884")
	})
}

func TestVectorNormalization(t *testing.T) {
	Convey("VectorNormalizationScenarioOne", t, func() {
		unitVector := Vector{Coordinates: []float64{5.581, -2.136}}
		unitVectorNormalization := unitVector.Normalization()
		magnitude := unitVectorNormalization.Magnitude()

		So(fmt.Sprintf("%.0f", magnitude), ShouldEqual, "1")
	})

	Convey("VectorNormalizationScenarioTwo", t, func() {
		unitVector := Vector{Coordinates: []float64{1.996, 3.108, -4.554}}
		unitVectorNormalization := unitVector.Normalization()
		magnitude := unitVectorNormalization.Magnitude()

		So(fmt.Sprintf("%.0f", magnitude), ShouldEqual, "1")
	})
}

func TestDotProduct(t *testing.T) {
	Convey("WithSameDimension: expecting 5", t, func() {
		v1 := Vector{Coordinates: []float64{1, 2, -1}}
		v2 := Vector{Coordinates: []float64{3, 1, 0}}
		result := v1.Dot(v2)
		fResult := fmt.Sprintf("%.3f", result)
		So(fResult, ShouldEqual, "5.000")
	})
}

func TestAngleBetweenVectors(t *testing.T) {
	Convey("Scenario1 expecting angle to be 0.87 rad", t, func() {
		v1 := Vector{Coordinates: []float64{1, 2, -1}}
		v2 := Vector{Coordinates: []float64{3, 1, 0}}

		angleRad := v1.AngleWith(v2, false)
		angleRadFormatted := fmt.Sprintf("%.2f", angleRad)
		So(angleRadFormatted, ShouldEqual, "0.87")
	})
}

func TestOrthogonalTo(t *testing.T) {
	Convey("expecting to not be orthogonal", t, func() {
		Convey("Scenario1", func() {
			vector1 := Vector{Coordinates: []float64{-7.579, -7.88}}
			vector2 := Vector{Coordinates: []float64{22.737, 23.64}}
			isOrthogonal := vector1.IsOrthogonalTo(vector2)
			So(isOrthogonal, ShouldBeFalse)
		})

		Convey("Scenario2", func() {
			vector1 := Vector{Coordinates: []float64{-2.029, 9.97, 4.172}}
			vector2 := Vector{Coordinates: []float64{-9.231, -6.639, -7.245}}
			isOrthogonal := vector1.IsOrthogonalTo(vector2)
			So(isOrthogonal, ShouldBeFalse)
		})
	})
	Convey("expecting to be orthogonal", t, func() {
		Convey("Scenario1", func() {
			vector1 := Vector{Coordinates: []float64{-2.328, -7.284, -1.214}}
			vector2 := Vector{Coordinates: []float64{-1.821, 1.072, -2.94}}
			isOrthogonal := vector1.IsOrthogonalTo(vector2)
			So(isOrthogonal, ShouldBeTrue)
		})

		Convey("Scenario2", func() {
			vector1 := Vector{Coordinates: []float64{2.118, 4.827}}
			vector2 := Vector{Coordinates: []float64{0, 0}}
			isOrthogonal := vector1.IsOrthogonalTo(vector2)
			So(isOrthogonal, ShouldBeTrue)
		})
	})
}

func TestParallelTo(t *testing.T) {
	Convey("Scenario1IsParallel", t, func() {
		vector1 := Vector{Coordinates: []float64{-7.579, -7.88}}
		vector2 := Vector{Coordinates: []float64{22.737, 23.64}}
		isParallel := vector1.IsParallelTo(vector2)
		So(isParallel, ShouldBeTrue)
	})
	Convey("Scenario2IsntParallel", t, func() {
		vector1 := Vector{Coordinates: []float64{-2.029, 9.97, 4.172}}
		vector2 := Vector{Coordinates: []float64{-9.231, -6.639, -7.245}}
		isParallel := vector1.IsParallelTo(vector2)
		So(isParallel, ShouldBeFalse)
	})
	Convey("Scenario3IsntParallel", t, func() {
		vector1 := Vector{Coordinates: []float64{-2.328, -7.284, -1.214}}
		vector2 := Vector{Coordinates: []float64{-1.821, 1.072, -2.94}}
		isParallel := vector1.IsParallelTo(vector2)
		So(isParallel, ShouldBeFalse)
	})
	Convey("Scenario4IsParallel", t, func() {
		vector1 := Vector{Coordinates: []float64{2.118, 4.827}}
		vector2 := Vector{Coordinates: []float64{0, 0}}
		isParallel := vector1.IsParallelTo(vector2)
		So(isParallel, ShouldBeTrue)
	})
}

func TestProject(t *testing.T) {
	Convey("CommonScenario1", t, func() {
		v1 := Vector{Coordinates: []float64{3.039, 1.879}}
		v2 := Vector{Coordinates: []float64{.825, 2.036}}

		projection := v1.Project(v2)

		expectedSlice := []float64{1.083, 2.672}
		So(len(projection.Coordinates), ShouldBeGreaterThanOrEqualTo, 0)
		for index, coordinate := range projection.Coordinates {
			rounderCoord := fmt.Sprintf("%.3f", coordinate)
			So(rounderCoord, ShouldEqual, fmt.Sprintf("%.3f", expectedSlice[index]))
		}
	})
}

func TestOrthogonal(t *testing.T) {
	Convey("CommonScenario1", t, func() {
		v1 := Vector{Coordinates: []float64{-9.88, -3.264, -8.159}}
		v2 := Vector{Coordinates: []float64{-2.155, -9.353, -9.473}}

		projection := v1.Orthogonal(v2)
		So(len(projection.Coordinates), ShouldBeGreaterThanOrEqualTo, 0)

		expectedSlice := []float64{-8.350, 3.376, -1.434}
		for index, coordinate := range projection.Coordinates {
			rounderCoord := fmt.Sprintf("%.3f", coordinate)
			So(rounderCoord, ShouldEqual, fmt.Sprintf("%.3f", expectedSlice[index]))
		}
		So(projection.IsOrthogonalTo(v2), ShouldBeTrue)
	})
}

func TestDecomposeVector(t *testing.T) {
	v1 := Vector{Coordinates: []float64{3.009, -6.172, 3.692, -2.51}}
	v2 := Vector{Coordinates: []float64{6.404, -9.144, 2.759, 8.718}}
	// Are these test cases required to sequential ???? don't reme

	Convey("Validating projection of vectors", t, func() {
		projectionParallel := v1.Project(v2)

		expectedSliceParallel := []float64{1.969, -2.811, .848, 2.680}
		So(projectionParallel.Coordinates, ShouldNotBeEmpty)

		for index, coordinate := range projectionParallel.Coordinates {
			rounderCoord := fmt.Sprintf("%.3f", coordinate)
			So(rounderCoord, ShouldEqual, fmt.Sprintf("%.3f", expectedSliceParallel[index]))
		}
	})

	Convey("Validating orthogonal of vectors", t, func() {
		projectionOrthogonal := v1.Orthogonal(v2)

		So(projectionOrthogonal.Coordinates, ShouldNotBeEmpty)

		expectedSliceOrthogonal := []float64{1.04, -3.361, 2.844, -5.19}
		for index, coordinate := range projectionOrthogonal.Coordinates {
			rounderCoord := fmt.Sprintf("%.3f", coordinate)
			So(rounderCoord, ShouldEqual, fmt.Sprintf("%.3f", expectedSliceOrthogonal[index]))
		}

		So(projectionOrthogonal.IsOrthogonalTo(v2), ShouldBeTrue)
	})
}

func TestCrossProduct(t *testing.T) {
	baseScenario := func(t *testing.T, v1, v2 Vector, expectedValues []string) {
		product := v1.CrossProduct(v2)

		xCoord := product.Coordinates[0]
		So(fmt.Sprintf("%.3f", xCoord), ShouldEqual, expectedValues[0])

		yCoord := product.Coordinates[1]
		So(fmt.Sprintf("%.3f", yCoord), ShouldEqual, expectedValues[1])

		zCoord := product.Coordinates[2]
		So(fmt.Sprintf("%.3f", zCoord), ShouldEqual, expectedValues[2])

		So(product.IsOrthogonalTo(v1), ShouldBeTrue)

		So(product.IsOrthogonalTo(v2), ShouldBeTrue)
	}

	Convey("Scenario1", t, func() {
		v1 := Vector{Coordinates: []float64{5, 3, -2}}
		v2 := Vector{Coordinates: []float64{-1, 0, 3}}
		baseScenario(t, v1, v2, []string{"9.000", "-13.000", "3.000"})
	})

	Convey("Scenario2", t, func() {
		v1 := Vector{Coordinates: []float64{8.462, 7.893, -8.187}}
		v2 := Vector{Coordinates: []float64{6.984, -5.975, 4.778}}
		baseScenario(t, v1, v2, []string{"-11.205", "-97.609", "-105.685"})
	})
}

func TestParallelogramArea(t *testing.T) {
	Convey("Expeting 1.0 area", t, func() {
		v1 := Vector{Coordinates: []float64{-8.987, -9.838, 5.031}}
		v2 := Vector{Coordinates: []float64{-4.268, -1.861, -8.866}}
		area := v1.ParallelogramArea(v2)
		So(fmt.Sprintf("%.3f", area), ShouldEqual, "142.122")
	})
}

func TestTriangleArea(t *testing.T) {
	Convey("Expecting 1.0 area", t, func() {
		v1 := Vector{Coordinates: []float64{1.5, 9.547, 3.691}}
		v2 := Vector{Coordinates: []float64{-6.007, 0.124, 5.772}}
		area := v1.TriangleArea(v2)
		So(fmt.Sprintf("%.3f", area), ShouldEqual, "42.565")
	})
}
