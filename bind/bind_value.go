package bind

import (
	"github.com/gopherjs/vecty"
)

type BindValue struct {
	Value    *string
	rerender func()
}

func Value(value *string) *BindValue {
	return &BindValue{
		Value: value,
	}
}

func ValueRerender(value *string, rerender func()) *BindValue {
	return &BindValue{
		Value:    value,
		rerender: rerender,
	}
}

func (b *BindValue) Get() string {
	if b.Value == nil {
		return ""
	}
	return *b.Value
}

func (b *BindValue) OnEvent(e *vecty.Event) {

	value := ""
	if obj := e.Target.Get("value"); obj != nil {
		value = obj.String()
	}

	if b.Value != nil {
		*b.Value = value
	}

	if b.rerender != nil {
		b.rerender()
	}
}
