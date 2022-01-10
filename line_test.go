package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLineInitialize(t *testing.T) {
	Convey("TheLineVectorDefiningDefaultsWithoutError", t, func() {
		normalVector := Vector{Coordinates: []float64{6, 6, 6}}
		line := Line{}
		err := line.Initialize(normalVector, 69)
		So(err, ShouldBeNil)
	})

	Convey("TheLineVectorWillNotPanicForEmptyValues", t, func() {
		normalVector := Vector{}
		line := Line{}
		err := line.Initialize(normalVector, 69)
		So(err, ShouldBeNil)
	})
}

func TestLineStr(t *testing.T) {
	Convey("RetrievingTheRightStringRepresentationOfAnLine", t, func() {
		normalVector := Vector{Coordinates: []float64{6, 6, 6}}
		line := Line{}
		_ = line.Initialize(normalVector, 69)
		lineString := line.Str()
		So(len(lineString), ShouldNotEqual, 0)
	})
}

func TestLineIsParallel(t *testing.T) {
	Convey("ParallelLinesIdentified", t, func() {
		line1 := Line{}
		err := line1.Initialize(Vector{Coordinates: []float64{1.182, 5.562}}, 6.744)
		So(err, ShouldBeNil)

		line2 := Line{}
		err = line2.Initialize(Vector{Coordinates: []float64{1.773, 8.343}}, 9.525)
		So(err, ShouldBeNil)

		// "the lines are parallel but was not identified as such"
		So(line1.IsParallel(line2), ShouldBeTrue)
	})

	Convey("UnParallelLinesIdentified", t, func() {
		line1 := Line{}
		err := line1.Initialize(Vector{Coordinates: []float64{7.204, 3.182}}, 86.8)
		// vector initialization for the line1 throwed an error
		So(err, ShouldBeNil)

		line2 := Line{}
		err = line2.Initialize(Vector{Coordinates: []float64{8.172, 4.114}}, 9.883)
		// vector initialization for the line2 throwed an error
		So(err, ShouldBeNil)

		// the lines are not parallel but was not identified as such
		So(line1.IsParallel(line2), ShouldBeFalse)
	})
}

func TestLineIsEqual(t *testing.T) {
	Convey("UnequalLines", t, func() {
		line1 := Line{}
		err := line1.Initialize(Vector{Coordinates: []float64{7.204, 3.182}}, 86.8)
		// "vector initialization for the line1 throwed an error"
		So(err, ShouldBeNil)

		line2 := Line{}
		err = line2.Initialize(Vector{Coordinates: []float64{8.172, 4.114}}, 9.883)
		// "vector initialization for the line2 throwed an error
		So(err, ShouldBeNil)

		// "the lines are not equal, but where identified as such"
		So(line1.IsEqual(line2), ShouldBeFalse)
	})

	Convey("EqualLines", t, func() {
		l1 := Line{}
		err := l1.Initialize(Vector{Coordinates: []float64{4.046, 2.836}}, 1.21)
		// "vector initialization for the line1 throwed an error"
		So(err, ShouldBeNil)

		l2 := Line{}
		err = l2.Initialize(Vector{Coordinates: []float64{10.115, 7.09}}, 3.025)
		// "vector initialization for the line2 throwed an error"
		So(err, ShouldBeNil)

		// "the lines are equal, but they where identified as such"
		So(l1.IsEqual(l2), ShouldBeTrue)
	})
}

func TestLine(t *testing.T) {
	Convey("Intersection", t, func() {
		l1 := Line{}
		err := l1.Initialize(Vector{Coordinates: []float64{7.204, 3.182}}, 8.68)
		// "Vector initialization for the l1 throwed an error"
		So(err, ShouldBeNil)

		l2 := Line{
			Dimensions:   0,
			NormalVector: Vector{},
			ConstantTerm: 0,
			BasePoint:    Vector{},
		}
		err = l2.Initialize(Vector{Coordinates: []float64{8.172, 4.114}}, 9.883)
		// "Vector initialization for the l2 throwed an error"
		So(err, ShouldBeNil)

		intersection, err := l1.IntersectWith(l2)
		// "Should not be throw a error %+v", err
		So(err, ShouldBeNil)

		parsedPointOne := fmt.Sprintf("%.3f", intersection.Coordinates[0])
		// "First point of intersection should be 1.173 returned %s", parsedPointOne
		So(parsedPointOne != "1.173", ShouldBeFalse)

		parsedPointTwo := fmt.Sprintf("%.3f", intersection.Coordinates[1])
		// "First point of intersection should be 0.073 returned %s", parsedPointTwo
		So(parsedPointTwo != "0.073", ShouldBeFalse)
	})
}
