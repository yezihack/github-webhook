package util

import (
	"os/exec"
)

// call script bash function
func CallScript(path string) (string, error) {
	cmd := exec.Command("/bin/bash", path)
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
