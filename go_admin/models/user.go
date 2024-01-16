package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}

func (user *User) GeneratePassword(passwordstring []byte) error {
	password, err := bcrypt.GenerateFromPassword([]byte(passwordstring), 14)
	user.Password = []byte(password)
	if err != nil{
		return err
	}
	return nil
}
