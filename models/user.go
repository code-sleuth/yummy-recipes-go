package models

import (
	"errors"
	"log"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db = dbConnect()

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

	// save new user to db
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers func
func GetUsers() (*[]User, error) {
	var userList []User

	// get all users from database
	if err := db.Find(&userList).Error; err != nil {
		return nil, err
	}

	if len(userList) > 0 {
		return &userList, nil
	}
	return nil, errors.New("no users in db")
}

// GetUser func
func GetUser(id string) (*User, error) {
	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if user.ID == id {
		return &user, nil
	}

	return nil, errors.New("cannot find user with given id")
}

// UpdateUser func
func UpdateUser(id, email, username, fullname string) (*User, error) {
	user, err := GetUser(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	user.Username = username
	user.Email = email
	user.Fullname = fullname

	if err := db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser func
func DeleteUser(id string) (string, error) {
	user, err := GetUser(id)
	if err != nil {
		return "", err
	}

	if err := db.Delete(&user).Error; err != nil {
		return "", err
	}

	return "delete successful", nil
}

func getPwdBytes(password string) []byte {
	// Return the password as a byte slice
	return []byte(password)
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
