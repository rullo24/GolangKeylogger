package main

import (
	"os/exec"
)

func main() {
	// Executing the command to kill the keylogger program
	cmd := exec.Command("taskkill", "/F", "/IM", "RC_keylogger.exe")
	cmd.Run()
}
