package stub_test

import (
	"testing"

	"github.com/ritabc/twg/suite/stub"
	"github.com/ritabc/twg/suite/suitetest"
)

// For this implementation (stub), we have only a few lines of code, because we wrote actual tests in suitetest package
func TestUserStore(t *testing.T) {
	us := &stub.UserStore{}
	suitetest.UserStore(t, us, nil, nil)
}

func TestUserStore_withStruct(t *testing.T) {
	us := stub.UserStore{}
	tests := suitetest.UserStoreSuite{
		UserStore: us,
	}
}
