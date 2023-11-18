package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	ID        uint
	Objective string
	Completed bool
}

func NewTodo(objective string) *Todo {
	return &Todo{Objective: objective, Completed: false}
}
