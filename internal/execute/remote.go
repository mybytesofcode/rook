package execute

import (
	"fmt"
	"os"
	"os/exec"
)

func RemoteExecute(path string, host string, port int, user string) error {
	var cmd *exec.Cmd
	cmd = exec.Command(
		"sh",
		"-c",
		fmt.Sprintf(
			"ssh -p %d %s@%s 'sudo bash -s' < %s",
			port,
			user,
			host,
			path,
		),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("remote execution failed: %w", err)
	}

	return err
}
