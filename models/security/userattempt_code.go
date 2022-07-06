package security

import (
	"time"

	"github.com/markoxley/daggertech"
	"github.com/markoxley/daggertech/clause"
	"github.com/markoxley/swordtech"
)

// AddUserAttempt adds a new IP attempt record
func AddUserAttempt(un string) *time.Time {
	if t := UserLocked(un); t != nil {
		return t
	}
	attemptMinutes := swordtech.GetParameter("Security", "UserThreshhold", "Number of minutes to check user attempts", "30").Int()
	attemptsAllowed := swordtech.GetParameter("Security", "UserAttempts", "Number of attempts allowed for a username", "3").Int()
	ipLockDuration := swordtech.GetParameter("Security", "UserLockDuration", "Minutes to lock username", "20").Int()
	if attemptMinutes == 0 || attemptsAllowed == 0 || ipLockDuration == 0 {
		return nil
	}
	epoch := time.Now().Add(time.Minute * time.Duration(-attemptMinutes))
	attempt := &UserAttempt{
		UserID: un,
	}
	daggertech.Save(attempt)

	if c := daggertech.Count("UserAttempt", &daggertech.Criteria{
		Where: clause.Equal("UserID", un).AndNotLess("CreateDate", epoch).ToString(),
	}); int64(c) > attemptsAllowed {
		t := AddUserLock(un, time.Minute*time.Duration(ipLockDuration))
		return &t
	}
	return nil
}

// ClearUserAttempt clears all locks and attempts for the specified username
func ClearUserAttempt(un string) {
	RemoveUserLock(un)
	daggertech.RemoveMany("UserAttempt", &daggertech.Criteria{
		Where: clause.Equal("Username", un).ToString(),
	})
}
