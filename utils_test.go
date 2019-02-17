package main

import "testing"

func TestSetPrecision(t *testing.T) {
	roundedNum := SetPrecision(12.13399123, 3)
	expectedNum := 12.134
	if roundedNum != expectedNum {
		t.Errorf("SetPrecision method returned %f expected %f", roundedNum, expectedNum)
	}
}
