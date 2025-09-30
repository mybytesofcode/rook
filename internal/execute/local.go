package execute

import (
	"fmt"
	"os"
	"os/exec"
)

func LocalExecute(path string) error {
	var cmd *exec.Cmd
	cmd = exec.Command(
		"sh",
		"-c",
		fmt.Sprintf(
			"bash -s < %s",
			path,
		),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("local execution failed: %w", err)
	}

	return err
}
