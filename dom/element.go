package dom

import (
	"syscall/js"
)

type Element struct {
	Component     *Component
	Props         *Props
	Children      []*Element
	TextContent   string
	Styles        *Styles
	EventHandlers *NativeEventHandlerMap
	nativeType    NativeComponentType
	stateHandlers map[any]any
}

// handleChildren will loop through the children and append them to the native element
// that this element represents.
func (el *Element) handleChildren(doc *js.Value, native *js.Value) {
	for _, child := range el.Children {
		res := child.Render(doc, native)

		if res != nil {
			native.Call("appendChild", *res)
		}
	}
}

func (el *Element) getSetStateHandler(valueAddr any, state any) any {
	if el.stateHandlers == nil {
		el.stateHandlers = make(map[any]any)
	}

	if existingState, hasState := el.stateHandlers[valueAddr]; hasState {
		return existingState
	}

	el.stateHandlers[valueAddr] = state

	return state
}

func (el *Element) setNativeType(nativeType NativeComponentType) {
	el.nativeType = nativeType
}

func (el *Element) GetNativeType() NativeComponentType {
	return el.nativeType
}

// Render should handle the rendering of the element held by this instance.
func (el *Element) Render(doc *js.Value, parent *js.Value) *js.Value {
	if el.Component == nil {
		return nil
	}

	var props Props = nil

	if el.Props != nil {
		props = *el.Props
	}

	childElement, other := (*el.Component)(props, el)

	elType := "block"

	if len(el.nativeType) > 0 {
		elType = string(el.nativeType)
	} else if elType == "block" {
		// Handles custom components.
		return childElement.Render(doc, parent)
	}

	// If we encounter a text component, then we only need to append the text to the
	// parent component.
	if el.nativeType == TEXT && parent != nil {
		parent.Set("innerHTML", other.(string))
		return nil
	}

	native := doc.Call("createElement", elType)

	if el.EventHandlers != nil {
		for eventType, handler := range *el.EventHandlers {
			native.Call("addEventListener", string(eventType), js.FuncOf(*handler))
		}
	}

	hasAddedText := false

	if len(el.TextContent) > 0 {
		// The text content takes precedence over everything else.
		native.Set("innerHTML", el.TextContent)
		hasAddedText = true
	}

	if el.Props != nil {
		if children, hasChildren := (*el.Props)["children"]; hasChildren {
			// If the children are an array of elements, then we need to go through all
			// of the children and handle them separately. Otherwise, if it's of string
			// type, simply append that to the innerHTML.
			if elChildren, ok := children.([]*Element); ok {
				el.Children = elChildren
				el.handleChildren(doc, &native)
			} else if strChild, ok := children.(string); ok && !hasAddedText {
				native.Set("innerHTML", strChild)
			}
		} else if el.Children != nil && len(el.Children) > 0 {
			el.handleChildren(doc, &native)
		}
	}

	if el.Styles != nil {
		style := native.Get("style")

		for rule, value := range *el.Styles {
			style.Set(rule, value)
		}
	}

	return &native
}

func CreateBareElement(c *Component, s *Styles, ch ...*Element) *Element {
	if c == nil {
		return nil
	}

	return &Element{
		Component:     c,
		Styles:        s,
		EventHandlers: nil,
		Props:         &Props{},
		Children:      ch,
	}
}

func attachToElement(e *Element, item interface{}) {
	if styles, isStyles := item.(*Styles); isStyles {
		e.Styles = styles
	} else if children, isChildren := item.([]*Element); isChildren {
		e.Children = children
	} else if props, isProps := item.(*Props); isProps {
		e.Props = props
	} else if eventHandlers, isEH := item.(*NativeEventHandlerMap); isEH {
		e.EventHandlers = eventHandlers
	} else if child, isChild := item.(*Element); isChild {
		e.Children = append(e.Children, child)
	} else if str, isStr := item.(string); isStr {
		e.TextContent = str
	}
}

// CreateElement will compose the element object and return it so that the calling code
// remains neat and standardized.
func CreateElement(c *Component, items ...interface{}) *Element {
	if c == nil {
		return nil
	}

	element := &Element{
		Component: c,
		Props:     &Props{},
		Children:  []*Element{},
	}

	for _, item := range items {
		attachToElement(element, item)
	}

	return element
}
