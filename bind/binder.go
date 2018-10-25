package bind

// const tagName = "bind"

// type Binder struct {
// 	model    interface{}
// 	OnRender func()
// }

// func (b *Binder) OnEvent(e *vecty.Event) {
// 	id := ""
// 	if obj := e.Target.Get("id"); obj != nil {
// 		id = obj.String()
// 	}

// 	value := ""
// 	if obj := e.Target.Get("value"); obj != nil {
// 		value = obj.String()
// 	}

// 	v := reflect.ValueOf(b.model)
// 	if !v.IsValid() {
// 		return
// 	}
// 	i := reflect.Indirect(v)
// 	s := i.Type()

// 	bindingHandled := false
// 	for i := 0; i < s.NumField(); i++ {
// 		field := s.Field(i)
// 		if tag, ok := field.Tag.Lookup(tagName); ok {
// 			if tag == id {
// 				v.Elem().Field(i).SetString(value)
// 				bindingHandled = true
// 				break
// 			}
// 		}

// 	}

// 	if !bindingHandled {
// 		panic("there is no binding for " + id + " please check your tags!")
// 	}

// 	// render updated content??
// 	if b.OnRender != nil {
// 		b.OnRender()
// 	}
// }

// func (b *Binder) Notify() {
// 	return

// 	d := dom.GetWindow().Document()

// 	v := reflect.ValueOf(b.model)
// 	i := reflect.Indirect(v)
// 	s := i.Type()

// 	for i := 0; i < s.NumField(); i++ {

// 		// Get the field, returns https://golang.org/pkg/reflect/#StructField
// 		field := s.Field(i)

// 		// Get the field tag value
// 		tag := field.Tag.Get(tagName)

// 		//fmt.Println("notify " + tag)

// 		newValue := v.Elem().Field(i).String()

// 		h := d.GetElementByID(tag)
// 		fmt.Println("notify tag: " + tag + " value: " + newValue)
// 		//h.SetInnerHTML(v.Elem().Field(i).String())
// 		h.SetAttribute("value", newValue)
// 	}

// 	//	js.Global.Get("document").Call("write", "Hello world!")
// }

// func (b *Binder) Value(id string) string {

// 	v := reflect.ValueOf(b.model)
// 	if !v.IsValid() {
// 		return ""
// 	}

// 	i := reflect.Indirect(v)

// 	s := i.Type()

// 	for i := 0; i < s.NumField(); i++ {

// 		// Get the field, returns https://golang.org/pkg/reflect/#StructField
// 		field := s.Field(i)

// 		// Get the field tag value
// 		tag := field.Tag.Get(tagName)

// 		if tag == id {
// 			return v.Elem().Field(i).String()
// 		}
// 	}

// 	return ""
// }

// func Bind(m interface{}) Binder {
// 	rv := reflect.ValueOf(m)
// 	if rv.Kind() != reflect.Ptr {
// 		panic("model has to be assigned as ptr!")
// 	}
// 	return Binder{
// 		model: m,
// 	}
// }
