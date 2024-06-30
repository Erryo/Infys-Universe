package api

import (
	"os"

	"github.com/labstack/echo/v4"
)

func UnAuthRedirect(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if GetClaim(c) == "" {
			c.Redirect(303, "/")
			return nil
		}
		return next(c)
	}
}

func SetUpRoutes(app *echo.Echo, dbh *DbHandler) {
	app.Static("/static", "static")
	app.GET("/favicon.ico", echo.StaticFileHandler("favicon.ico", os.DirFS("static/images/")))

	// Public
	app.GET("/", ShowAbout)
	app.GET("/about", ShowAbout)
	app.GET("/signUp", ShowSignUp)
	app.GET("/signIn", ShowSignIn)

	// Private
	app.GET("/user/home", ShowHome, UnAuthRedirect)
	app.GET("/user/tictactoe", ShowTTT, UnAuthRedirect)
	app.GET("/user/account", ShowAccount, UnAuthRedirect)
	app.GET("/user/schedule", ShowSchedule, UnAuthRedirect)

	// User Management
	app.GET("/LogOut", LogOut)
	app.GET("/Delete", dbh.DeleteUser)
	app.POST("/SignUp", dbh.SignUp)
	app.POST("/SignIn", dbh.SignIn)
}
