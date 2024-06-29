package api

import (
	"github.com/Erryo/Infys-Universe/ui"
	"github.com/Erryo/Infys-Universe/ui/about"
	"github.com/labstack/echo/v4"
)

func ShowAbout(c echo.Context) error {
	lang := c.QueryParam("lang")
	username := GetClaim(c)
	if lang == "ro" {
		return Render(c, about.RoAboutMe(username))
	}
	if lang == "de" {
		return Render(c, about.DeAboutMe(username))
	}
	return Render(c, about.AboutMe(username))
}

func ShowSignUp(c echo.Context) error {
	return Render(c, ui.SignUp())
}

func ShowSignIn(c echo.Context) error {
	return Render(c, ui.SignIn())
}

func ShowHome(c echo.Context) error {
	return c.String(200, "Username is:"+GetClaim(c))
}
