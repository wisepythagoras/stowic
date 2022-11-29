package main

import (
	"syscall/js"

	"github.com/wisepythagoras/stowic/dom"
)

// myComponent Shows how to create and manage a custom component.
var myComponent dom.Component = func(p dom.Props, e *dom.Element) (*dom.Element, any) {
	var initialCount = 1
	count, setCount := dom.UseState(&initialCount, e)

	var onButtonClick dom.NativeEventHandler = func(this js.Value, args []js.Value) any {
		println("Button clicked", *count) // time.Now().GoString()
		setCount((*count) + 1)
		return nil
	}

	return dom.CreateElement(
		&dom.Div, &dom.Styles{"color": "blue", "border": "1px solid blue"}, &p,
		dom.CreateElement(&dom.P, "This is a custom component"),
		dom.CreateElement(&dom.Button, "Click me", &dom.NativeEventHandlerMap{
			dom.ONCLICK: &onButtonClick,
		}),
		dom.CreateElement(&dom.Span, "This is another span"),
	), nil
}

func main() {
	lock := make(chan int, 1)
	divProps := make(dom.Props)
	divProps["test"] = 123

	rootChildren := []*dom.Element{
		dom.CreateElement(&dom.Div, &divProps, dom.CreateElement(&dom.Span, "Hello, world!")),
		dom.CreateElement(&dom.Span, "This is a span!!!", &dom.Styles{
			"fontSize": "21px",
			"color":    "red",
		}, nil),
		dom.CreateElement(&dom.Span, "This is a span!!!", &dom.Styles{
			"fontSize":   "22px",
			"color":      "blue",
			"fontWeight": "bold",
		}, nil),
		dom.CreateElement(&dom.H1, "This is a header"),
		dom.CreateElement(
			&dom.Div, &dom.Styles{"padding": "10px"}, nil, &dom.Props{},
			dom.CreateElement(&myComponent, &dom.Props{"test": 123}),
			dom.CreateElement(&dom.Div, dom.CreateElement(&dom.Div, "This is a test")),
			dom.CreateElement(&dom.Div, dom.CreateElement(&dom.Div, "Nesting test")),
		),
	}

	// This is the equivalent to the initial render of a React app.
	dom.Render(dom.CreateElement(&dom.Div, rootChildren), "hello")

	<-lock
}
