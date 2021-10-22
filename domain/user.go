package domain

import "github.com/google/uuid"

type UserID struct {
	value string
}

func NewUserID() UserID {
	return UserID{
		value: uuid.New().String(),
	}
}

func (u UserID) String() string {
	return u.value
}

func (u UserID) MarshalJSON() ([]byte, error) {
	return []byte(u.value), nil
}

func (u *UserID) UnmarshalJSON(b []byte) error {
	u = &UserID{value: string(b)}
	return nil
}
