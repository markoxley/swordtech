package users

// UserState is the result of the user logging in
type UserState int

const (
	// UserOK indicates that the user has successfully logged in
	UserOK UserState = iota
	// UserInvalid indicates that the user has not successfully logged in
	UserInvalid
	// UserPasswordExpired indicates that the user has successfully logged in, but
	// they need to change the password
	UserPasswordExpired
	// UserLocked indicates that the user account has been locked
	UserLocked
)
