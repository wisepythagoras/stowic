package dom

type Props map[string]interface{}

type Component interface {
	Render(props Props) *Element
}
