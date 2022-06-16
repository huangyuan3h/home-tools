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

/**
* register the account to
**/
func Register(email string, password string) (token string, err error) {
	var t = ""
	// todo: add password base 64
	// save to db
	// generate token
	// return token
	return t, nil
}

func Login(email string, password string) (token string, err error) {
	var t = ""
	// todo:
	// search user name and encoded password
	// compare the password and email
	// generate token
	// return token
	return t, nil

}

// simple version of reset password, this is not correct process flow
func ResetPassword(email string, password string) (token string, err error) {
	var t = ""
	// todo:
	// search user name and encoded password
	// compare the password and email
	// generate token
	// return token
	return t, nil
}
