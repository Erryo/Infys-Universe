package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Erryo/Infys-Universe/types"
	"github.com/labstack/echo/v4"
)

const API_CALL string = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

// Delay represents the minimum amount of minutes required to pass
func isTimeOk(delay int, c echo.Context) bool {
	cookie, err := c.Cookie("time")
	if err != nil {
		c.Logger().Print(err)
		return false
	}

	cookieTime, err := time.Parse(time.RFC3339, cookie.Value)
	if err != nil {
		c.Logger().Print(err)
		return false
	}

	return time.Since(cookieTime) > (time.Minute * 10)
}

func getWeather(key, city string) (error, types.CityData) {
	if key == "" {
		key = os.Getenv("OPENWEATHER_API_KEY")
	}

	var cityData types.CityData
	query := fmt.Sprintf(API_CALL, city, key)
	response, err := http.Get(query)
	if err != nil {
		return err, cityData
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err, cityData
	}

	err = json.Unmarshal(body, &cityData)
	if err != nil {
		return err, cityData
	}

	fmt.Printf("The current temperature in %s is %.2f°C\n", cityData.Name, cityData.Main.Temp)
	return nil, cityData
}
