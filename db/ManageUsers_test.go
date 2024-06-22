package db

import (
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

func TestCreateUser(t *testing.T) {
	db := ConnectDB()

	properUser := types.User{
		Username:  "John Doe",
		Password:  "John Doe's password",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Now(),
	}

	missingPassword := types.User{
		Username:  "John Doe",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Now(),
	}
	missingUsername := types.User{
		Password:  "John Doe's password",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Now(),
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
		CreatedAt: time.Now(),
	}

	missingUsername := types.User{
		Password:  "John Doe's password",
		Email:     "johndoe@gmail.com",
		CreatedAt: time.Now(),
	}
	emptyUser := types.User{}
	longUsername := types.User{
		Username: "123456789012345678901234567890",
		Password: "Idk",
	}
	testCases := []tCU{
		{user: properUser, desc: "Test a proper user", wantErr: nil, wantUser: properUser},
		{user: emptyUser, desc: "Test a emptyUser user", wantErr: types.ErrFieldEmpty, wantUser: types.User{}},
		{user: longUsername, desc: "Test a long username user", wantErr: types.ErrValueTooLong, wantUser: types.User{}},
		{user: missingUsername, desc: "Test a missing username user", wantErr: types.ErrFieldEmpty, wantUser: types.User{}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			user, err := GetUser(db, tC.user.Username)
			if err != tC.wantErr {
				t.Fatal(err)
			}
			if user != tC.wantUser {
				t.Fatal("Wrong user")
			}
		})
	}
}
