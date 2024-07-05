package main

import (
	"BE23TODO/app/config"
	"BE23TODO/app/databases"
	"BE23TODO/app/migrations"
	"BE23TODO/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbPosgres := databases.InitDBpostgre(cfg)
	migrations.RunMigrations(dbPosgres)

	// create new instance echo
	e := echo.New()

	routes.InitRouter(e, dbPosgres)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
