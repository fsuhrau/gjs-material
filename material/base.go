package material

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Color int

const (
	Default Color = iota
	Primary
	Secondary
	Info
	Success
	Danger
	Warning
	Transparent
	Dark
	White
)

func MarkupColor(prefix string, color Color) vecty.Applyer {
	c := ""
	switch color {
	case Primary:
		c = "primary"
	case Secondary:
		c = "secondary"
	case Info:
		c = "info"
	case Success:
		c = "success"
	case Danger:
		c = "danger"
	case Warning:
		c = "warning"
	case Transparent:
		c = "transparent"
	case Dark:
		c = "dark"
	case White:
		c = "white"
	}
	return vecty.MarkupIf(c != "", vecty.Class(prefix+"-"+c))
}

type Alignment int

const (
	Left Alignment = iota
	Center
	Right
)

type nativeInputer interface {
	NativeInput() (*vecty.HTML, string)
}

func findProp(key string, h *vecty.HTML) *js.Object {
	k := js.InternalObject(h)
	if k == js.Undefined {
		return nil
	}
	k = k.Get("properties")
	if k == js.Undefined {
		return nil
	}
	k = k.Get("$" + key)
	if k == js.Undefined {
		return nil
	}
	return k.Get("v").Get("$val")
}

func FindID(moc vecty.MarkupOrChild) string {
	switch t := moc.(type) {
	case *vecty.MarkupList:
		return FindID(*t)
	case nativeInputer:
		_, id := t.NativeInput()
		return id
	case vecty.Applyer:
		d := elem.Div()
		t.Apply(d)
		return FindID(d)
	case *vecty.HTML:
		id := findProp("id", t)
		if id == nil {
			return ""
		}
		return id.String()
	case vecty.Component:
		if h, ok := t.Render().(*vecty.HTML); ok {
			return FindID(h)
		}
	}
	return ""
}

func MarkupOnly(moc vecty.MarkupOrChild) *vecty.MarkupList {
	// TODO: handle vecty.List
	switch t := moc.(type) {
	case vecty.ComponentOrHTML:
		return nil
	case vecty.MarkupList:
		return &t
	}
	return nil
}

type StaticComponent struct {
	vecty.Core
	Child vecty.ComponentOrHTML
}

// RenderStoredChild is a helper which provides a Component which wraps the
// provided ComponentOrHTML. It exists as a workaround to a vecty issue.
//
// See: https://github.com/gopherjs/vecty/issues/191
func RenderStoredChild(child vecty.ComponentOrHTML) *StaticComponent {
	return &StaticComponent{Child: child}
}

func (c *StaticComponent) Render() vecty.ComponentOrHTML {
	switch t := c.Child.(type) {
	case vecty.List:
		return elem.Div(t)
	}
	return c.Child
}

func (c *StaticComponent) SkipRender(prev vecty.Component) bool {
	switch prev.(type) {
	case *StaticComponent:
		return true
	}
	return false
}
