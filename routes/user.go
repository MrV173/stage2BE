package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	r := repository.RepositoryUser(mysql.DB)
	h := handlers.UserHandler(r)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
}
