package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dbName := os.Getenv("DBNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	connStr := fmt.Sprintf("postgres://postgres:%v@localhost:5432/%v?sslmode=disable", password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected succesfully")
	return db
}

func CreateUsersTable(db *sql.DB) error {
	userQuery := `                        
    CREATE TABLE IF NOT EXISTS users(     
        username VARCHAR(25) PRIMARY KEY, 
        password VARCHAR(25) NOT NULL,             
        email VARCHAR(25) NOT NULL,                
        createdAt VARCHAR(24) NOT NULL                   
    )                                     
`
	_, err := db.Exec(userQuery)
	if err != nil {
		return err
	}
	return nil
}

func CreateUserSubjectLinkTable(db *sql.DB) error {
	userSubjectLinkQuery := `                        
    CREATE TABLE IF NOT EXISTS userSubjectLink(      
    username VARCHAR(25),                            
    subject VARCHAR(25),                             
    PRIMARY KEY(username,subject),                   
    FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE
    )                                                
    `
	_, err := db.Exec(userSubjectLinkQuery)
	if err != nil {
		return err
	}
	return nil
}

func CreateLessonsTable(db *sql.DB) error {
	lessonsQuery := `                                 
    CREATE TABLE IF NOT EXISTS lessons(               
    id SMALLSERIAL PRIMARY KEY,                       
    day VARCHAR(15),                                  
    lno SMALLINT,                                     
    name VARCHAR(25),                                 
    startAt VARCHAR(25),                                 
    endAt VARCHAR(25),                                
    username varchar(25),                             
    FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE, 
    UNIQUE(day,lno,username)                          
    )                                                 
    `

	_, err := db.Exec(lessonsQuery)
	if err != nil {
		return err
	}
	return nil
}

//func CreateUserLessonLinkTable(db *sql.DB) error {
//	userLessonLinkQuery := `
//    CREATE TABLE IF NOT EXISTS userLessonLink(
//    username VARCHAR(25),
//    lid SMALLINT,
//    PRIMARY KEY(username,LID),
//    FOREIGN KEY(username) REFERENCES users(username),
//    FOREIGN KEY(lid) REFERENCES lessons(id)
//    )
//    `
//	_, err := db.Exec(userLessonLinkQuery)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
