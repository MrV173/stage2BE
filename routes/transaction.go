package routes

import (
	"landtick/handlers"
	"landtick/pkg/middleware"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Group) {
	r := repository.TransactionRepository(mysql.DB)
	h := handlers.TransactionHandler(r)

	e.GET("/transactions", middleware.Auth(h.FindTransaction))
	e.GET("/transaction/:id", middleware.Auth(h.GetTransaction))
	e.POST("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction)))
}
