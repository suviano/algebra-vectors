package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetPrecision(t *testing.T) {

	Convey("Expected precision definition", t, func() {
		roundedNum := SetPrecision(12.13399123, 3)
		expectedNum := 12.134
		if roundedNum != expectedNum {
			// SetPrecision method returned %f expected %f
			So(roundedNum, ShouldEqual, expectedNum)
		}
	})
}
