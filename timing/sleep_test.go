package timing_test

import (
	"testing"

	"github.com/ritabc/twg/timing"
)

func TestPollUntil(t *testing.T) {
	fn := func() bool {
		return false
	}
	err := timing.PollUntil(fn, 2)
	if err != timing.ErrExceededMaxTries {
		t.Errorf("PollUntil() err = %s, want %s", err, timing.ErrExceededMaxTries)
	}
}
