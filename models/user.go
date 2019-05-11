package models

import (
	"fmt"
	"log"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID       string `gorm:"primary_key"`
	Email    string `gorm:"type:varchar(256);unique_index"`
	Username string `gorm:"type:varchar(20);unique_index"`
	Fullname string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(500)"`
}

// CreateUser func
func CreateUser(email, username, fullname, password string) (*User, error) {
	var user User
	user.ID = uuid.Must(uuid.NewV4()).String()
	user.Email = email
	user.Username = username
	user.Fullname = fullname
	pwd := getPwdBytes(password)
	user.Password = hashAndSalt(pwd)
	db.Create(&user)
	return &user, nil
}

func getPwdBytes(password string) []byte {
	var pwd string
	// Read the  password
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Fatal(err)
	}
	// Return the password as a byte slice
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
