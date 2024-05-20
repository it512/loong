package schema

import (
	"github.com/google/uuid"
)

func uid() uuid.UUID {
	return UUIDv7()
}

func UUIDv7() uuid.UUID {
	id := uuid.Must(uuid.NewV7())
	return id
}

func NewID() string {
	return UUIDv7().String()
}
