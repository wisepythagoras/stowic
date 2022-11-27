package dom

func CreateTextComponent(text string) *Component {
	var textComponent Component = func(props Props, el *Element) (*Element, any) {
		el.setNativeType(TEXT)
		return nil, text
	}

	return &textComponent
}
