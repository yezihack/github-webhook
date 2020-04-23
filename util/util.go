package util

import (
	"fmt"
	"net/http"
	"os"
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

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// response html
func Response(w http.ResponseWriter, statusCode int, format string, msg ...interface{}) {
	w.WriteHeader(statusCode)
	_, _ = fmt.Fprintf(w, format, msg...)
}
