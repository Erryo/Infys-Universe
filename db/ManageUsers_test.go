package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/Erryo/Infys-Universe/types"
)

type tCU struct {
	user     types.User
	desc     string
	wantErr  error
	wantUser types.User
}

type tCL struct {
	lesson     types.Lesson
	desc       string
	wantErr    error
	wantLesson []types.Lesson
}

func TestCreateUser(t *testing.T) {
	db := ConnectDB()

	properUser := types.User{
		Username:  "John Doe",
		Password:  "John Doe's password",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.DateOnly),
	}

	missingPassword := types.User{
		Username:  "John Doe",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Now().Format(time.DateOnly),
	}
	missingUsername := types.User{
		Password:  "John Doe's password",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Now().Format(time.DateOnly),
	}
	emptyUser := types.User{}
	longUsername := types.User{
		Username: "123456789012345678901234567890",
		Password: "Idk",
	}
	longPassword := types.User{
		Username: "Idk",
		Password: "123456789012345678901234567890",
	}
	testCases := []tCU{
		{user: properUser, desc: "Test a proper user", wantErr: nil},
		{user: emptyUser, desc: "Test a emptyUser user", wantErr: types.ErrFieldEmpty},
		{user: longUsername, desc: "Test a long username user", wantErr: types.ErrValueTooLong},
		{user: longPassword, desc: "Test a long username user", wantErr: types.ErrValueTooLong},
		{user: missingUsername, desc: "Test a missing username user", wantErr: types.ErrFieldEmpty},
		{user: missingPassword, desc: "Test a missing password user", wantErr: types.ErrFieldEmpty},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := CreateUser(db, tC.user)
			if err != tC.wantErr {
				t.Fatal(err)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	db := ConnectDB()

	properUser := types.User{
		Username:  "John Doe",
		Password:  "John Doe's password",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.DateOnly),
	}

	missingUsername := types.User{
		Password:  "John password missing",
		Email:     "johndoe1@gmail.com",
		CreatedAt: time.Now().Format(time.DateOnly),
	}
	emptyUser := types.User{}
	longUsername := types.User{
		Username: "123456789012345678901234567890",
		Password: "Idk",
	}
	wrongUser := types.User{
		Username: "NOT ME A",
	}
	testCases := []tCU{
		{user: properUser, desc: "Test a proper user", wantErr: nil, wantUser: properUser},
		{user: emptyUser, desc: "Test a emptyUser user", wantErr: types.ErrFieldEmpty, wantUser: types.User{}},
		{user: longUsername, desc: "Test a long username user", wantErr: types.ErrValueTooLong, wantUser: types.User{}},
		{user: missingUsername, desc: "Test a missing username user", wantErr: types.ErrFieldEmpty, wantUser: types.User{}},
		{user: wrongUser, desc: "Get a wrong User ", wantErr: types.ErrNoRows},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			user, err := GetUser(db, tC.user.Username)
			if err != tC.wantErr {
				t.Fatal(err)
			}
			if user != tC.wantUser {
				fmt.Println(user)
				fmt.Println(properUser)
				t.Fatal("Wrong user")
			}
		})
	}
}

func TestCreateLesson(t *testing.T) {
	db := ConnectDB()

	proper := types.Lesson{
		Id:        0,
		Day:       "Monday",
		Lno:       1,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		Username:  "John Doe",
	}
	p2roper := types.Lesson{
		Id:        0,
		Day:       "Sunday",
		Lno:       2,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		Username:  "John Doe",
	}
	p3roper := types.Lesson{
		Id:        0,
		Day:       "Monday",
		Lno:       2,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		Username:  "John Doe",
	}
	noDay := types.Lesson{
		Id:        0,
		Lno:       1,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
	}

	empty := types.Lesson{}

	longDay := types.Lesson{
		Id:        0,
		Day:       "123456789012345678901234567890",
		Lno:       1,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
	}

	testCases := []tCL{
		{lesson: proper, desc: "Test a proper Lessson", wantErr: nil},
		{lesson: p2roper, desc: "Test a 2nd proper Lessson", wantErr: nil},
		{lesson: p3roper, desc: "Test a 2nd proper Lessson", wantErr: nil},
		{lesson: proper, desc: "Test a repeating Lessson", wantErr: types.ErrDuplicateKey},
		{lesson: noDay, desc: "Test a missing Day Lessson", wantErr: types.ErrFieldEmpty},
		{lesson: empty, desc: "Test a empty Lessson", wantErr: types.ErrFieldEmpty},
		{lesson: longDay, desc: "Test a long field Lesson", wantErr: types.ErrValueTooLong},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := CreateLesson(db, tC.lesson)
			err = PgerrorTransform(err)
			if err != tC.wantErr {
				t.Fatal(err)
			}
		})
	}
}

func TestGetLesson(t *testing.T) {
	db := ConnectDB()

	proper := types.Lesson{
		Id:        1,
		Day:       "Monday",
		Lno:       1,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),

		Username: "John Doe",
	}
	propers := []types.Lesson{proper}
	noUser := types.Lesson{
		Id:        0,
		Lno:       1,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
	}

	empty := types.Lesson{}

	longUser := types.Lesson{
		Id:        0,
		Username:  "123456789012345678901234567890",
		Lno:       1,
		Name:      "Romanian",
		StartTime: time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
		EndTime:   time.Date(time.Now().Year(), time.Now().Month(), 22, 19, 30, 0, 0, &time.Location{}).Format(time.TimeOnly),
	}

	wrongUser := types.Lesson{
		Username: "NOT ME",
	}
	testCases := []tCL{
		{lesson: proper, desc: "Get a proper Lessson", wantErr: nil, wantLesson: propers},
		{lesson: proper, desc: "Get a repeating Lessson", wantErr: nil, wantLesson: propers},
		{lesson: noUser, desc: "Get a missing user Lessson", wantErr: types.ErrFieldEmpty},
		{lesson: empty, desc: "Get a empty Lessson", wantErr: types.ErrFieldEmpty},
		{lesson: longUser, desc: "Get a long field Lesson", wantErr: types.ErrValueTooLong},
		{lesson: wrongUser, desc: "Get a wrong User Lesson", wantErr: nil},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lessons, err := GetLesson(db, tC.lesson.Username)
			if err != tC.wantErr {
				t.Fatal(err)
			}
			fmt.Println(lessons)
			fmt.Println()
		})
	}
}
