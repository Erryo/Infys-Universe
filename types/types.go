package types

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Username  string `sql:"username"`
	Password  string `sql:"Password"`
	Email     string `sql:"email"`
	CreatedAt string `sql:"createdAt"`
}
type Lesson struct {
	Id        int
	Day       string
	Lno       int
	Name      string
	StartTime string
	EndTime   string
	Username  string
}
type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}
