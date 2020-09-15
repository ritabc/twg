package gen

import (
	"fmt"
	"sync"
	"testing"
)

var m sync.Mutex
var count = 0

// Email generates a unique email address every time it's called. It's intended to be used for creating new user accounts without worrying about an email address already being used.
// Note: This does not check the DB to verify that th email is not taken, so if you are generating email addresses another way it is possible to have collisions

func Email(t *testing.T) string {
	m.Lock()
	defer m.Unlock()
	t.Log("log seed here if generating random data for bug-finding purposes")
	ret := fmt.Sprintf("user-%d@example.com", count)
	count++
	return ret
}
