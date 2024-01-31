package easywalk

import (
	"github.com/google/uuid"
	"time"
)

type SimplyEntityInterface interface {
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
}
