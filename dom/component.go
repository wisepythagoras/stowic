package dom

type Props map[string]interface{}

type Component func(Props, *Element) (*Element, any)

type Styles map[string]string
