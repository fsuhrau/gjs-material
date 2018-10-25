package material

import (
	"github.com/fsuhrau/gjs-material/bind"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type TextArea struct {
	vecty.Core
	ID          string
	Placeholder string
	Type        prop.InputType
	Value       bind.Bind
	Rows        int
	Cols        int
}

// Render implements the vecty.Component interface.
func (c *TextArea) Render() vecty.ComponentOrHTML {
	return elem.TextArea(
		vecty.Markup(
			vecty.Property("rows", c.Rows),
			vecty.Property("cols", c.Cols),
			vecty.Class("form-control"),
			prop.ID(c.ID),
			vecty.MarkupIf(c.Placeholder != "", prop.Placeholder(c.Placeholder)),
			prop.Value(c.Value.Get()),
			event.Input(c.Value.OnEvent),
		),
	)
}
