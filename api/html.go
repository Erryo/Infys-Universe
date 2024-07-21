package api

import (
	"time"

	"github.com/Erryo/Infys-Universe/types"
	"github.com/Erryo/Infys-Universe/ui"
	"github.com/Erryo/Infys-Universe/ui/about"
	"github.com/Erryo/Infys-Universe/ui/user"
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
	username := GetClaim(c)
	if isTimeOk(10, c) {
		_, weather := getWeather("", "Goettingen")
		weather.Name = "Goettingen"
		c.SetCookie(createTimeCookie(time.Now().Format(time.RFC3339)))
		return Render(c, user.Home(username, weather))
	}
	return Render(c, user.Home(username, types.CityData{}))
}

func ShowTTT(c echo.Context) error {
	return c.String(200, "Username is:"+GetClaim(c))
}

func ShowAccount(c echo.Context) error {
	return c.String(200, "Username is:"+GetClaim(c))
}

func ShowSchedule(c echo.Context) error {
	return c.String(200, "Username is:"+GetClaim(c))
}
