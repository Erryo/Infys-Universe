package main

import (
	"fmt"
	"os"

	"github.com/Erryo/Infys-Universe/api"
	"github.com/Erryo/Infys-Universe/db"
	"github.com/labstack/echo/v4"
)

var isDev string

func DisableCache(next echo.HandlerFunc) echo.HandlerFunc {
	if isDev != "TRUE" {
		return next
	}
	return echo.HandlerFunc(func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-store")
		return next(c)
	})
}

func main() {
	isDev = os.Getenv("DEV")
	fmt.Println("Dev:", isDev)

	app := echo.New()

	dB := db.ConnectDB()
	db.CreateUsersTable(dB)
	db.CreateUserSubjectLinkTable(dB)
	db.CreateLessonsTable(dB)
	defer dB.Close()

	app.Use(DisableCache)
	app.Static("/static", "static")
	app.GET("/", api.ShowAbout)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app.Logger.Fatal(app.Start(port))
}
