package dom

import (
	"syscall/js"
)

type Props map[string]interface{}

type Component func(Props, *Element) (*Element, any)

type Styles map[string]string

type NativeEventHandlerType string

// Documentation on these: https://developer.mozilla.org/en-US/docs/Web/API/Element
const (
	ONCLICK      NativeEventHandlerType = "click"
	ONBLUR       NativeEventHandlerType = "blur"
	ONFOCUS      NativeEventHandlerType = "focus"
	ONFOCUSIN    NativeEventHandlerType = "focusin"
	ONFOCUSOUT   NativeEventHandlerType = "focusout"
	ONKEYDOWN    NativeEventHandlerType = "keydown"
	ONKEYPRESS   NativeEventHandlerType = "keypress"
	ONKEYUP      NativeEventHandlerType = "keyup"
	ONMOUSEENTER NativeEventHandlerType = "mouseenter"
	ONMOUSEDOWN  NativeEventHandlerType = "mousedown"
	ONMOUSELEAVE NativeEventHandlerType = "mouseleave"
	ONMOUSEMOVE  NativeEventHandlerType = "mousemove"
	ONMOUSEOVER  NativeEventHandlerType = "mouseover"
	ONMOUSEUP    NativeEventHandlerType = "mouseup"
)

type NativeEventHandler func(this js.Value, args []js.Value) any

type NativeEventHandlerMap map[NativeEventHandlerType]*NativeEventHandler
