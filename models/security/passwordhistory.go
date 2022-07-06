package security

import "github.com/markoxley/daggertech"

// PasswordHistory is the list of passwords previously used
type PasswordHistory struct {
	daggertech.Model
	UserID      string `daggertech:"type:uuid,key:true"`
	OldPassword string `daggertech:"size:128"`
}
