package main

import (
	"context"
	"go-mongodb/db"
	"go-mongodb/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	client := db.ConnectMongo()
	defer client.Disconnect(context.TODO())

	e := echo.New()
	routes.Routes(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))
}
