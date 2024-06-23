package api

import (
	"github.com/Erryo/Infys-Universe/ui"
	"github.com/labstack/echo/v4"
)

func ShowAbout(c echo.Context) error {
	return Render(c, ui.AboutMe())
}
