package repository

import (
	go_easywalk "github.com/easywalk/go/restful"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

type SimplyRepository[T go_easywalk.SimplyEntityInterface] struct {
	DB *gorm.DB
}

// Create is a generic method for create operation
// @param data - pointer to entity
// @return uuid of created entity, error
func (r *SimplyRepository[T]) Create(data T) (uuid.UUID, error) {
	data.SetID(uuid.New())
	tx := r.DB.Create(&data)
	if tx.Error != nil {
		log.Printf("Error in repository Create operation - %v", tx.Error)
	}

	return data.GetID(), tx.Error
}

func (r *SimplyRepository[T]) Update(id uuid.UUID, mapFields map[string]any) (int64, error) {

	fromDB, err := r.Read(id)
	if err != nil {
		return 0, err
	}

	// print all fields of T
	for i := 0; i < reflect.TypeOf(fromDB).Elem().NumField(); i++ {
		for key, value := range mapFields {

			lowerFieldName := strings.ToLower(reflect.TypeOf(fromDB).Elem().Field(i).Name)
			lowerKeyName := strings.ToLower(key)

			if lowerFieldName != lowerKeyName {
				continue
			}

			reflect.ValueOf(fromDB).Elem().Field(i).Set(reflect.ValueOf(value))
		}
	}

	// update T
	tx := r.DB.Save(&fromDB)
	if tx.Error != nil {
		log.Printf("Error in repository UpdateByID operation - %v", tx.Error)
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (r *SimplyRepository[T]) Delete(id uuid.UUID) (int64, error) {
	var deleted T
	tx := r.DB.Delete(&deleted, id)
	if tx.Error != nil {
		log.Printf("Error in repository DeleteByID operation - %v", tx.Error)
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}