//
// Package goui provides a cross platform GUI solution for Go developers.
//
// It uses Cocoa/WebKit for macOS, MSHTML for Windows and Gtk/WebKit for Linux.
//
// It provides two way bindings between Go and Javascript.
//
//
package goui

import "C"
import (
	"encoding/json"
	"path/filepath"
	"strings"
)

/*type iWindow interface {
	create(Settings)
	exit()
	activate()
	invokeJS(string)
}*/

// Settings is to configure the window's appearance
type Settings struct {
	Title     string
	Url       string
	Left      int
	Top       int
	Width     int
	Height    int
	Resizable bool
	Debug     bool
}

//as goui designed to support only single-page application, it is reasonable to hold a window globally
var w *window

// Create is to create a native window with a webview
//
func Create(settings Settings) (err error) {
	return CreateWithMenu(settings, nil)
}

func CreateWithMenu(settings Settings, menuDefs []MenuDef) (err error) {
	if !strings.HasSuffix(settings.Url, "http") {
		settings.Url, err = filepath.Abs(settings.Url)
		settings.Url = "file://" + settings.Url
	}

	if err != nil {
		return
	}

	w.create(settings, menuDefs)
	defer w.exit()

	return

}

// Service is to add a backend service for frontend to invoke.
// params:
//	url - the url act as an unique identifier of a service, for example, "user/login", "blog/get/:id".
//	handler - the function that handle the client request.
func Service(url string, handler func(*Context)) {
	route := new(route)
	route.handler = handler
	parseRoute(url, route)
}

type JSServiceOptions struct {
	Url     string      `json:"url"`
	Data    interface{} `json:"data"`
	Success string      `json:"success"`
	Error   string      `json:"error"`
}

// RequestJSService is to send a request to the front end
func RequestJSService(options JSServiceOptions) (err error) {
	ops, err := json.Marshal(options)
	if err != nil {
		return
	}

	w.invokeJS("goui.handleRequest(" + string(ops) + ")")
	return
}

func ActivateWindow() {
	//window.activate()
}

//InvokeJavascriptFunc is for the backend to invoke frontend javascript directly.
//params:
//	name - javascript function name
//	params - the parameters
/*func InvokeJavascriptFunc(name string, params ...interface{})  {
	js := fiputil.MkString(params,name + "(",",",")")
	worker.invokeJS(js)
}
*/
func OpenDevTools() {

}
