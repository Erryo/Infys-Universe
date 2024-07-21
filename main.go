package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Erryo/Infys-Universe/api"
	"github.com/Erryo/Infys-Universe/db"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	isDev = os.Getenv("DEV")
	fmt.Println("Dev:", isDev)

	app := echo.New()
	app.Use(DisableCache)

	dB := db.ConnectDB()
	db.CreateUsersTable(dB)
	db.CreateUserSubjectLinkTable(dB)
	db.CreateLessonsTable(dB)
	defer dB.Close()
	dbh := api.NewDbHandler(dB)

	api.SetUpRoutes(app, dbh)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(port))
}
