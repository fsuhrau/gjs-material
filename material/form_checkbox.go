package material

import (
	"github.com/fsuhrau/gjs-material/bind"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FormField is a vecty-material formfield component.
type Checkbox struct {
	vecty.Core
	Root     vecty.MarkupOrChild
	Input    vecty.ComponentOrHTML
	Label    string
	Name     string
	Checked  bool
	Disabled bool
	Inline   bool
	Radio    bool
	Val      string
	Value    bind.Bind
}

// Render implements the vecty.Component interface.
func (c *Checkbox) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			c,
		),
		elem.Label(
			vecty.Markup(
				vecty.Class("form-check-label"),
			),
			elem.Input(
				vecty.Markup(
					vecty.Class("form-check-input"),
					vecty.MarkupIf(c.Radio, prop.Type(prop.TypeRadio)),
					vecty.MarkupIf(!c.Radio, prop.Type(prop.TypeCheckbox)),
					prop.Value(c.Value.Get()),
					vecty.Property("disabled", c.Disabled),
					vecty.MarkupIf(c.Name != "", vecty.Attribute("name", c.Name)),
				),
			),
			vecty.Text(c.Label),
			elem.Span(
				vecty.Markup(
					vecty.MarkupIf(!c.Radio, vecty.Class("form-check-sign")),
					vecty.MarkupIf(c.Radio, vecty.Class("circle")),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("check"),
					),
				),
			),
		),
	)
}

func (c *Checkbox) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("form-check"),
		vecty.MarkupIf(c.Disabled, vecty.Class("disabled")),
		vecty.MarkupIf(c.Inline, vecty.Class("form-check-inline")),
		vecty.MarkupIf(c.Radio, vecty.Class("form-check-radio")),
	).Apply(h)
}
