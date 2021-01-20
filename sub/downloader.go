package sub

import "os/exec"

type Downloader struct {
	CloneCmd *exec.Cmd
}

func (d *Downloader) Download(repo, dst string) (string, error) {
	var cmd *exec.Cmd
	if d.CloneCmd == nil {
		cmd = exec.Command("git", "cloned", repo, dst)
	} else {
		name, args := d.CloneCmd.Args[0], append(d.CloneCmd.Args[1:], repo, dst)
		cmd = exec.Command(name, args...)
	}
	bytes, err := cmd.Output()
	if err != nil {
		return string(bytes), err
	}
	return string(bytes), nil
}
