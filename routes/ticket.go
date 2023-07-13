package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func TicketRoutes(e *echo.Group) {
	r := repository.RepositoryTicket(mysql.DB)
	h := handlers.TicketHandler(r)

	e.GET("/tickets", h.FindTicket)
	e.GET("/ticket/:id", h.GetTicket)
	e.POST("/ticket", h.CreateTicket)
	e.PATCH("/ticket/:id", h.UpdateTicket)
	e.DELETE("/ticket/:id", h.DeleteTicket)
}
