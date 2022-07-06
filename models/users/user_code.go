package users

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/markoxley/daggertech"
	"github.com/markoxley/daggertech/clause"
	"github.com/markoxley/swordtech"
	"github.com/markoxley/swordtech/models/security"
)

// HashPassword creates the password hash
func (u *User) HashPassword(p string) string {
	// The Prophecies of Nostradamus Century II, Quatrain 27
	salt := `The divine word will be struck from the sky, 
	One who cannot proceed any further: 
	The secret closed up with the revelation, 
	Such that they will march over and ahead.`

	hsh := sha512.Sum512([]byte(*u.ID + p + salt))
	return hex.EncodeToString(hsh[:])
}

// UpdatePassword updates the user password
func (u *User) UpdatePassword(newPassword, oldPassword string) error {
	if u.IsNew() {
		return errors.New("User record does not have an ID")
	}
	op := u.HashPassword(oldPassword)
	if u.Password != op {
		return errors.New("Old password does not match")
	}
	np := u.HashPassword(newPassword)
	if daggertech.Count("PasswordHistory", &daggertech.Criteria{
		Where: clause.Equal("UserId", *u.ID).AndEqual("OldPassword", np).ToString(),
	}) > 0 {
		return errors.New("Password has already been used")
	}
	u.Password = np
	expTime := swordtech.GetParameter("User", "PasswordExpiry", "Number of months for password expiry", "6")
	if expTime.Int() == 0 {
		expTime.Set(6)
	}
	yr := int(expTime.Int() / 12)
	mnth := int(expTime.Int() % 12)
	exp := time.Now().AddDate(yr, mnth, 0)
	u.PasswordExpiry = &exp
	if !daggertech.Save(u) {
		return errors.New("Unable to update password due to an internal error")
	}
	ph := &security.PasswordHistory{
		UserID:      *u.ID,
		OldPassword: np,
	}
	daggertech.Save(ph)
	return nil
}

// FullName returns the fullname of the user
func (u *User) FullName() string {
	return strings.Trim(fmt.Sprintf("%s %s", strings.Trim(u.FirstName, " "), strings.Trim(u.LastName, " ")), " ")
}

// ForcePasswordExpiry sets the password forced flag
func (u *User) ForcePasswordExpiry() {
	u.PasswordForceChange = true
	daggertech.Save(u)
}

// Lock adds a lock to the user record
func (u *User) Lock(d time.Duration) {
	security.AddUserLock(u.UserID, d)
}

// GetUserByID returns the user record for the passed ID
func GetUserByID(userID *string) *User {
	if userID == nil {
		return nil
	}
	if result, ok := daggertech.First(&User{}, &daggertech.Criteria{
		Where: clause.Equal("ID", *userID).ToString(),
	}); ok {
		if user, ok := result.(*User); ok {
			return user
		}
	}
	return nil
}

// GetUserByUserID returns the user with the specified UserID
func GetUserByUserID(userID string) *User {
	if m, ok := daggertech.First(&User{}, &daggertech.Criteria{
		Where: clause.Equal("UserID", userID).ToString(),
	}); ok {
		if user, ok := m.(*User); ok {
			return user
		}
	}
	return nil
}

// ValidateUser returns the user record and response code when a user attempts to sign in
func ValidateUser(w http.ResponseWriter, r *http.Request, username, password string) (*User, UserState, *time.Time) {

	user, userState, timeStamp := TestUserLogin(username, password)
	switch userState {
	case UserOK:
		swordtech.SetSessionVar(r, w, "UserID", user.ID)
	case UserPasswordExpired:
		swordtech.SetSessionVar(r, w, "ExpiredUserID", user.ID)
	}
	return user, userState, timeStamp
}

// TestUserLogin tests the username and password combination entered
func TestUserLogin(username, password string) (*User, UserState, *time.Time) {
	// Check for user lock
	if t := security.UserLocked(username); t != nil {
		return nil, UserLocked, t
	}
	if model, ok := daggertech.First(&User{}, &daggertech.Criteria{
		Where: clause.Equal("UserID", username).ToString(),
	}); ok {
		if user, ok := model.(*User); ok {
			hsh := user.HashPassword(password)
			if hsh == user.Password {
				if (user.PasswordExpiry != nil && user.PasswordExpiry.After(time.Now())) || user.PasswordForceChange {
					return user, UserPasswordExpired, nil
				}
				now := time.Now()
				user.LastLogin = &now
				daggertech.Save(user)
				return user, UserOK, nil
			}
		}
	}

	// User record does not exist, create user attempt
	// If it returns a time, return the furthest in the future
	userT := security.AddUserAttempt(username)

	if userT == nil {
		return nil, UserInvalid, nil
	}

	return nil, UserLocked, userT
}

// CreateUser creates a new user record
func CreateUser(userID, firstname, lastname, email, password string) (*User, bool) {
	u := &User{
		UserID:              userID,
		FirstName:           firstname,
		LastName:            lastname,
		Email:               email,
		Password:            "",
		PasswordForceChange: true,
	}
	if ok := daggertech.Save(u); !ok {
		return nil, false
	}
	u.Password = u.HashPassword(password)
	daggertech.Save(u)
	return u, true
}
