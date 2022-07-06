package security

import (
	"time"

	"github.com/markoxley/daggertech"
	"github.com/markoxley/daggertech/clause"
	"github.com/markoxley/daggertech/order"
)

// UserLocked returns true if the passed IP address is locked out
func UserLocked(un string) *time.Time {
	if m, ok := daggertech.First(&UserLock{}, &daggertech.Criteria{
		Where: clause.Equal("UserID", un).ToString(),
		Order: order.Desc("Release").ToString(),
	}); ok {
		if r, ok := m.(*UserLock); ok {
			if r.Release.Before(time.Now()) {
				RemoveUserLock(un)
				return nil
			}
			return &r.Release
		}
	}
	return nil
}

// RemoveUserLock removes a user lock
func RemoveUserLock(un string) {
	daggertech.RemoveMany("UserLock", &daggertech.Criteria{
		Where: clause.Equal("UserID", un).ToString(),
	})
}

// AddUserLock adds a lock for the username
func AddUserLock(un string, d time.Duration) time.Time {
	ClearUserAttempt(un)
	t := time.Now().Add(d)
	l := &UserLock{
		UserID:  un,
		Release: t,
	}
	daggertech.Save(l)
	return t
}
