package api

import (
	"os"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(app *echo.Echo, dbh *DbHandler) {
	app.Static("/static", "static")
	app.GET("/favicon.ico", echo.StaticFileHandler("favicon.ico", os.DirFS("static/images/")))
	app.GET("/", ShowAbout)
	app.GET("/About", ShowAbout)
	app.GET("/SignUp", ShowSignUp)
	app.POST("/SignUp", dbh.SignUp)
	app.GET("/SignIn", ShowSignIn)
}
