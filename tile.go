package tile

import (
	"math"
	"sort"
)

func TileForPoint(longitude, latitude float64) int64 {
	x := lon2X(longitude)
	y := lat2Y(latitude)

	return xy2Tile(x, y)
}

func TilesForArea(minLongitude, minLatitude, maxLongitude, maxLatitude float64) []int64 {
	minX := lon2X(minLongitude)
	maxX := lon2X(maxLongitude)
	minY := lat2Y(minLatitude)
	maxY := lat2Y(maxLatitude)

	var tiles []int64

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			tile := xy2Tile(x, y)
			tiles = append(tiles, tile)
		}
	}

	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i] < tiles[j]
	})

	return tiles
}

func lon2X(lon float64) int64 {
	n := math.Round((lon + 180.0) * 65535.0 / 360.0)
	return int64(n + 0.5)
}

func lat2Y(lat float64) int64 {
	n := math.Round((lat + 90.0) * 65535.0 / 180.0)
	return int64(n + 0.5)
}

func xy2Tile(x, y int64) int64 {
	tile := int64(0)

	for i := 15; i >= 0; i-- {
		tile = (tile << 1) | ((x >> i) & 1)
		tile = (tile << 1) | ((y >> i) & 1)
	}

	return tile
}
