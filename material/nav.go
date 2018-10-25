package material

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type NavStyle int

const (
	NavLink NavStyle = iota
	NavTabs
)

type NavItem struct {
	vecty.Core
	index     int
	Name      string
	Active    bool
	Disabled  bool
	HRef      string
	IsVisible func(id string) bool
	OnClick   func(this *NavItem, e *vecty.Event)
}

func (c *NavItem) onClick(e *vecty.Event) {
	if c.OnClick != nil {
		c.Active = true
		c.OnClick(c, e)
	}
}

func (b *NavItem) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.Markup(
				b,
				prop.Href(b.HRef),
			),
			vecty.Text(b.Name),
		),
	)
}

func (b *NavItem) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("nav-link"),
		vecty.MarkupIf(b.Active, vecty.Class("active")),
		vecty.MarkupIf(b.Disabled, vecty.Class("disabled")),
	).Apply(h)
}

type NavItems []NavItem

type Nav struct {
	vecty.Core
	//OnClick   func(this *Nav, e *vecty.Event)
	Style     NavStyle
	Alignment Alignment
	NavItems  NavItems
	Vertical  bool
	Href      string
}

func (b *Nav) Render() vecty.ComponentOrHTML {
	var items vecty.List
	for i := range b.NavItems {
		isVisible := true
		if b.NavItems[i].IsVisible != nil {
			isVisible = b.NavItems[i].IsVisible(b.NavItems[i].Name)
		}
		if !isVisible {
			continue
		}

		items = append(items, &NavItem{
			index:    i,
			Name:     b.NavItems[i].Name,
			Active:   b.NavItems[i].Active,
			Disabled: b.NavItems[i].Disabled,
			HRef:     b.NavItems[i].HRef,
			OnClick:  b.onClick,
		})
	}

	return elem.UnorderedList(
		vecty.Markup(
			b,
		),
		items,
	)
}

func (b *Nav) Apply(h *vecty.HTML) {
	vecty.Markup(
		vecty.Class("nav"),
		vecty.MarkupIf(b.Alignment == Center, vecty.Class("justify-content-center")),
		vecty.MarkupIf(b.Alignment == Right, vecty.Class("justify-content-end")),
		vecty.MarkupIf(b.Vertical, vecty.Class("flex-column")),
	).Apply(h)
}

func (b *Nav) onClick(item *NavItem, e *vecty.Event) {
	for i := range b.NavItems {
		if item != &b.NavItems[i] {
			b.NavItems[i].Active = false
		}
	}
}
