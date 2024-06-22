package db

import (
	"database/sql"
	"reflect"

	"github.com/Erryo/Infys-Universe/types"
)

func CreateUser(db *sql.DB, user types.User) error {
	v := reflect.ValueOf(user)

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).String()
		if value == "" {
			return types.ErrFieldEmpty
		}
		if len(value) > 25 {
			return types.ErrValueTooLong
		}
	}

	query := `                                           
    INSERT INTO users(username,password,email,createdat) 
    VALUES($1,$2,$3,$4)                                  
    `
	_, err := db.Exec(query, user.Username, user.Password, user.Email, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *sql.DB, username string) (types.User, error) {
	if username == "" {
		return types.User{}, types.ErrFieldEmpty
	}
	if len(username) > 25 {
		return types.User{}, types.ErrValueTooLong
	}
	user := types.User{}

	query := `
    SELECT * FROM users WHERE username=$1`
	db.QueryRow(query, username).Scan(&user)
	return user, nil
}
