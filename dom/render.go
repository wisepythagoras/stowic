//go:build js && wasm
// +build js,wasm

package dom

import (
	"syscall/js"
)

func Render(el *Element, id string) {
	doc := js.Global().Get("document")
	tree := el.Render(&doc)

	target := doc.Call("getElementById", id)
	target.Set("innerHTML", "")

	if tree != nil {
		target.Call("appendChild", *tree)
	}
}
