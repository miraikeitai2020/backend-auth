package database

import(
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

const(
	GROM_TEST_DEFAULT_ID = "id"
	GROM_TEST_DEFAULT_PASS = "password"
)

func InitializeDBMoc() (*DB, sqlmock.Sqlmock, error){
	db, mock, err := sqlmock.New()
    if err != nil {
        return nil, nil, err
    }

    gdb, err := gorm.Open("mysql", db)
    if err != nil {
        return nil, nil, err
    }
    return &DB{DB: gdb}, mock, nil
}

func TestGetUserInfo(t *testing.T) {
	db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
	}
	defer db.DB.Close()
	db.DB.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta(
        "SELECT * FROM `users` WHERE id = ?")).
        WithArgs(GROM_TEST_DEFAULT_ID).
        WillReturnRows(
            sqlmock.NewRows([]string{"id", "pass"}).
			AddRow(GROM_TEST_DEFAULT_ID, GROM_TEST_DEFAULT_PASS),
		)
	
	result := db.GetUserInfo("id")
	if result.ID != "id" {
		t.Errorf("id value is not `id`")
	}
	if result.Pass != "password" {
		t.Errorf("pass value is not `password`")
	}
}

func TestInsertUserInfo(t *testing.T) {
	db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
	}
	defer db.DB.Close()
	db.DB.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta(
        "INSERT INTO `users` (id, pass) VALUES (?, ?)")).
        WithArgs(GROM_TEST_DEFAULT_ID, GROM_TEST_DEFAULT_PASS)
	
	db.InsertUserInfo("id", "pass")
}
