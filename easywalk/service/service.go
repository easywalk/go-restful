package service

import (
	"github.com/easywalk/go-restful"
	"github.com/easywalk/go-restful/easywalk/repository"
	"github.com/google/uuid"
)

type SimplyServiceInterface[T easywalk.SimplyEntityInterface] interface {
	Create(data T) (*T, error)
	UpdateByID(id uuid.UUID, mapFields map[string]any) (*T, error)
	DeleteByID(id uuid.UUID) (int64, error)

	ReadByID(id uuid.UUID) (*T, error)
	FindAll(mapFields map[string]any) ([]T, error)
}

func NewGenericService[T easywalk.SimplyEntityInterface](repo repository.SimplyRepositoryInterface[T]) SimplyServiceInterface[T] {
	return &SimplyService[T]{Repo: repo}
}
