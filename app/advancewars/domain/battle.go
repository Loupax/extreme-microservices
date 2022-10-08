package domain

import "github.com/google/uuid"

type BattleID struct {
	value string
}

func NewBattleID() BattleID {
	return BattleID{
		value: uuid.New().String(),
	}
}

// Everything below this line can be generated
func (g BattleID) String() string {
	return g.value
}

func (u BattleID) MarshalJSON() ([]byte, error) {
	return []byte(u.value), nil
}

func (u *BattleID) UnmarshalJSON(b []byte) error {
	u = &BattleID{value: string(b)}
	return nil
}
