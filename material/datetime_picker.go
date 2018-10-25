package material

import (
	"github.com/fsuhrau/gjs-material/bind"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type DateTimePicker struct {
	vecty.Core
	ID          string
	Placeholder string
	Value       bind.Bind
}

// Render implements the vecty.Component interface.
func (c *DateTimePicker) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("form-group"),
		),
		elem.Label(
			vecty.Markup(
				vecty.Class("label-control"),
			),
			vecty.Text(c.Placeholder),
		),
		elem.Input(
			vecty.Markup(
				prop.ID(c.ID),
				prop.Type(prop.TypeText),
				vecty.Class("form-control", "datetimepicker"),
				vecty.MarkupIf(c.Placeholder != "", prop.Placeholder(c.Placeholder)),
				event.Input(c.Value.OnEvent),
			),
		),
		elem.Script(
			vecty.Text("$('.datetimepicker').datetimepicker({icons: {time: \"fa fa-clock-o\",date: \"fa fa-calendar\",up: \"fa fa-chevron-up\",down: \"fa fa-chevron-down\",previous: 'fa fa-chevron-left',next: 'fa fa-chevron-right',today: 'fa fa-screenshot',clear: 'fa fa-trash',close: 'fa fa-remove'}});"),
		),
	)
}
