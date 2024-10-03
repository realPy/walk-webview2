


# walk-webview2

This package provides walk (https://github.com/lxn/walk) widget for using the Microsoft Edge WebView2 . It is inspired from [go-webview2]
(https://github.com/jchv/go-webview2) and provides a compatible API.

The internal webview in walk cant load wasm files. Walk-webview2 bring support with a recent edge webview.



## Example with hogosuru framework

If you are using Windows 10+, the WebView2 runtime should already be installed. If you don't have it installed, you can download and install a copy from Microsoft's website:

[WebView2 runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)

After that, you should be able to run go-webview2 directly:

Build the example
```
GOOS="windows" go build  -ldflags="-H windowsgui" -o webview-walk.exe example/main.go
```





