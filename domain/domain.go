package domain

type Terrain int

const (
	// Sea is only passable by units that have a Floating or Flying chassis
	Sea Terrain = iota
	// Land is accessible by ground and flying units
	Land
)

type Chassis int

const (
	Wheels Chassis = iota
	Threads
	Wings
	BoatFrame
)
