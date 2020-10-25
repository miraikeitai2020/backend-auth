package database

import(
	"fmt"

	// import gorm library
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"github.com/miraikeitai2020/backend-auth/pkg/config"
	"github.com/miraikeitai2020/backend-auth/pkg/server/model"
)

const(
	DB_DRIVER_MYSQL = "mysql"
	// SQL Query
	QUERY_FORMAT_INSERT_USER_INFO = "INSERT INTO `users` (id, pass) VALUES (?, ?)"
	QUERY_FORMAT_GET_USER_INFO = "SELECT * FROM `users` WHERE id = ?"
)

type DB struct {
	DB	*gorm.DB
}

func Init() (*DB, error) {
	token := config.GetConnectionToken()
	for i:=0; i<5; i++ {
		if db, err := gorm.Open(DB_DRIVER_MYSQL, token);err == nil {
			return &DB{DB: db}, nil
		}
		config.BackOff(i)
	}
	return nil, fmt.Errorf("Faild connection database")
}

func (db *DB) InsertUserInfo(id, pass string) {
	db.DB.Exec(QUERY_FORMAT_INSERT_USER_INFO, id, pass)
}

func (db *DB) GetUserInfo(id string) (info model.User) {
	db.DB.Raw(QUERY_FORMAT_GET_USER_INFO, id).Scan(&info)
	return
}
