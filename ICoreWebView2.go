//go:build windows
// +build windows

package webview2

import "unsafe"

type ICoreWebView2 struct {
	vtbl *iCoreWebView2Vtbl
}

func (e *ICoreWebView2) GetBrowserProcessID() uintptr {
	var ptr uintptr
	e.vtbl.GetBrowserProcessID.Call(uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(&ptr)))
	return ptr
}

func (e *ICoreWebView2) AddRef() uintptr {
	r, _, _ := e.vtbl.AddRef.Call(uintptr(unsafe.Pointer(e)))
	return r
}

func (e *ICoreWebView2) Release() uintptr {
	r, _, _ := e.vtbl.Release.Call(uintptr(unsafe.Pointer(e)))
	return r
}

func (e *ICoreWebView2) AddWindowCloseRequested(callback *genericHandler, token *uintptr) uintptr {
	r, _, _ := e.vtbl.AddWindowCloseRequested.Call(uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(token)))
	return r
}

func (e *ICoreWebView2) AddNavigationCompleted(callback *genericHandler, token *uintptr) uintptr {
	r, _, _ := e.vtbl.AddNavigationCompleted.Call(uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(callback)), uintptr(unsafe.Pointer(token)))
	return r
}

// ICoreWebView2

type iCoreWebView2Vtbl struct {
	_IUnknownVtbl
	GetSettings                            ComProc
	GetSource                              ComProc
	Navigate                               ComProc
	NavigateToString                       ComProc
	AddNavigationStarting                  ComProc
	RemoveNavigationStarting               ComProc
	AddContentLoading                      ComProc
	RemoveContentLoading                   ComProc
	AddSourceChanged                       ComProc
	RemoveSourceChanged                    ComProc
	AddHistoryChanged                      ComProc
	RemoveHistoryChanged                   ComProc
	AddNavigationCompleted                 ComProc
	RemoveNavigationCompleted              ComProc
	AddFrameNavigationStarting             ComProc
	RemoveFrameNavigationStarting          ComProc
	AddFrameNavigationCompleted            ComProc
	RemoveFrameNavigationCompleted         ComProc
	AddScriptDialogOpening                 ComProc
	RemoveScriptDialogOpening              ComProc
	AddPermissionRequested                 ComProc
	RemovePermissionRequested              ComProc
	AddProcessFailed                       ComProc
	RemoveProcessFailed                    ComProc
	AddScriptToExecuteOnDocumentCreated    ComProc
	RemoveScriptToExecuteOnDocumentCreated ComProc
	ExecuteScript                          ComProc
	CapturePreview                         ComProc
	Reload                                 ComProc
	PostWebMessageAsJSON                   ComProc
	PostWebMessageAsString                 ComProc
	AddWebMessageReceived                  ComProc
	RemoveWebMessageReceived               ComProc
	CallDevToolsProtocolMethod             ComProc
	GetBrowserProcessID                    ComProc
	GetCanGoBack                           ComProc
	GetCanGoForward                        ComProc
	GoBack                                 ComProc
	GoForward                              ComProc
	GetDevToolsProtocolEventReceiver       ComProc
	Stop                                   ComProc
	AddNewWindowRequested                  ComProc
	RemoveNewWindowRequested               ComProc
	AddDocumentTitleChanged                ComProc
	RemoveDocumentTitleChanged             ComProc
	GetDocumentTitle                       ComProc
	AddHostObjectToScript                  ComProc
	RemoveHostObjectFromScript             ComProc
	OpenDevToolsWindow                     ComProc
	AddContainsFullScreenElementChanged    ComProc
	RemoveContainsFullScreenElementChanged ComProc
	GetContainsFullScreenElement           ComProc
	AddWebResourceRequested                ComProc
	RemoveWebResourceRequested             ComProc
	AddWebResourceRequestedFilter          ComProc
	RemoveWebResourceRequestedFilter       ComProc
	AddWindowCloseRequested                ComProc
	RemoveWindowCloseRequested             ComProc

	//ICoreWebView2_2
	AddWebResourceResponseReceived    ComProc
	RemoveWebResourceResponseReceived ComProc
	NavigateWithWebResourceRequest    ComProc
	AddDOMContentLoaded               ComProc
	Remove_DOMContentLoaded           ComProc
	GetCookieManager                  ComProc
	GetEnvironment                    ComProc

	//ICoreWebView2_3
	ClearVirtualHostNameToFolderMapping ComProc
	GetIsSuspended                      ComProc
	Resume                              ComProc
	SetVirtualHostNameToFolderMapping   ComProc
	TrySuspend                          ComProc

	//ICoreWebView2_4
	_ [4]ComProc

	//ICoreWebView2_5
	_ [2]ComProc

	//ICoreWebView2_6
	_ ComProc
	//ICoreWebView2_7
	_ ComProc
	//ICoreWebView2_8
	_ [7]ComProc
	//ICoreWebView2_9
	_ [9]ComProc
	//ICoreWebView2_10
	_ [2]ComProc
	//ICoreWebView2_11
	_ [3]ComProc
	//ICoreWebView2_12
	_ [3]ComProc
	//ICoreWebView2_13
	GetProfile ComProc
}
