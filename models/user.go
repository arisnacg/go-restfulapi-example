package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"arisnacg/go-restfulapi-example/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint
	Username  string    `gorm:"size:255;not null; unique;" json:"username"`
	Password  string    `gorm:"size:255;not null;" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func (user *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {
	// convert password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// remove space from username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Auth(username string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	generatedToken, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return generatedToken, nil

}
