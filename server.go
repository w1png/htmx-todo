package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/w1png/htmx/config"
	"github.com/w1png/htmx/handlers"
)

type HttpServer struct {
	echo *echo.Echo
}

func NewHttpServer() *HttpServer {
	server := &HttpServer{
		echo: echo.New(),
	}

	server.GatherRoutes()

	return server
}

func (s *HttpServer) GatherRoutes() {
	s.echo.GET("/", handlers.IndexHandler)
	s.echo.PUT("/todo/:id", handlers.MarkTodoHandler)
	s.echo.DELETE("/todo/:id", handlers.DeleteTodoHandler)

	s.echo.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	s.echo.POST("/create", handlers.CreateTodoHandler)

	s.echo.Use(middleware.Logger())
}

func (s *HttpServer) Run() error {
	return s.echo.Start(fmt.Sprintf(":%s", config.ConfigInstance.ServerPort))
}
