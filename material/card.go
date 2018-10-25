package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type CardHeader struct {
	vecty.Core
	Color         Color
	Title         string
	SubTitle      string
	Icon          *Icon
	Text          string
	Custom        vecty.ComponentOrHTML
	TextAlignment Alignment
}

func (c *CardHeader) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			c,
		),
		vecty.If(c.Title != "",
			elem.Heading4(
				vecty.Markup(
					vecty.Class("card-title"),
				),
				vecty.Text(c.Title),
			),
		),
		vecty.If(c.SubTitle != "",
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("category"),
				),
				vecty.Text(c.SubTitle),
			),
		),
		vecty.If(c.Text != "",
			elem.Div(
				vecty.Markup(
					vecty.Class("card-text"),
				),
				elem.Heading4(
					vecty.Markup(
						vecty.Class("card-title"),
					),
					vecty.Text(c.Text),
				),
			),
		),
		vecty.If(c.Icon != nil,
			elem.Div(
				vecty.Markup(
					vecty.Class("card-icon"),
				),
				c.Icon,
			),
		),
		c.Custom,
	)
}

func (c *CardHeader) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("card-header"),
		vecty.MarkupIf(c.Icon != nil, vecty.Class("card-header-icon")),
		vecty.MarkupIf(c.Text != "", vecty.Class("card-header-text")),
		MarkupColor("card-header", c.Color),
		vecty.MarkupIf(c.TextAlignment == Center, vecty.Class("text-center")),
		vecty.MarkupIf(c.TextAlignment == Right, vecty.Class("text-right")),
	).Apply(h)
}

type CardImage struct {
	vecty.Core
	Src    string
	Alt    string
	Bottom bool
}

func (c *CardImage) Render() vecty.ComponentOrHTML {
	return elem.Image(
		vecty.Markup(
			vecty.MarkupIf(!c.Bottom, vecty.Class("card-img-top")),
			vecty.MarkupIf(c.Bottom, vecty.Class("card-img-bottom")),
			prop.Src(c.Src),
			vecty.Attribute("alt", c.Alt),
		),
	)
}

type Card struct {
	vecty.Core
	Color         Color
	TextColor     Color
	TextAlignment Alignment
	Header        *CardHeader
	Body          vecty.ComponentOrHTML
	Footer        vecty.ComponentOrHTML
	Image         vecty.ComponentOrHTML
	ImageOverlay  vecty.ComponentOrHTML
}

func (c *Card) Render() vecty.ComponentOrHTML {
	var topImage vecty.ComponentOrHTML
	var bottomImage vecty.ComponentOrHTML
	var topOverlay vecty.ComponentOrHTML
	var bottomOverlay vecty.ComponentOrHTML
	if img, ok := c.Image.(*CardImage); ok {
		if !img.Bottom {
			topImage = c.Image
			topOverlay = c.ImageOverlay
		} else {
			bottomImage = c.Image
			bottomOverlay = c.ImageOverlay
		}
	}

	return elem.Div(
		vecty.Markup(
			c,
		),
		vecty.If(c.Header != nil, c.Header),
		topImage,
		topOverlay,
		elem.Div(
			vecty.Markup(
				vecty.Class("card-body"),
			),
			c.Body,
		),
		bottomImage,
		bottomOverlay,
		vecty.If(c.Footer != nil,
			elem.Div(
				vecty.Markup(
					vecty.Class("card-footer", "text-muted"),
				),
				c.Footer,
			),
		),
	)
}

func (c *Card) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("card"),
		vecty.MarkupIf(c.TextAlignment == Center, vecty.Class("text-center")),
		vecty.MarkupIf(c.TextAlignment == Right, vecty.Class("text-right")),
		MarkupColor("bg", c.Color),
		MarkupColor("text", c.TextColor),
	).Apply(h)
}
