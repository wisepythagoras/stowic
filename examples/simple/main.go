package main

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/wisepythagoras/stowic/dom"
)

// myComponent Shows how to create and manage a custom component.
var myComponent dom.Component = func(p dom.Props, e *dom.Element) (*dom.Element, any) {
	var onButtonClick dom.NativeEventHandler = func(this js.Value, args []js.Value) any {
		fmt.Println("Button clicked", time.Now().GoString())
		return nil
	}

	return dom.CreateElement(
		&dom.Div, &dom.Styles{"color": "blue", "border": "1px solid blue"}, nil, &p,
		dom.CreateTextElement(&dom.P, "This is a custom component", nil, nil),
		dom.CreateTextElement(&dom.Button, "Click me", nil, &dom.NativeEventHandlerMap{
			dom.ONCLICK: &onButtonClick,
		}),
		dom.CreateTextElement(&dom.Span, "This is another span", nil, nil),
	), nil
}

func main() {
	lock := make(chan int, 1)
	divProps := make(dom.Props)
	divProps["test"] = 123

	rootProps := make(dom.Props)
	rootProps["children"] = []*dom.Element{
		dom.CreateElement(&dom.Div, nil, nil, &divProps, dom.CreateTextElement(&dom.Span, "Hello, world!", nil, nil)),
		dom.CreateTextElement(&dom.Span, "This is a span!!!", &dom.Styles{
			"fontSize": "21px",
			"color":    "red",
		}, nil),
		dom.CreateTextElement(&dom.Span, "This is a span!!!", &dom.Styles{
			"fontSize":   "22px",
			"color":      "blue",
			"fontWeight": "bold",
		}, nil),
		dom.CreateTextElement(&dom.H1, "This is a header", nil, nil),
		dom.CreateElement(
			&dom.Div, &dom.Styles{"padding": "10px"}, nil, &dom.Props{},
			dom.CreateElement(&myComponent, nil, nil, &dom.Props{"test": 123}),
			dom.CreateBareElement(&dom.Div, nil, dom.CreateTextElement(&dom.Div, "This is a test", nil, nil)),
			dom.CreateBareElement(&dom.Div, nil, dom.CreateTextElement(&dom.Div, "Nesting test", nil, nil)),
		),
	}

	// This is the equivalent to the initial render of a React app.
	dom.Render(dom.CreateElement(&dom.Div, nil, nil, &rootProps), "hello")

	<-lock
}
