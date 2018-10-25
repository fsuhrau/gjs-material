package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Select struct {
	vecty.Core
	Label   vecty.ComponentOrHTML
	OnClick func(this *Select, e *vecty.Event)
	Color   Color
	Markup  vecty.MarkupList
	Options []string
}

func (c *Select) Render() vecty.ComponentOrHTML {
	var options vecty.List
	for i := range c.Options {
		options = append(options, elem.Option(
			vecty.Text(c.Options[i]),
		))
	}

	return elem.Select(
		vecty.Markup(
			c,
			c.Markup,
		),
		options,
	)
}

func (c *Select) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("form-control", "selectpicker"),
		vecty.Data("style", "btn btn-link"),
	).Apply(h)
}

func (c *Select) onClick(e *vecty.Event) {
	if c.OnClick != nil {
		c.OnClick(c, e)
	}
}
