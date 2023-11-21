package main
import "syscall"

// Dynamic loading of user32.dll for Windows UI interactions.
var (
	user32 = syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

// isKeyPressed checks if the specified key is currently pressed
func isKeyPressed(key_address int) bool {
	result, _, _ := getAsyncKeyState.Call(uintptr(key_address))
	return result&0x8000 != 0
}