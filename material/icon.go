package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// supported icons https://material.io/tools/icons/?style=baseline
type Icon struct {
	vecty.Core
	Color Color
	Name  string
}

// Render implements the vecty.Component interface.
func (c *Icon) Render() vecty.ComponentOrHTML {
	return elem.Italic(
		vecty.Markup(
			vecty.Class("material-icons"),
		),
		vecty.Text(c.Name),
	)
}

// func (i *Icon) Apply(h *vecty.HTML) {
// 	vecty.Markup(
// 		vecty.Class("icon"),
// 		MarkupColor("icon", i.Color),
// 	).Apply(h)
// }
