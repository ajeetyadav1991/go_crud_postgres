package internal

import (
	//"time"
	"errors"
)
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)


// The User Object or struct is a representation of `users` table which we created in postgres database
type User struct {
	ID        int64     `json:"id"`
	//CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Location  string   `json:"location"`
	Age       int64    `json:"age"`
}

/* `User` starts with capital character. It means it is public or exported. It can be accessed by the other packages.*/
