package domain

import "github.com/google/uuid"

type GameID struct {
	value string
}

func NewGameID() GameID {
	return GameID{
		value: uuid.New().String(),
	}
}

func (g GameID) String() string {
	return g.value
}

func (u GameID) MarshalJSON() ([]byte, error) {
	return []byte(u.value), nil
}

func (u *GameID) UnmarshalJSON(b []byte) error {
	u = &GameID{value: string(b)}
	return nil
}
