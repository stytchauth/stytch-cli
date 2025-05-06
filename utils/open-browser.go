package utils

import (
	"fmt"
	"os"
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
	args = append([]string{url}, args...)
	err := exec.Command(cmd, args...).Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
