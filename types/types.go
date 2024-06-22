package types

import "time"

type User struct {
	Username  string    `sql:"username"`
	Password  string    `sql:"Password"`
	Email     string    `sql:"email"`
	CreatedAt time.Time `sql:"createdAt"`
}
type Lesson struct {
	Id        int
	Day       string
	Lno       int
	Name      string
	StartTime time.Time
	EndTime   time.Time
}
