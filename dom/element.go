package dom

import (
	"syscall/js"
)

type Element struct {
	Component   Component
	Props       *Props
	Children    []*Element
	TextContent *Text
	Styles      map[string]string
}

// handleChildren will loop through the children and append them to the native element
// that this element represents.
func (el *Element) handleChildren(doc *js.Value, native *js.Value) {
	for _, child := range el.Children {
		if text, ok := child.Component.(*Text); ok {
			native.Set("innerHTML", text.Contents)
		} else {
			native.Call("appendChild", *child.Render(doc))
		}
	}
}

// Render should handle the rendering of the element held by this instance.
func (el *Element) Render(doc *js.Value) *js.Value {
	if el.Component == nil {
		return nil
	}

	var props Props = nil

	if el.Props != nil {
		props = *el.Props
	}

	el.Component.Render(props)

	elType := "block"

	if nativeType, ok := el.Component.(NativeComponent); ok {
		elType = string(nativeType.Type())
	}

	native := doc.Call("createElement", elType)

	if el.Props != nil {
		if children, hasChildren := (*el.Props)["children"]; hasChildren {
			if elChildren, ok := children.([]*Element); ok {
				el.Children = elChildren
				el.handleChildren(doc, &native)
			} else if strChild, ok := children.(string); ok {
				native.Set("innerHTML", strChild)
			}
		} else if el.Children != nil && len(el.Children) > 0 {
			el.handleChildren(doc, &native)
		}
	} else if el.TextContent != nil {
		native.Set("innerHTML", el.TextContent.Contents)
	}

	if el.Styles != nil {
		style := native.Get("style")

		for rule, value := range el.Styles {
			style.Set(rule, value)
		}
	}

	return &native
}
