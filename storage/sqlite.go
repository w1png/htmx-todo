package storage

import (
	"fmt"

	"github.com/w1png/htmx/errors"
	"github.com/w1png/htmx/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteStorage struct {
	db *gorm.DB
}

func NewSqliteStorage(path string) (*SqliteStorage, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	storage := &SqliteStorage{db: db}

	err = storage.autoMigrate()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s SqliteStorage) autoMigrate() error {
	return s.db.AutoMigrate(&models.Todo{})
}

func (s SqliteStorage) SaveTodo(todo *models.Todo) error {
	return s.db.Save(todo).Error
}

func (s SqliteStorage) GetTodo(ID uint) (*models.Todo, error) {
	var todo models.Todo
	result := s.db.Where("ID = ?", ID).First(&todo)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &todo, errors.NewNotFoundError(fmt.Sprintf("todo with ID %d", ID))
		}
		return nil, result.Error
	}

	return &todo, nil
}

func (s SqliteStorage) GetTodos() ([]*models.Todo, error) {
	var todos []*models.Todo
	result := s.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func (s SqliteStorage) GetActiveTodos() ([]*models.Todo, error) {
	var todos []*models.Todo
	result := s.db.Where("completed = ?", false).Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func (s SqliteStorage) DeleteTodo(ID uint) error {
	_, err := s.GetTodo(ID)
	if err != nil {
		return err
	}

	return s.db.Where("ID = ?", ID).Delete(&models.Todo{}).Error
}
