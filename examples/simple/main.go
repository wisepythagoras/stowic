package main

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/wisepythagoras/stowic/dom"
)

var myComponent dom.Component = func(p dom.Props, e *dom.Element) (*dom.Element, any) {
	var onButtonClick dom.NativeEventHandler = func(this js.Value, args []js.Value) any {
		fmt.Println("Button clicked", time.Now().GoString())
		return nil
	}

	return &dom.Element{
		Component: &dom.P,
		Props:     &dom.Props{},
		Children: []*dom.Element{
			{
				Component:   &dom.Button,
				TextContent: "Click me",
				EventHandlers: dom.NativeEventHandlerMap{
					dom.ONCLICK: &onButtonClick,
				},
			},
			{
				Component:   &dom.Span,
				TextContent: "This is another span",
			},
		},
		Styles:      dom.Styles{"color": "blue"},
		TextContent: "This is a custom component",
	}, nil
}

func main() {
	lock := make(chan int, 1)
	divProps := make(dom.Props)
	divProps["children"] = []*dom.Element{
		{
			Component: dom.CreateTextComponent("Hello, world!"),
		},
	}

	rootProps := make(dom.Props)
	rootProps["children"] = []*dom.Element{
		{
			Component: &dom.Div,
			Props:     &divProps,
		},
		{
			Component: &dom.Span,
			// Props: &dom.Props{
			// 	"children": "This is a span",
			// },
			TextContent: "This is a span!!!",
			Styles: dom.Styles{
				"fontSize": "21px",
				"color":    "red",
			},
		},
		{
			Component: &dom.H1,
			Props: &dom.Props{
				"children": "This is a header",
			},
		},
		{
			Component: &dom.Div,
			Styles:    dom.Styles{"padding": "10px"},
			Props: &dom.Props{
				"children": []*dom.Element{
					{
						Component: &myComponent,
					},
					{
						Component: &dom.Div,
						Props: &dom.Props{
							"children": []*dom.Element{
								{
									Component: &dom.Div,
									Props: &dom.Props{
										"children": "Nesting test",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	element := dom.Element{
		Component: &dom.Div,
		Props:     &rootProps,
	}

	dom.Render(&element, "hello")
	<-lock
}
