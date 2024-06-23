package main

import (
	"os"

	"github.com/Erryo/Infys-Universe/api"
	"github.com/Erryo/Infys-Universe/db"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	dB := db.ConnectDB()
	db.CreateUsersTable(dB)
	db.CreateUserSubjectLinkTable(dB)
	db.CreateLessonsTable(dB)
	defer dB.Close()

	app.GET("/", api.ShowHtml)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(port))
}
