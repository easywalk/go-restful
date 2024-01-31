package repository

import (
	"github.com/google/uuid"
	"log"
)

func (r *SimplyRepository[T]) Read(id uuid.UUID) (T, error) {
	var out T
	tx := r.DB.First(&out, id)
	return out, tx.Error
}

func (r *SimplyRepository[T]) ReadAll() ([]T, error) {
	var out []T
	tx := r.DB.Find(&out)
	return out, tx.Error
}

func (r *SimplyRepository[T]) FindAll(mapFields map[string]any) ([]T, error) {
	var out []T
	tx := r.DB.Where(mapFields).Find(&out)
	if tx.Error != nil {
		log.Printf("Error in repository FindAll operation - %v", tx.Error)
	}
	return out, tx.Error
}
