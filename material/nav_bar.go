package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type NavBarItem struct {
	vecty.Core
	index     int
	Name      string
	Active    bool
	Disabled  bool
	HRef      string
	IsVisible func(id string) bool
	OnClick   func(this *NavBarItem, e *vecty.Event)
}

func (c *NavBarItem) onClick(e *vecty.Event) {
	if c.OnClick != nil {
		c.Active = true
		c.OnClick(c, e)
	}
}

func (b *NavBarItem) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.Markup(
			b,
		),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("nav-link"),
				prop.Href(b.HRef),
			),
			vecty.If(!b.Active, vecty.Text(b.Name)),
			vecty.If(b.Active,
				vecty.Text(b.Name+" ",
					elem.Span(
						vecty.Markup(
							vecty.Class("sr-only"),
						),
						vecty.Text("(current)"),
					),
				),
			),
		),
	)
}

func (b *NavBarItem) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("nav-item"),
		vecty.MarkupIf(b.Active, vecty.Class("active")),
		vecty.MarkupIf(b.Disabled, vecty.Class("disabled")),
	).Apply(h)
}

type NavBarItems []NavBarItem

type NavBarPlacement int

const (
	NavBarPlacementDefault NavBarPlacement = iota
	NavBarPlacementFixedTop
	NavBarPlacementFixedBottom
	NavBarPlacementStickyTop
)

type NavBar struct {
	vecty.Core
	OnClick       func(this *NavBar, e *vecty.Event)
	Color         Color
	LightColor    bool
	ColorOnScroll int
	Placement     NavBarPlacement
	Items         NavBarItems
	Brand         vecty.ComponentOrHTML
}

func (c *NavBar) Render() vecty.ComponentOrHTML {
	var items vecty.List
	for i := range c.Items {
		isVisible := true
		if c.Items[i].IsVisible != nil {
			isVisible = c.Items[i].IsVisible(c.Items[i].Name)
		}

		if !isVisible {
			continue
		}

		items = append(items, &NavBarItem{
			index:    i,
			Name:     c.Items[i].Name,
			Active:   c.Items[i].Active,
			Disabled: c.Items[i].Disabled,
			HRef:     c.Items[i].HRef,
			OnClick:  c.onClick,
		})
	}
	return elem.Navigation(
		vecty.Markup(
			c,
		),
		elem.Div(
			vecty.Markup(vecty.Class("container")),
			vecty.If(c.Brand != nil,
				c.Brand,
				// elem.Anchor(
				// 	vecty.Markup(
				// 		vecty.Class("navbar-brand"),
				// 		prop.Href("#"),
				// 	),
				// 	vecty.Text(),
				// ),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("navbar-toggler"),
					prop.Type(prop.TypeButton),
					vecty.Data("toggle", "collapse"),
					vecty.Attribute("aria-expanded", "false"),
					vecty.Attribute("aria-label", "Toggle navigation"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("sr-only"),
					),
					vecty.Text("Toggle navigation"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("navbar-toggler-icon"),
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("navbar-toggler-icon"),
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("navbar-toggler-icon"),
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("collapse", "navbar-collapse"),
				),
				elem.UnorderedList(
					vecty.Markup(
						vecty.Class("navbar-nav", "mr-auto"),
					),
					items,
				),
			),
		),
	)
}

func (c *NavBar) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("navbar", "navbar-expand-lg"),
		vecty.MarkupIf(c.Placement == NavBarPlacementFixedTop, vecty.Class("fixed-top")),
		vecty.MarkupIf(c.Placement == NavBarPlacementFixedBottom, vecty.Class("fixed-bottom")),
		vecty.MarkupIf(c.Placement == NavBarPlacementStickyTop, vecty.Class("sticky-top")),
		vecty.MarkupIf(c.LightColor, vecty.Class("navbar-light")),
		vecty.MarkupIf(!c.LightColor, vecty.Class("navbar-dark")),
		MarkupColor("bg", c.Color),
		vecty.MarkupIf(c.Color == Transparent, vecty.Class("navbar-transparent")),
		vecty.MarkupIf(c.Color == Transparent && c.ColorOnScroll > 0,
			vecty.Class("navbar-color-on-scroll"),
			vecty.Attribute("color-on-scroll", c.ColorOnScroll),
		),
	).Apply(h)
}

func (b *NavBar) onClick(item *NavBarItem, e *vecty.Event) {
	for i := range b.Items {
		if item != &b.Items[i] {
			b.Items[i].Active = false
		}
	}
}
