package bind

import (
	"github.com/gopherjs/vecty"
)

type Bind interface {
	Get() string
	OnEvent(e *vecty.Event)
}
