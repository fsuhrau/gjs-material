package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Badge struct {
	vecty.Core
	Text  string
	Color Color
	Pill  bool
	HRef  string
}

func (c *Badge) Render() vecty.ComponentOrHTML {
	var component vecty.ComponentOrHTML
	if c.HRef != "" {
		component = elem.Anchor(
			vecty.Markup(
				c,
				prop.Href(c.HRef),
			),
			vecty.Text(c.Text),
		)
	} else {
		component = elem.Span(
			vecty.Markup(
				c,
			),
			vecty.Text(c.Text),
		)
	}

	return component
}

func (c *Badge) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("badge"),
		vecty.MarkupIf(c.Pill, vecty.Class("badge-pill")),
		MarkupColor("badge", c.Color),
	).Apply(h)
}
