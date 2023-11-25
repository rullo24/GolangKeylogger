package main

import (
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
	"fmt"
)

// CallbackFunc is the type definition for the callback function
type CallbackFunc func(keys []uint32)

// Constants from Windows API
const (
	WM_KEYDOWN     = 0x0100
	WM_SYSKEYDOWN  = 0x0104
	WM_KEYUP       = 0x0101
	WM_SYSKEYUP    = 0x0105
	HC_ACTION      = 0
	WH_KEYBOARD_LL = 13
)

// KBDLLHOOKSTRUCT structure
type KBDLLHOOKSTRUCT struct {
	VkCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

var (
	user32DLL          = syscall.MustLoadDLL("user32.dll")
	setWindowsHookEx   = user32DLL.MustFindProc("SetWindowsHookExW")
	getMessage         = user32DLL.MustFindProc("GetMessageW")
	translateMessage   = user32DLL.MustFindProc("TranslateMessage")
	dispatchMessage    = user32DLL.MustFindProc("DispatchMessageW")
	unhookWindowsHookEx = user32DLL.MustFindProc("UnhookWindowsHookEx")

	modUser32           = windows.NewLazyDLL("user32.dll")
	procCallNextHookEx  = modUser32.NewProc("CallNextHookEx")
)

var (
	hookHandle uintptr
	currentlyPressedKeys []uint32
	callbackFunc CallbackFunc
	shiftPressed bool
)

// SetHook sets up the keyboard hook
func SetHook(callback CallbackFunc) {
	callbackFunc = callback

	// Set up the keyboard hook using WH_KEYBOARD_LL (Low-Level Keyboard Hook)
	ret, _, _ := syscall.Syscall6(setWindowsHookEx.Addr(), 5, uintptr(WH_KEYBOARD_LL), syscall.NewCallback(_keyboardCallback), 0, 0, 0, 0)
	hookHandle = ret
}

// Unhook removes the keyboard hook
func Unhook() {
	// Unhook the keyboard hook
	syscall.Syscall(unhookWindowsHookEx.Addr(), 1, uintptr(hookHandle), 0, 0)
}

// _keyboardCallback is the callback function for the keyboard hook
func _keyboardCallback(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode == HC_ACTION {
		kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))

		// Check if the key is pressed (wParam can be WM_KEYDOWN or WM_SYSKEYDOWN)
		if wParam == WM_KEYDOWN {
			// Turning the Shift Modifier On
			if kbdStruct.VkCode == 160 || kbdStruct.VkCode == 161 {
				shiftPressed = true
			}

			// Checking if the key that has just been pressed is a key on the keyboard (excluding nums and letters)
			printableKeys := keyModifierMap[int(kbdStruct.VkCode)]

			// Varying output key based on the shift key modifier
			if shiftPressed {
				fmt.Printf("%c", printableKeys.ModifiedKey)
			} else {
				fmt.Printf("%c", printableKeys.BaseKey)
			}

		} else if wParam == WM_KEYUP {
			// Setting the Shift Multiplier Off
			if kbdStruct.VkCode == 160 || kbdStruct.VkCode == 161 {
				shiftPressed = false
			}
		}

		// Running a function of your choosing
		callbackFunc(currentlyPressedKeys)
	}

	// Call the next hook in the chain using CallNextHookEx
	ret, _, _ := procCallNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
	return ret
}