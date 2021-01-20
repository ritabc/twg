package sub_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/ritabc/twg/sub"
)

func TestDemo(t *testing.T) {
	fmt.Println(os.Args[0:])
	cmd := exec.Command(os.Args[0], "-test.run=Test_GitCloneSubprocess")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("err = %s\n", err)
	}
	fmt.Println(string(out))
}

// This is not a real test - it is used to simulate the git subprocess
func Test_GitCloneSubprocess(t *testing.T) {
	if os.Getenv("GO_RUNNING_SUBPROCESS") != "1" {
		fmt.Println("Skipping because not a subprocess")
		return
	}
	os.Mkdir("test-123", 777)
	os.Exit(1)
}

func TestDownloader_Download(t *testing.T) {
	var d sub.Downloader
	d.CloneCmd = exec.Command(os.Args[0], "-test.run=Test_GitCloneSubprocess")
	d.CloneCmd.Env = append(os.Environ(), "GO_RUNNING_SUBPROCESS=1")
	wantDir := "test-123"
	msg, err := d.Download("https://github.com/joncalhoun/form.git", wantDir)
	if err != nil {
		t.Errorf("Download() err = %s; want nil", err)
		t.Errorf("Download() output: %s", msg)
	}
	if _, err := os.Stat(wantDir); os.IsNotExist(err) {
		t.Errorf("Download() didn't create dir %s", wantDir)
		t.Errorf("Download() output: %s", msg)
	}
}
