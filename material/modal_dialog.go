package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type ModalDialog struct {
	vecty.Core
	ID         string
	Title      string
	Detail     string
	CancelText string
	AcceptText string
	Open       bool
	NoBackdrop bool
	Scrollable bool
	AcceptBtn  *Button
	CancelBtn  *Button
	OnAccept   func(this *ModalDialog, e *vecty.Event)
	OnCancel   func(this *ModalDialog, e *vecty.Event)
}

// Render implements the vecty.Component interface.
func (c *ModalDialog) Render() vecty.ComponentOrHTML {

	acceptButton := c.AcceptBtn
	if acceptButton == nil {
		acceptButton = &Button{
			Color: Primary,
			Label: vecty.Text(c.AcceptText),
		}
	}

	cancelButton := c.CancelBtn
	if cancelButton == nil {
		cancelButton = &Button{
			Color: Secondary,
			Label: vecty.Text(c.CancelText),
			//Data:  event.Click(c.OnCancel),
		}
	}

	// Built-in root element.
	return elem.Div(
		vecty.Markup(
			vecty.Property("id", c.ID),
			vecty.Class("modal"),
			vecty.Class("fade"),
			vecty.Attribute("tabindex", "-1"),
			vecty.Attribute("role", "dialog"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("modal-dialog"),
				vecty.Attribute("role", "document"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("modal-content"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("modal-header"),
					),
					elem.Heading5(
						vecty.Markup(
							vecty.Class("modal-title"),
						),
						vecty.Text(c.Title),
					),
					elem.Button(
						vecty.Markup(
							prop.Type(prop.TypeButton),
							vecty.Class("close"),
							vecty.Data("dismiss", "modal"),
							vecty.Attribute("aria-label", "Close"),
						),
						elem.Span(
							vecty.Markup(
								vecty.Attribute("aria-hidden", "true"),
							),
							vecty.Text("&times;"),
						),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("modal-body"),
					),
					elem.Paragraph(
						vecty.Text("Modal body text goes here."),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("modal-footer"),
					),
					acceptButton,
					cancelButton,
				),
			),
		),
	)
}

// func (c *ModalDialog) Apply(h *vecty.HTML) {
// 	vecty.Markup(
// 		vecty.Class("mdc-dialog"),
// 		vecty.MarkupIf(d.Role == "", vecty.Attribute("role", "dialog")),
// 		vecty.MarkupIf(d.Role != "", vecty.Attribute("role", d.Role)),
// 		vecty.MarkupIf(d.Open, vecty.Class("mdc-dialog--open")),
// 		vecty.MarkupIf(!d.Open, vecty.Attribute("aria-hidden", "true")),
// 		d.ariaLabelledBy(h),
// 		d.ariaDescribedBy(h),
// 	).Apply(h)
// }

// func (d *Dialog) labelID(h *vecty.HTML) string {
// 	id := FindID(h)
// 	if id == "" {
// 		return ""
// 	}
// 	return id + "-label"
// }

// func (d *Dialog) ariaLabelledBy(h *vecty.HTML) vecty.Applyer {
// 	if d.labelID(h) == "" {
// 		return nil
// 	}
// 	return vecty.Attribute("aria-labelledby", d.labelID(h))
// }

// func (d *Dialog) descriptionID(h *vecty.HTML) string {
// 	id := FindID(h)
// 	if id == "" {
// 		return ""
// 	}
// 	return id + "-description"
// }

// func (d *Dialog) ariaDescribedBy(h *vecty.HTML) vecty.Applyer {
// 	if d.descriptionID(h) == "" {
// 		return nil
// 	}
// 	return vecty.Attribute("aria-describedby", d.descriptionID(h))
// }

func (d *ModalDialog) onCancel(e *vecty.Event) {
	if d.OnCancel != nil {
		d.OnCancel(d, e)
	}
}

func (d *ModalDialog) onAccept(e *vecty.Event) {
	if d.OnAccept != nil {
		d.OnAccept(d, e)
	}
}
