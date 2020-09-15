package git

import (
	"os/exec"
	"strings"
)

var execCommand = exec.Command

// return current git version that the user is using
func Version() string {
	cmd := execCommand("git", "version")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	n := len("git version ")
	version := string(stdout[n:])
	return strings.TrimSpace(version)
}
