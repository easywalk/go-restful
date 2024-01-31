package service

import (
	"database/sql"
	go_easywalk "github.com/easywalk/go/restful"
	"github.com/google/uuid"
)

type SimplyService[T go_easywalk.SimplyEntityInterface] struct {
	Repo go_easywalk.SimplyRepositoryInterface[T]
}

func (s SimplyService[T]) Create(data T) (*T, error) {
	id, err := s.Repo.Create(data)
	if err != nil {
		return nil, err
	}

	// get data from db
	data, err = s.Repo.Read(id)
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (s SimplyService[T]) UpdateByID(id uuid.UUID, mapFields map[string]any) (*T, error) {
	affected, err := s.Repo.Update(id, mapFields)
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, sql.ErrNoRows
	}

	return s.ReadByID(id)
}

func (s SimplyService[T]) DeleteByID(id uuid.UUID) (int64, error) {
	affected, err := s.Repo.Delete(id)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (s SimplyService[T]) FindAll(mapFields map[string]any) ([]T, error) {
	data, err := s.Repo.FindAll(mapFields)
	if err != nil {
		return nil, err
	}

	return data, nil
}
