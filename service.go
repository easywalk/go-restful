package easywalk

import (
	"github.com/easywalk/go-restful/easywalk/service"
	"github.com/google/uuid"
)

type SimplyServiceInterface[T SimplyEntityInterface] interface {
	Create(data T) (*T, error)
	UpdateByID(id uuid.UUID, mapFields map[string]any) (*T, error)
	DeleteByID(id uuid.UUID) (int64, error)

	ReadByID(id uuid.UUID) (*T, error)
	FindAll(mapFields map[string]any) ([]T, error)
}

func NewGenericService[T SimplyEntityInterface](repo SimplyRepositoryInterface[T]) SimplyServiceInterface[T] {
	return &service.SimplyService[T]{Repo: repo}
}
