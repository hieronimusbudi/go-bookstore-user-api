package users

// import (
// 	"github.com/hieronimusbudi/go-bookstore-utils-new/rest_errors"
// )

import (
	"github.com/hieronimusbudi/go-bookstore-utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

//Validate func
func (user *User) Validate() errors.RestErr {
	return errors.NewBadRequestError("hello")
}
