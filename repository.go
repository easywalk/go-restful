package go_easywalk

import (
	"github.com/easywalk/go/restful/easywalk/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SimplyRepositoryInterface[T SimplyEntityInterface] interface {
	Create(data T) (uuid.UUID, error)
	Update(id uuid.UUID, mapFields map[string]any) (int64, error)
	Delete(id uuid.UUID) (int64, error)

	// Query
	Read(id uuid.UUID) (T, error)
	ReadAll() ([]T, error)
	FindAll(mapFields map[string]any) ([]T, error)
}

// NewSimplyRepository is a factory method for create new simplyRepository
// @param db - pointer to gorm DB
// @return pointer to simplyRepository
// @example
//
//	repo := NewSimplyRepository(db)
func NewSimplyRepository[T SimplyEntityInterface](db *gorm.DB) SimplyRepositoryInterface[T] {
	var table T
	db.AutoMigrate(&table)
	return &repository.SimplyRepository[T]{DB: db}
}
