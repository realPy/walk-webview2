//go:build windows
// +build windows

package webview2

import (
	"strings"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

// ICoreWebView2Environment

type iCoreWebView2EnvironmentVtbl struct {
	_IUnknownVtbl
	CreateCoreWebView2Controller     ComProc
	CreateWebResourceResponse        ComProc
	GetBrowserVersionString          ComProc
	AddNewBrowserVersionAvailable    ComProc
	RemoveNewBrowserVersionAvailable ComProc
	//ICoreWebView2Environment2
	CreateWebResourceRequest ComProc
	//ICoreWebView2Environment3
	CreateCoreWebView2CompositionController ComProc
	CreateCoreWebView2PointerInfo           ComProc
	//ICoreWebView2Environment4
	GetAutomationProviderForWindow ComProc
	//ICoreWebView2Environment5
	AddBrowserProcessExited    ComProc
	RemoveBrowserProcessExited ComProc
	//ICoreWebView2Environment6
	CreatePrintSettings ComProc
	//ICoreWebView2Environment7
	GetUserDataFolder ComProc
	//ICoreWebView2Environment8
	AddProcessInfosChanged    ComProc
	RemoveProcessInfosChanged ComProc
	GetProcessInfos           ComProc
}

type ICoreWebView2Environment struct {
	vtbl *iCoreWebView2EnvironmentVtbl
}

func (i *ICoreWebView2Environment) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment) GetUserDataFolder() string {
	var (
		LPWSTR   *uint16
		datapath string
	)

	i.vtbl.GetUserDataFolder.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&LPWSTR)))
	datapath = strings.Clone(utf16PtrToString(LPWSTR))
	windows.CoTaskMemFree(unsafe.Pointer(LPWSTR))
	return datapath
}

func (i *ICoreWebView2Environment) GetBrowserVersionString() string {
	var (
		LPWSTR  *uint16
		version string
	)

	i.vtbl.GetBrowserVersionString.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&LPWSTR)))
	version = strings.Clone(utf16PtrToString(LPWSTR))
	windows.CoTaskMemFree(unsafe.Pointer(LPWSTR))
	return version
}

func (i *ICoreWebView2Environment) AddBrowserProcessExited(callback *genericHandler, token *uintptr) uintptr {
	r, _, _ := i.vtbl.AddBrowserProcessExited.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(token)))
	return r
}

func (i *ICoreWebView2Environment) AddProcessInfosChanged(callback *genericHandler, token *uintptr) uintptr {
	r, _, _ := i.vtbl.AddProcessInfosChanged.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(token)))
	return r
}

func (i *ICoreWebView2Environment) CreateCoreWebView2Controller(hwnd win.HWND, callback *genericHandler) {
	i.vtbl.CreateCoreWebView2Controller.Call(uintptr(unsafe.Pointer(i)), uintptr(hwnd), uintptr(unsafe.Pointer(callback)))
}

func (i *ICoreWebView2Environment) GetProcessInfos(collections **ICoreWebView2ProcessInfoCollection) uintptr {
	r, _, _ := i.vtbl.GetProcessInfos.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(collections)))
	return r
}
