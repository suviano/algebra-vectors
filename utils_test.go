package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetPrecision(t *testing.T) {

	Convey("Expected precision definition", t, func() {
		originalNum := 12.13399123
		roundedNum := SetPrecision(originalNum, 3)
		expectedNum := 12.134
		So(originalNum, ShouldNotAlmostEqual, expectedNum)
		So(roundedNum, ShouldEqual, expectedNum)
	})
}
