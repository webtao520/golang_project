package golog

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	kSTD_OUTPUT_HANDLE                  uintptr = 4294967285
	kENABLE_VIRTUAL_TERMINAL_PROCESSING         = 0x0004
	kINVALID_HANDLE_VALUE               uintptr = ^uintptr(0)
)

func init() {
	EnableVT100()
}

func EnableVT100() error {
	kernel32 := mustLoadLibrary("kernel32.dll")
	getStdHandle := mustGetProcAddress(kernel32, "GetStdHandle")
	getConsoleMode := mustGetProcAddress(kernel32, "GetConsoleMode")
	setConsoleMode := mustGetProcAddress(kernel32, "SetConsoleMode")

	h, _, err := syscall.Syscall(getStdHandle, 1, kSTD_OUTPUT_HANDLE, 0, 0)
	if h == kINVALID_HANDLE_VALUE {
		return errors.New(fmt.Sprintf("enable VT100 support failed: %s", err.Error()))
	}

	mode := uint32(0)
	r, _, err := syscall.Syscall(getConsoleMode, 2, h, uintptr(unsafe.Pointer(&mode)), 0)
	if r == 0 {
		return errors.New(fmt.Sprintf("enable VT100 support failed: %s", err.Error()))
	}

	mode = mode | kENABLE_VIRTUAL_TERMINAL_PROCESSING
	r, _, err = syscall.Syscall(setConsoleMode, 2, h, uintptr(mode), 0)
	if r == 0 {
		return errors.New(fmt.Sprintf("enable VT100 support failed: %s", err.Error()))
	}

	return nil
}

func mustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}
	return uintptr(lib)
}

func mustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}
	return uintptr(addr)
}
