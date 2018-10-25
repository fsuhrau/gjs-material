package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FormField is a vecty-material formfield component.
type FormField struct {
	vecty.Core
	Input    vecty.ComponentOrHTML
	Label    string
	Help     string
	AlignEnd bool
}

// Render implements the vecty.Component interface.
func (f *FormField) Render() vecty.ComponentOrHTML {
	inputID := FindID(f.Input)
	return elem.Div(
		vecty.Markup(
			f,
		),
		elem.Label(
			vecty.Markup(
				vecty.MarkupIf(inputID != "",
					prop.For(inputID),
				),
			),
			vecty.Text(f.Label),
		),
		f.Input,
		vecty.If(f.Help != "",
			elem.Small(
				vecty.Markup(
					vecty.Class("form-text", "text-muted"),
				),
				vecty.Text(f.Help),
			),
		),
	)
}

func (f *FormField) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("form-group"),
	).Apply(h)
}
