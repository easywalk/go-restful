package easywalk

import (
	"github.com/google/uuid"
	"time"
)

type SimplyEntityInterface interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}
