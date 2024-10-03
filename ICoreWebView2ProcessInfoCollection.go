//go:build windows
// +build windows

package webview2

import (
	"unsafe"
)

// ICoreWebView2ProcessInfoCollection

type iCoreWebView2ProcessInfoCollectionVtbl struct {
	_IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2ProcessInfoCollection struct {
	vtbl *iCoreWebView2ProcessInfoCollectionVtbl
}

func (i *ICoreWebView2ProcessInfoCollection) GetCount() uint {
	var count uint

	i.vtbl.GetCount.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&count)))

	return count
}
