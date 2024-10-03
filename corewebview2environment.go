//go:build windows
// +build windows

package webview2

import (
	"syscall"
	"unsafe"
)

func CreateCoreWebView2EnvironmentWithOptions(callback *genericHandler) {

	syscall.SyscallN(createCoreWebView2Environment.Addr(), uintptr(0), uintptr(0), uintptr(0), uintptr(unsafe.Pointer(callback)))
}
