package material

import (
	"github.com/fsuhrau/gjs-material/bind"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Notification struct {
	vecty.Core
	Title      bind.Bind
	Message    bind.Bind
	Icon       vecty.ComponentOrHTML
	Color      Color
	Size       ButtonSize
	DataTarget string
	DataToggle string
	Rounded    bool
	Link       bool
	Disabled   bool
	Mini       bool
	Markup     vecty.MarkupList
}

func (c *Notification) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			c,
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("container"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("alert-icon"),
				),
				c.Icon,
				elem.Italic(
					vecty.Markup(
						vecty.Class("material-icons"),
					),
					vecty.Text("info_outline"),
				),
			),
			elem.Button(
				vecty.Markup(
					prop.Type(prop.TypeButton),
					vecty.Class("close"),
					vecty.Data("dismiss", "alert"),
					vecty.Attribute("aria-label", "Close"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Attribute("aria-hidden", "true"),
					),
					elem.Italic(
						vecty.Markup(
							vecty.Class("material-icons"),
						),
						vecty.Text("clear"),
					),
				),
			),
			elem.Bold(
				vecty.Text(c.Title.Get()+":"),
			),
			vecty.Text(c.Message.Get()),
		),
	)
}

func (c *Notification) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("alert"),
		MarkupColor("alert", c.Color),
	).Apply(h)
}
