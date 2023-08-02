package main

import (
	"fmt"
	"googlemaps.github.io/maps"
)

func encode(coords [][]float64) (string, error) {
	path := make([]maps.LatLng, len(coords))

	for i, coord := range coords {
		path[i] = maps.LatLng{Lat: coord[0], Lng: coord[1]}
	}

	return maps.Encode(path), nil
}

func decode(encoded string) ([][]float64, error) {
	path, err := maps.DecodePolyline(encoded)
	if err != nil {
		return nil, err
	}

	coords := make([][]float64, len(path))
	for i, point := range path {
		coords[i] = []float64{point.Lat, point.Lng}
	}

	return coords, nil
}

func main() {
	encoded, err := encode([][]float64{{37.7749, -122.4194}, {34.0522, -118.2437}})
	if err != nil {
		panic(err)
	}
	fmt.Println("Encoded: ", encoded)

	decoded, err := decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded: ", decoded)
}