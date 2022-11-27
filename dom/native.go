package dom

type NativeComponentType string

const (
	DIV      NativeComponentType = "div"
	SPAN     NativeComponentType = "span"
	HEADING1 NativeComponentType = "h1"
	TEXT     NativeComponentType = "text"
)

// Native div element.

var Div Component = func(props Props, el *Element) (*Element, any) {
	el.setNativeType(DIV)
	return nil, nil
}

// Native span element.

var Span Component = func(props Props, el *Element) (*Element, any) {
	el.setNativeType(SPAN)
	return nil, nil
}

// Native h1 element.

var H1 Component = func(props Props, el *Element) (*Element, any) {
	el.setNativeType(HEADING1)
	return nil, nil
}
