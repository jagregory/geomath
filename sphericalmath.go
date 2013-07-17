package geomath

import . "math"

const r float64 = 6371.0
const radian float64 = Pi / 180.0

// Measures the distance between two points on a sphere using the spherical law of cosines
// http://en.wikipedia.org/wiki/Spherical_law_of_cosines
func CosineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	lond := degreesToRadians(lon2 - lon1)
	lat1r := degreesToRadians(lat1)
	lat2r := degreesToRadians(lat2)

	return Acos(Sin(lat1r)*Sin(lat2r)+Cos(lat1r)*Cos(lat2r)*Cos(lond)) * r
}

// Measures the distance between two points on a sphere using the haversine formula
// http://en.wikipedia.org/wiki/Haversine_formula
func HaversineDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	latd := degreesToRadians(lat2 - lat1)
	lond := degreesToRadians(lon2 - lon1)
	lat1d := degreesToRadians(lat1)
	lat2d := degreesToRadians(lat2)

	a := Sin(latd/2)*Sin(latd/2) + Sin(lond/2)*Sin(lond/2)*Cos(lat1d)*Cos(lat2d)
	c := 2 * Atan2(Sqrt(a), Sqrt(1-a))

	return r * c
}

func degreesToRadians(d float64) float64 {
	return d * radian
}
