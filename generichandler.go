//go:build windows
// +build windows

package webview2

import (
	"github.com/google/uuid"
)

type _GenericHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type genericHandler struct {
	vtbl      *_GenericHandlerVtbl
	ref       uint32
	uuid      uuid.UUID
	name      string
	callback  func(args ...uintptr) uintptr
	onrelease func()
}

func _GenericHandlerInterface(this *genericHandler, refiid, object uintptr) uintptr {
	return 0
}

func _GenericHandlerAddRef(this *genericHandler) uintptr {
	this.ref++
	return uintptr(this.ref)
}

func _GenericHandlerRelease(this *genericHandler) uintptr {
	this.ref--
	if this.ref == 0 {
		if this.onrelease != nil {
			this.onrelease()
		}
		deleteRefGO(this.uuid)
	}
	return uintptr(this.ref)
}

func _GenericHandlerInvoke1(this *genericHandler, arg uintptr) uintptr {
	return this.callback(arg)
}

func _GenericHandlerInvoke2(this *genericHandler, arg1, arg2 uintptr) uintptr {
	return this.callback(arg1, arg2)
}

func _GenericHandlerInvoke3(this *genericHandler, arg1, arg2, arg3 uintptr) uintptr {
	return this.callback(arg1, arg2, arg3)
}

var _GenericHandlerFn = _GenericHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_GenericHandlerInterface),
		NewComProc(_GenericHandlerAddRef),
		NewComProc(_GenericHandlerRelease),
	},
	NewComProc(_GenericHandlerInvoke1),
}

func GenericHandler(name string, lenargs int, callback func(args ...uintptr) uintptr, onrelease func()) *genericHandler {

	hfunc := _GenericHandlerFn

	switch lenargs {
	case 1:
		hfunc.Invoke = NewComProc(_GenericHandlerInvoke1)
	case 2:
		hfunc.Invoke = NewComProc(_GenericHandlerInvoke2)
	case 3:
		hfunc.Invoke = NewComProc(_GenericHandlerInvoke3)

	}
	h := &genericHandler{
		name:      name,
		vtbl:      &hfunc,
		uuid:      uuid.New(),
		callback:  callback,
		onrelease: onrelease,
	}

	setRefGO(h.uuid, h)
	return h
}
