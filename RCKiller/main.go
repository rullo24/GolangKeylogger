package main

import (
	"os/exec"
)

func main() {
	// Executing the command to kill the keylogger program
	cmd := exec.Command("taskkill", "/F", "/IM", "keylog.exe")
	cmd.Run()
}
