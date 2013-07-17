package geomath

import (
	. "strconv"
	"testing"
)

const precision = 8

var cosineTests = []struct {
	lat1   float64
	lon1   float64
	lat2   float64
	lon2   float64
	result float64
}{
	{0, 0, 0, 0, 0},
	{-33.884098, 151.206823, -33.865965, 151.205667, 2.019120017019051},
	{-33.7028012445, 151.098494013, -33.8458927727, 151.21182505, 19.04963200715648},
}

func TestCosineDistance(t *testing.T) {
	for ci, tt := range cosineTests {
		result := CosineDistance(tt.lat1, tt.lon1, tt.lat2, tt.lon2)
		roundedResult := FormatFloat(result, 'f', precision, 64)

		if roundedResult != FormatFloat(tt.result, 'f', precision, 64) {
			t.Errorf("Case %d: Expected %e got %s", ci+1, tt.result, roundedResult)
		}
	}
}

var haversineTests = []struct {
	lat1   float64
	lon1   float64
	lat2   float64
	lon2   float64
	result float64
}{
	{0, 0, 0, 0, 0},
	{-33.884098, 151.206823, -33.865965, 151.205667, 2.019120017019051},
	{-33.7028012445, 151.098494013, -33.8458927727, 151.21182505, 19.04963200715648},
}

func TestHaversineDistance(t *testing.T) {
	for ci, tt := range haversineTests {
		result := HaversineDistance(tt.lat1, tt.lon1, tt.lat2, tt.lon2)
		roundedResult := FormatFloat(result, 'f', precision, 64)

		if roundedResult != FormatFloat(tt.result, 'f', precision, 64) {
			t.Errorf("Case %d: Expected %e got %s", ci+1, tt.result, roundedResult)
		}
	}
}
