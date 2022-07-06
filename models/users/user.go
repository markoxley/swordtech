package users

import (
	"time"

	"github.com/markoxley/daggertech"
)

// User is the struct for user details
type User struct {
	daggertech.Model
	UserID              string     `daggertech:"size:128,key:true"`
	FirstName           string     `daggertech:"size:128"`
	LastName            string     `daggertech:"size:128"`
	Email               string     `daggertech:"size:128"`
	Password            string     `daggertech:"size:128"`
	LastLogin           *time.Time `daggertech:"type:time"`
	PasswordExpiry      *time.Time `daggertech:"type:time"`
	PasswordForceChange bool       `daggertech:""`
	UnlockTime          *time.Time `daggertech:"type:time"`
}
