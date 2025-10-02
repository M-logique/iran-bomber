//go:build windows
// +build windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)




func enableVirtualTerminalProcessing() {
kernel32 := syscall.NewLazyDLL("kernel32.dll")
    setConsoleMode := kernel32.NewProc("SetConsoleMode")
    getConsoleMode := kernel32.NewProc("GetConsoleMode")

    var mode uint32
    handle := syscall.Handle(os.Stdout.Fd())

    getConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
    mode |= 0x0004 // ENABLE_VIRTUAL_TERMINAL_PROCESSING
    setConsoleMode.Call(uintptr(handle), uintptr(mode))
}
