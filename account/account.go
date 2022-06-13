package account

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	Email    string `gorm:"primarykey"`
	Password string // base64(password)
}

func get_db() (db *gorm.DB, err error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8",
	}), &gorm.Config{})
}
