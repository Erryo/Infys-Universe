package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Erryo/Infys-Universe/db"
	"github.com/Erryo/Infys-Universe/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	user.Password = c.FormValue("password")
	user.CreatedAt = time.Now().Format(time.TimeOnly)

	err := db.CreateUser(dbh.db, user)
	if err == nil {
		c.Response().Header().Set("HX-Redirect", "/signIn")
		c.Response().WriteHeader(200)
		c.Redirect(303, "/SignIn")
		return nil
	}
	if err == types.ErrDuplicateKey {
		return c.String(200, "User already exists")
	}
	return err
}

func (dbh *DbHandler) SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := db.GetUser(dbh.db, username)
	if err == nil {

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return c.String(200, "User password is wrong")
		}

		err, t := CreateJWT(user.Username)
		if err != nil {
			return err
		}

		cookie := &http.Cookie{
			Name:     "auth",
			Value:    t,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 72),
			HttpOnly: true,
		}
		c.SetCookie(cookie)
		c.Response().Header().Set("HX-Redirect", "/user/home")
		c.Response().WriteHeader(200)
		c.Redirect(303, "/user/home")
		return nil
	}
	if err == types.ErrDuplicateKey {
		return c.String(200, "Username is wrong")
	}
	if err == types.ErrNoRows {
		return c.String(200, "User does not exist")
	}
	return err
}

func (dbh *DbHandler) DeleteUser(c echo.Context) error {
	username := GetClaim(c)
	if username == "" {
		return types.ErrFieldEmpty
	}
	err := db.DeleteUser(dbh.db, username)
	if err != nil {
		return err
	}
	DeleteJWT(c)
	c.Redirect(303, "/")
	return nil
}

func LogOut(c echo.Context) error {
	username := GetClaim(c)
	if username == "" {
		return types.ErrFieldEmpty
	}
	DeleteJWT(c)
	c.Redirect(303, "/")
	return nil
}

func CreateJWT(username string) (error, string) {
	claims := &types.JwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	return err, t
}

func GetClaim(c echo.Context) string {
	inter, err := c.Cookie("auth")
	if err != nil {
		return ""
	}
	tokenString := inter.Value

	token, err := jwt.ParseWithClaims(tokenString, &types.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*types.JwtCustomClaims); ok {
		return claims.Name
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return ""
}

func DeleteJWT(c echo.Context) {
	cookie := &http.Cookie{
		Name:    "auth",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}

	c.SetCookie(cookie)
}
