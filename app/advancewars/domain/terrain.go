package domain

type TerrainTile int

const (
	// Sea is only passable by units that have a Floating or Flying chassis
	Sea TerrainTile = iota
	// Land is accessible by ground and flying units
	Land
	// Road has the smallest movement penalty for wheeled vehicles
	Road
)
