package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type InfoArea struct {
	vecty.Core
	Icon       *Icon
	Horizontal bool
	Title      vecty.ComponentOrHTML
	Content    vecty.ComponentOrHTML
	Link       vecty.ComponentOrHTML
}

func (c *InfoArea) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			c,
		),
		vecty.If(c.Icon != nil,
			elem.Div(
				vecty.Markup(
					vecty.Class("icon"),
					MarkupColor("icon", c.Icon.Color),
				),
				c.Icon,
			),
		),
		vecty.If(c.Horizontal,
			elem.Div(
				vecty.Markup(
					vecty.Class("description"),
				),
				elem.Heading4(
					vecty.Markup(
						vecty.Class("info-title"),
					),
					c.Title,
				),
				c.Content,
				c.Link,
			),
		),
		vecty.If(!c.Horizontal,
			elem.Heading4(
				vecty.Markup(
					vecty.Class("info-title"),
				),
				c.Title,
			),
			c.Content,
			c.Link,
		),
	)
}

func (c *InfoArea) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("info"),
		vecty.MarkupIf(c.Horizontal, vecty.Class("info-horizontal")),
	).Apply(h)
}
