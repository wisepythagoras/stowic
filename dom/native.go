package dom

type NativeComponentType string

const (
	DIV      NativeComponentType = "div"
	SPAN     NativeComponentType = "span"
	HEADING1 NativeComponentType = "h1"
)

type NativeComponent interface {
	Render(props Props) *Element
	Type() NativeComponentType
}

// Native div element.

type Div struct {
}

func (d *Div) Render(props Props) *Element {
	println("The div's render function was called")
	return nil
}

func (d *Div) Type() NativeComponentType {
	return DIV
}

// Native span element.

type Span struct {
}

func (s *Span) Render(props Props) *Element {
	println("The span's render function was called")
	return nil
}

func (s *Span) Type() NativeComponentType {
	return SPAN
}

// Native h1 element.

type H1 struct {
}

func (h *H1) Render(props Props) *Element {
	println("The h1's render function was called")
	return nil
}

func (h *H1) Type() NativeComponentType {
	return HEADING1
}
