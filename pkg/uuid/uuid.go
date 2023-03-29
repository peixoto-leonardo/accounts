package uuid

import "github.com/google/uuid"

type UUID string

func (u UUID) String() string {
	return string(u)
}

func New() UUID {
	return UUID(uuid.New().String())
}
