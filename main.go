package main

import (
	"os"

	"github.com/Erryo/Infys-Universe/api"
	"github.com/Erryo/Infys-Universe/db"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db := db.ConnectDB()
	defer db.Close()

	app.GET("/", api.ShowHtml)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(port))
}
