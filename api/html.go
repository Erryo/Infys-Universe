package api

import (
	"github.com/Erryo/Infys-Universe/ui"
	"github.com/labstack/echo/v4"
)

func ShowAbout(c echo.Context) error {
	lang := c.QueryParam("lang")
	if lang == "ro" {
		return Render(c, ui.RoAboutMe())
	}
	if lang == "de" {
		return Render(c, ui.DeAboutMe())
	}
	return Render(c, ui.AboutMe())
}

func ShowSignUp(c echo.Context) error {
	return Render(c, ui.SignUp())
}

func ShowSignIn(c echo.Context) error {
	return Render(c, ui.SignIn())
}
