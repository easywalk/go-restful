package easywalk

import "github.com/google/uuid"

type SimplyEntityInterface interface {
	GetID() uuid.UUID
	SetID(id uuid.UUID)
}
