package api

import (
	"github.com/Erryo/Infys-Universe/ui"
	"github.com/labstack/echo/v4"
)

func ShowAbout(c echo.Context) error {
	lang := c.QueryParam("lang")
	c.Logger().Printf("My langs is %v", lang)
	if lang == "ro" {
		return Render(c, ui.RoAboutMe())
	}
	if lang == "de" {
		return Render(c, ui.DeAboutMe())
	}
	return Render(c, ui.AboutMe())
}
