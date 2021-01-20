package timing

import (
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	now := time.Now()
	timeNow = func() time.Time {
		return now
	}
	user := User{}
	SaveUser(&user)
	// if user was created more than a second ago, error
	if user.UpdatedAt != now {
		t.Errorf("user.UpdatedAt = %v, want ~%v", user.UpdatedAt, now)
	}
}

func TestUserSaver_Save(t *testing.T) {
	now := time.Now()
	us := UserSaver{
		now: func() time.Time {
			return now
		},
	}
	user := User{}
	us.Save(&user)
	if time.Now().Sub(user.UpdatedAt) > 1*time.Second {
		t.Errorf("user.UpdatedAt = %v, want ~%v", user.UpdatedAt, time.Now())
	}
}
