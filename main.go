package main

import (
	"landtick/database"
	"landtick/pkg/mysql"
	"landtick/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file")
	}
	e := echo.New()
	mysql.DatabaseInit()
	database.RunMigration()

	e.Static("/uploads", "./uploads")

	routes.Routeinit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start("localhost:5000"))
}
