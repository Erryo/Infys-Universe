package db

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/Erryo/Infys-Universe/types"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func PgerrorTransform(err error) error {
	if err == nil {
		return nil
	}
	fmt.Println(err)
	pgerr, ok := err.(*pq.Error)
	if ok {
		if pgerr.Code == "23505" {
			return types.ErrDuplicateKey
		}
		if pgerr.Code == "23503" || pgerr.Code == "42830" {
			return types.ErrDuplicateKey
		}
	}

	if err == sql.ErrNoRows {
		return types.ErrNoRows
	}

	return err
}

func CreateUser(db *sql.DB, user types.User) error {
	v := reflect.ValueOf(user)

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).String()
		if value == "" {
			return types.ErrFieldEmpty
		}
		if len(value) > 25 {
			fmt.Println(value)
			return types.ErrValueTooLong
		}
	}
	// Hash

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	query := `                                           
    INSERT INTO users(username,password,email,createdat) 
    VAlues($1,$2,$3,$4)                                  
    `
	_, err = db.Exec(query, user.Username, user.Password, user.Email, user.CreatedAt)
	if err != nil {
		err = PgerrorTransform(err)
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
	err := db.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		err = PgerrorTransform(err)
		return types.User{}, err
	}

	return user, nil
}

func CreateLesson(db *sql.DB, lesson types.Lesson) error {
	v := reflect.ValueOf(lesson)

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).String()
		if value == "" {
			fmt.Println(value)
			return types.ErrFieldEmpty
		}
		if len(value) > 25 {
			return types.ErrValueTooLong
		}
	}

	query := `                                                                         
    INSERT INTO lessons(day,lno,name,startat,endat,username)
    VAlues($1,$2,$3,$4,$5,$6) RETURNING id                                                          
    `
	var id int
	err := db.QueryRow(query, lesson.Day, lesson.Lno, lesson.Name, lesson.StartTime, lesson.EndTime, lesson.Username).Scan(&id)
	fmt.Print(id)
	if err != nil {
		err = PgerrorTransform(err)
		return err
	}
	return nil
}

func GetLesson(db *sql.DB, username string) ([]types.Lesson, error) {
	if username == "" {
		return nil, types.ErrFieldEmpty
	}
	if len(username) > 25 {
		return nil, types.ErrValueTooLong
	}

	var lessons []types.Lesson
	query := `                                                                                      
    SELECT * FROM lessons WHERE username=$1`
	// err := db.QueryRow(query, username).Scan(&lesson.Id, &lesson.Day, &lesson.Lno, &lesson.Name, &lesson.StartTime, &lesson.EndTime, &lesson.Username)

	rows, err := db.Query(query, username)
	if err != nil {
		err = PgerrorTransform(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		lesson := types.Lesson{}
		rows.Scan(&lesson.Id, &lesson.Day, &lesson.Lno, &lesson.Name, &lesson.StartTime, &lesson.EndTime, &lesson.Username)
		fmt.Println(lesson)
		lessons = append(lessons, lesson)
	}

	fmt.Println(lessons)
	return lessons, nil
}

func CreateSubject(db *sql.DB, subject, username string) error {
	if subject == "" || username == "" {
		return types.ErrFieldEmpty
	}
	if len(subject) > 25 || len(username) > 25 {
		return types.ErrValueTooLong
	}

	query := `
    INSERT INTO userSubjectLink (username,subject) VALUES ($1,$2)
    `
	_, err := db.Exec(query, username, subject)
	if err != nil {
		err = PgerrorTransform(err)
		return err
	}
	return nil
}

func GetSubject(db *sql.DB, username string) ([]string, error) {
	if username == "" {
		return nil, types.ErrFieldEmpty
	}
	if len(username) > 25 {
		return nil, types.ErrValueTooLong
	}
	query := `
    SELECT (subject) FROM userSubjectLink WHERE username=$1
    `
	rows, err := db.Query(query, username)
	if err != nil {
		fmt.Println(err)
		err = PgerrorTransform(err)
		return nil, err
	}

	defer rows.Close()

	var subjects []string
	for rows.Next() {
		var subject string
		rows.Scan(&subject)
		subjects = append(subjects, subject)
	}
	fmt.Println(subjects)
	return subjects, nil
}

func DeleteUser(db *sql.DB, username string) error {
	if username == "" {
		return types.ErrFieldEmpty
	}
	if len(username) > 25 {
		return types.ErrValueTooLong
	}
	query := `
    DELETE FROM users WHERE username=$1
    `
	_, err := db.Exec(query, username)
	if err != nil {
		err = PgerrorTransform(err)
		return err
	}
	return nil
}

func DeleteLesson(db *sql.DB, id int) error {
	if id < 1 {
		return types.ErrInvalidArgument
	}
	if id > 32767 {
		return types.ErrValueTooLong
	}
	query := `
    DELETE FROM lessons WHERE id=$1
    `
	_, err := db.Exec(query, id)
	if err != nil {
		err = PgerrorTransform(err)
		return err
	}
	return nil
}

func DeleteSubject(db *sql.DB, username, subject string) error {
	if username == "" || subject == "" {
		return types.ErrFieldEmpty
	}
	if len(username) > 25 || len(subject) > 25 {
		return types.ErrValueTooLong
	}
	query := `                                
    DELETE FROM userSubjectLink WHERE username=$1 AND subject=$2            
    `
	_, err := db.Exec(query, username, subject)
	if err != nil {
		err = PgerrorTransform(err)
		return err
	}
	return nil
}
