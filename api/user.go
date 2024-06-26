package api

import (
	"database/sql"
	"time"

	"github.com/Erryo/Infys-Universe/db"
	"github.com/Erryo/Infys-Universe/types"
	"github.com/labstack/echo/v4"
)

type DbHandler struct {
	db *sql.DB
}

func NewDbHandler(db *sql.DB) *DbHandler {
	return &DbHandler{
		db: db,
	}
}

func (dbh *DbHandler) SignUp(c echo.Context) error {
	user := types.User{}
	user.Username = c.FormValue("username")
	user.Email = c.FormValue("email")
	password := c.FormValue("password")
	user.CreatedAt = time.Now().Format(time.TimeOnly)
	user.Password = password // Encrypt

	c.Logger().Printf("User:%v", user)
	err := db.CreateUser(dbh.db, user)
	if err == nil {
		c.Logger().Print("Success")
		c.Response().Header().Set("HX-Redirect", "/SignIn")
		c.Response().WriteHeader(200)
		c.Redirect(303, "/SignIn")
		return nil
	}
	if err == types.ErrDuplicateKey {
		return c.String(200, "User already exists")
	}
	return err
}
