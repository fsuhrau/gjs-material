package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type ButtonSize int

const (
	Regular ButtonSize = iota
	Small
	Large
)

type Button struct {
	vecty.Core
	Label      vecty.ComponentOrHTML
	Icon       vecty.ComponentOrHTML
	OnClick    func(this *Button, e *vecty.Event)
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

func (c *Button) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			c,
			c.Markup,
		),
		c.Icon,
		RenderStoredChild(c.Label),
	)
}

func (c *Button) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("btn"),
		prop.Type(prop.TypeButton),
		event.Click(c.onClick),
		vecty.Property("disabled", c.Disabled),
		vecty.MarkupIf(c.DataTarget != "", vecty.Attribute("data-target", c.DataTarget)),
		vecty.MarkupIf(c.DataToggle != "", vecty.Attribute("data-toggle", c.DataToggle)),
		vecty.MarkupIf(c.Rounded, vecty.Class("btn-round")),
		vecty.MarkupIf(c.Mini, vecty.Class("btn-fab", "btn-fab-mini")),
		vecty.MarkupIf(c.Link, vecty.Class("btn-link")),
		MarkupColor("btn", c.Color),
		vecty.MarkupIf(c.Size == Small, vecty.Class("btn-sm")),
		vecty.MarkupIf(c.Size == Large, vecty.Class("btn-lg")),
	).Apply(h)
}

func (c *Button) onClick(e *vecty.Event) {
	if c.OnClick != nil {
		c.OnClick(c, e)
	}
}
