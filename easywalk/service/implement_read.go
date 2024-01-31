package service

import "github.com/google/uuid"

func (s SimplyService[T]) ReadByID(id uuid.UUID) (*T, error) {
	data, err := s.Repo.Read(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
