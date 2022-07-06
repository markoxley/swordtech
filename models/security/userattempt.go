package security

import (
	"github.com/markoxley/daggertech"
)

// UserAttempt stores the attempts to sign in as a user
type UserAttempt struct {
	daggertech.Model
	UserID string `daggertech:"size:128,key:true"`
}
