package dom

type NativeComponentType string

const (
	DIV       NativeComponentType = "div"
	SPAN      NativeComponentType = "span"
	HEADING1  NativeComponentType = "h1"
	BUTTON    NativeComponentType = "button"
	PARAGRAPH NativeComponentType = "p"

	// This is a placeholder type.
	TEXT NativeComponentType = "text"
)

// Native div component.

var Div Component = func(props Props, el *Element) (*Element, any) {
	el.setNativeType(DIV)
	return nil, nil
}

// Native span component.

var Span Component = func(props Props, el *Element) (*Element, any) {
	el.setNativeType(SPAN)
	return nil, nil
}

// Native h1 component.

var H1 Component = func(props Props, el *Element) (*Element, any) {
	el.setNativeType(HEADING1)
	return nil, nil
}

// Native button component.

var Button Component = func(p Props, el *Element) (*Element, any) {
	el.setNativeType(BUTTON)
	return nil, nil
}

// Native paragraph component.

var P Component = func(p Props, el *Element) (*Element, any) {
	el.setNativeType(PARAGRAPH)
	return nil, nil
}
