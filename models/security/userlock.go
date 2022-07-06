package security

import (
	"time"

	"github.com/markoxley/daggertech"
)

// UserLock stores the locked out IP addresses
type UserLock struct {
	daggertech.Model
	UserID  string    `daggertech:"size:128,key:true"`
	Release time.Time `daggertech:""`
}
