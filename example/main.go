//go:build windows
// +build windows

package main

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	webview2 "github.com/realPy/walk-webview2"
)

func main() {
	var mw *walk.MainWindow

	var ww *webview2.WebView2Widget

	m := declarative.MainWindow{
		AssignTo: &mw,
		Title:    "Hello google example",
		MinSize:  declarative.Size{Width: 600, Height: 400},
		Layout:   declarative.Grid{MarginsZero: true, SpacingZero: true},
		Children: []declarative.Widget{
			webview2.WebView2Container{
				AssignTo:   &ww,
				InitialURL: "https://realpy.github.io/hogosuru-brotlidecoder/index.html",
			},
		},
	}

	m.Create()
	mw.Run()
}
