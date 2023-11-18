package storage

import (
	"github.com/w1png/htmx/config"
	"github.com/w1png/htmx/errors"
	"github.com/w1png/htmx/models"
)

type Storage interface {
	SaveTodo(todo *models.Todo) error
	GetTodo(ID uint) (*models.Todo, error)
	GetTodos() ([]*models.Todo, error)
	GetActiveTodos() ([]*models.Todo, error)
	DeleteTodo(ID uint) error
}

var StorageInstance Storage

func InitStorage() error {
	var err error

	switch config.ConfigInstance.StorageType {
	case "sqlite":
		StorageInstance, err = NewSqliteStorage(config.ConfigInstance.SqliteFileName)
	case "in_memory":
		StorageInstance, err = NewSqliteStorage("file::memory:?cache=shared")
	default:
		return errors.NewInvalidStorageTypeError(config.ConfigInstance.StorageType)
	}

	if err != nil {
		return err
	}

	return nil
}
