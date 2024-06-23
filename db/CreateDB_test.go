package db

import "testing"

func TestConnectDb(t *testing.T) {
	_ = ConnectDB()
}

func TestCreateUserTable(t *testing.T) {
	db := ConnectDB()
	err := CreateUsersTable(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateUserSubjTable(t *testing.T) {
	db := ConnectDB()
	err := CreateUserSubjectLinkTable(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateLessonsTable(t *testing.T) {
	db := ConnectDB()
	err := CreateLessonsTable(db)
	if err != nil {
		t.Fatal(err)
	}
}

//func TestCreateUserLessonTable(t *testing.T) {
//	db := ConnectDB()
//	err := CreateUserLessonLinkTable(db)
//	if err != nil {
//		t.Fatal(err)
//	}
//}
