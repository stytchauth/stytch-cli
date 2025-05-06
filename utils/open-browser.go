package utils

import (
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	default: // Linux and others
		cmd = "xdg-open"
	}
	if cmd != "" {
		args = append([]string{url}, args...)
		exec.Command(cmd, args...).Start()
	}
}