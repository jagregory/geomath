# geomath

    go get github.com/jagregory/geomath.git

Provides two functions for calcuating distances on a spheres (aka the earth). `HaversineDistance` calculates the distance using the [Haversine formula](http://en.wikipedia.org/wiki/Haversine_formula) and `CosineDistance` uses the [Spherical Law of Cosines](http://en.wikipedia.org/wiki/Spherical_law_of_cosines). Both can be used interchangably except for their accuracy (hint: Haversine is more accurate). Both implementations make an assumption that the radius of the earth is the magical constant `6371`.

    HaversineDistance(-33.7028012445, 151.098494013, -33.8458927727, 151.21182505) => 19.04963200715648