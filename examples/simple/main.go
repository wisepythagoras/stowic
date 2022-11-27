package main

import (
	"github.com/wisepythagoras/stowic/dom"
)

func divComponent() {
	//
}

func main() {
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
}
