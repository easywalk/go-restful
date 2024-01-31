package easywalk

import (
	"github.com/google/uuid"
	"time"
)

type SimplyEntityInterface interface {
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	SetCreatedAt(time time.Time)
	SetUpdatedAt(time time.Time)
}
