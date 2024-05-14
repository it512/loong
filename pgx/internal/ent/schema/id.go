package schema

import (
	"github.com/google/uuid"
)

func uid() uuid.UUID {
	id := uuid.Must(uuid.NewV7())
	return id
}

func NewID() string {
	return uid().String()
}
