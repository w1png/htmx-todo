package handlers

import (
	"html/template"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/w1png/htmx/errors"
	"github.com/w1png/htmx/models"
	"github.com/w1png/htmx/storage"
)

func sendTodos(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("public/templates/index.html"))

	todos, err := storage.StorageInstance.GetTodos()
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(c.Response(), "Todos", todos)
	if err != nil {
		return errors.NewTemplateParsingError(err.Error())
	}

	return nil
}

func IndexHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("public/templates/index.html"))

	todos, err := storage.StorageInstance.GetTodos()
	if err != nil {
		return err
	}

	err = tmpl.Execute(c.Response(), todos)
	if err != nil {
		return errors.NewTemplateParsingError(err.Error())
	}

	return nil
}

func CreateTodoHandler(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}

	objective := c.FormValue("objective")

	todo := models.NewTodo(objective)

	err = storage.StorageInstance.SaveTodo(todo)
	if err != nil {
		return err
	}

	return sendTodos(c)
}

func MarkTodoHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	todo, err := storage.StorageInstance.GetTodo(uint(id))
	if err != nil {
		return err
	}

	todo.Completed = !todo.Completed

	err = storage.StorageInstance.SaveTodo(todo)
	if err != nil {
		return err
	}

	return sendTodos(c)
}

func DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	err = storage.StorageInstance.DeleteTodo(uint(id))
	if err != nil {
		return err
	}

	return sendTodos(c)
}
