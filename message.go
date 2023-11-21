package main

import (
	"syscall"
	"unsafe"
)

func MessageSetup() Message {
    var msg Message
    msg.HWnd = syscall.Handle(0) // Replace with an appropriate value
    msg.Message = 0              // Replace with an appropriate value
    msg.WParam = uintptr(0)       // Replace with an appropriate value
    msg.LParam = uintptr(0)       // Replace with an appropriate value
    msg.Time = 0                 // Replace with an appropriate value
    msg.Pt = struct{ X, Y int32 }{0, 0} // Replace with appropriate values

    return msg
}

func MessageLoop(msg Message) {
	syscall.Syscall(getMessage.Addr(), 3, uintptr(unsafe.Pointer(&msg)), 0, 0)

	syscall.Syscall(translateMessage.Addr(), 1, uintptr(unsafe.Pointer(&msg)), 0, 0)
	syscall.Syscall(dispatchMessage.Addr(), 1, uintptr(unsafe.Pointer(&msg)), 0, 0)
}