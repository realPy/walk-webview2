//go:build windows
// +build windows

package webview2

import (
	"unsafe"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

const webview2Class = "WebView2 Class"

var createCoreWebView2Environment *windows.LazyProc

func init() {
	web2 := windows.NewLazyDLL("WebView2Loader.dll")

	createCoreWebView2Environment = web2.NewProc("CreateCoreWebView2EnvironmentWithOptions")

	walk.AppendToWalkInit(func() {
		walk.MustRegisterWindowClass(webview2Class)
	})
}

type WebView2Widget struct {
	walk.WidgetBase
	wview *WebView2
}

type WebView2LayoutItem struct {
	walk.LayoutItemBase
}

func NewWebView2Widget(parent walk.Container, webviewinterface WebView2Interface) (*WebView2Widget, error) {
	w := new(WebView2Widget)

	if err := walk.InitWidget(
		w,
		parent,
		webview2Class,
		win.WS_VISIBLE,
		0); err != nil {

		return nil, err
	}

	w.wview = NewWebView2(w.Handle(), webviewinterface, true)
	CreateCoreWebView2EnvironmentWithOptions(GenericHandler("EnvironmentCompleted", 2, w.wview.EnvironmentCompleted, nil))

	return w, nil
}

func (w *WebView2Widget) CreateLayoutItem(ctx *walk.LayoutContext) walk.LayoutItem {

	return walk.NewGreedyLayoutItem()
}
func (w *WebView2Widget) onResize(x, y int32) {
	if w.wview != nil {
		if w.wview.controller != nil {
			w.wview.controller.PutBounds(win.RECT{0, 0, x, y})
		}
	}

}

func (w *WebView2Widget) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {

	switch msg {
	case win.WM_WINDOWPOSCHANGED:
		wp := (*win.WINDOWPOS)(unsafe.Pointer(lParam))

		if wp.Flags&win.SWP_NOSIZE != 0 {
			break
		}

		w.onResize(int32(wp.Cx), int32(wp.Cy))

	}

	return w.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
}

func (w *WebView2Widget) Dispose() {
	w.wview.Close()
	w.WidgetBase.Dispose()
}

func (li *WebView2LayoutItem) LayoutFlags() walk.LayoutFlags {
	return 0
}

type WebView2Container struct {
	AssignTo   **WebView2Widget
	InitialURL string
}

func (c WebView2Container) Create(builder *declarative.Builder) error {

	w, _ := NewWebView2Widget(builder.Parent(), &c)

	if c.AssignTo != nil {
		*c.AssignTo = w
	}

	return builder.InitWidget(c, w, func() error {
		return nil
	})
}

func (c *WebView2Container) GetInitialURL() string {
	return c.InitialURL
}
