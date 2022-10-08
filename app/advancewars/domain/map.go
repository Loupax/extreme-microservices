package domain

import "github.com/google/uuid"

type Terrain [][]TerrainTile
type Units []UnitInstance
type MapID struct {
	value string
}

func NewMapID() BattleID {
	return BattleID{
		value: uuid.New().String(),
	}
}

// Everything below this line can be generated
func (g MapID) String() string {
	return g.value
}

func (u MapID) MarshalJSON() ([]byte, error) {
	return []byte(u.value), nil
}

func (u *MapID) UnmarshalJSON(b []byte) error {
	u = &MapID{value: string(b)}
	return nil
}
