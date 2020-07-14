package sleep

import (
	"testing"
	"time"
)

// Allows us to inspect what happens when `go test` is run
// go tool looks at source file, builds new binary out of all test source files
// stores binary in temporary folder
// see by $ ps -u rbennett | grep go
// binary sleep.test can be run
// running `go build` doesn't include any files ending in '_test.go'
func TestTmpExecutable(t *testing.T) {
	time.Sleep(time.Minute)
}
