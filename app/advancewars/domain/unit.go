package domain

type Unit int

const (
	Infantry Unit = iota
	TransportHelicopter
	Submarine
	RocketLauncher
	Recon
	Mech
	WarTank
	LightTank
	Lander
	Fighter
	Cruiser
	Battleship
	Bomber
	BattleHelicopter
	Artillery
	APC
	AntiAirTank
	AntiAirMissileLauncher
)

type UnitInstance struct {
	Unit
	X  int
	Y  int
	HP int
}
