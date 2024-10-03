//go:build windows
// +build windows

package webview2

import (
	"log"
	"os"
	"time"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

type WebView2Interface interface {
	GetInitialURL() string
}

type WebView2 struct {
	webviewinterface WebView2Interface
	environment      *ICoreWebView2Environment
	controller       *ICoreWebView2Controller
	webview          *ICoreWebView2
	Hwnd             win.HWND
	datapath         string
	processID        uintptr
	cleancache       bool
}

func NewWebView2(h win.HWND, i WebView2Interface, cleancache bool) *WebView2 {
	var w WebView2
	w.Hwnd = h
	w.webviewinterface = i
	w.cleancache = cleancache

	return &w
}

func (e *WebView2) WindowCloseRequested(args ...uintptr) uintptr {

	return 0
}

func (e *WebView2) BrowserProcessExited(args ...uintptr) uintptr {

	return 0
}

func (e *WebView2) AddProcessInfosChanged(args ...uintptr) uintptr {

	env := (*ICoreWebView2Environment)(unsafe.Pointer(args[0]))
	var collections *ICoreWebView2ProcessInfoCollection
	env.GetProcessInfos(&collections)
	if collections != nil {
		if collections.GetCount() == 0 {

		}
	}

	return 0
}

func (e *WebView2) NavigationCompleted(args ...uintptr) uintptr {

	return 0
}

func (e *WebView2) EnvironmentCompleted(args ...uintptr) uintptr {
	res := int64(args[0])

	if int64(res) < 0 {
		log.Fatalf("Creating environment failed with %08x", res)
	}

	env := (*ICoreWebView2Environment)(unsafe.Pointer(args[1]))

	env.AddRef()
	e.environment = env

	env.CreateCoreWebView2Controller(e.Hwnd, GenericHandler("CreateCoreWebView2ControllerCompleted", 2, e.CreateCoreWebView2ControllerCompleted, nil))

	return 0
}

func (e *WebView2) CreateCoreWebView2ControllerCompleted(args ...uintptr) uintptr {
	res := int64(args[0])

	controller := (*ICoreWebView2Controller)(unsafe.Pointer(args[1]))

	if res < 0 {
		log.Fatalf("Creating controller failed with %08x", res)
	}

	controller.AddRef()
	e.controller = controller

	var token uintptr

	controller.GetCoreWebView2(&e.webview)
	e.webview.AddRef()

	e.webview.AddNavigationCompleted(GenericHandler("NavigationCompleted", 2, e.NavigationCompleted, nil), &token)

	e.webview.AddWindowCloseRequested(GenericHandler("WindowCloseRequested", 1, e.WindowCloseRequested, nil), &token)

	e.datapath = e.environment.GetUserDataFolder()
	e.processID = e.webview.GetBrowserProcessID()

	var bound win.RECT
	win.GetClientRect(e.Hwnd, &bound)

	e.controller.PutBounds(bound)
	if e.webviewinterface != nil {

		_, _, _ = e.webview.vtbl.Navigate.Call(
			uintptr(unsafe.Pointer(e.webview)),
			uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(e.webviewinterface.GetInitialURL()))),
		)
	}

	e.controller.PutIsVisible(true)

	return 0
}

func (e *WebView2) Close() {

	e.controller.Close()
	e.controller.Release()
	e.webview.Release()
	e.environment.Release()
	if e.cleancache {
		count := 0
		for {
			err := os.RemoveAll(e.datapath)
			if err == nil || count == 10 {
				break
			}
			time.Sleep(10 * time.Millisecond)
			count++
		}
	}
}
