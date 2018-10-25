package main

import (
	"fmt"

	"github.com/fsuhrau/gjs-material/bind"
	"github.com/fsuhrau/gjs-material/material"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.SetTitle("Card Demo")
	vecty.AddStylesheet("../assets/css/material-kit.css?v=2.0.4")
	vecty.AddStylesheet("https://maxcdn.bootstrapcdn.com/font-awesome/latest/css/font-awesome.min.css")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Roboto+Slab:400,700|Material+Icons")
	vecty.RenderBody(&CardView{})
}

// CardView is our main page component.
type CardView struct {
	vecty.Core
	Email    string
	Password string
}

// Render implements the vecty.Component interface.
func (p *CardView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("row"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("col-md-6"),
				),
				&material.Card{
					Header: &material.CardHeader{
						Color:    material.Primary,
						Title:    "Card Title",
						SubTitle: "Card Subtitle",
					},
					Body: elem.Form(
						vecty.Markup(
							vecty.Class("contact-form"),
						),
						&material.FormField{
							Label: "Email",
							Input: &material.Input{
								Type:        prop.TypeEmail,
								ID:          "Email",
								Placeholder: "Email",
								Value:       bind.Value(&p.Email),
							},
						},
						&material.FormField{
							Label: "Password",
							Input: &material.Input{
								Type:        prop.TypePassword,
								ID:          "password",
								Placeholder: "Password",
								Value:       bind.Value(&p.Password),
							},
						},
						&material.Button{
							Color: material.Primary,
							Label: vecty.Text("Submit"),
							OnClick: func(this *material.Button, e *vecty.Event) {
								js.Global.Call("alert", fmt.Sprintf("email: %s password: %s", p.Email, p.Password))
							},
						},
					),
				},
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/core/jquery.min.js"),
				vecty.Attribute("type", "text/javascript"),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/core/popper.min.js"),
				vecty.Attribute("type", "text/javascript"),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/core/bootstrap-material-design.min.js"),
				vecty.Attribute("type", "text/javascript"),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/plugins/moment.min.js"),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/plugins/bootstrap-datetimepicker.js"),
				vecty.Attribute("type", "text/javascript"),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/plugins/nouislider.min.js"),
				vecty.Attribute("type", "text/javascript"),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/plugins/jquery.sharrre.js"),
				vecty.Attribute("type", "text/javascript"),
			),
		),
		elem.Script(
			vecty.Markup(
				prop.Src("../assets/js/material-kit.min.js?v=2.0.4"),
				vecty.Attribute("type", "text/javascript"),
			),
		),
		// elem.Script(
		// 	vecty.Text("$(document).ready(function(){\n    \n    \n    var _gaq = _gaq || [];\n_gaq.push(['_setAccount', 'UA-46172202-1']);\n_gaq.push(['_trackCardView']);\n\n(function() {\n    var ga = document.createElement('script');\n    ga.type = 'text/javascript';\n    ga.async = true;\n    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';\n    var s = document.getElementsByTagName('script')[0];\n    s.parentNode.insertBefore(ga, s);\n})();\n\n    // Facebook Pixel Code Don't Delete\n!function(f,b,e,v,n,t,s){if(f.fbq)return;n=f.fbq=function(){n.callMethod?\nn.callMethod.apply(n,arguments):n.queue.push(arguments)};if(!f._fbq)f._fbq=n;\nn.push=n;n.loaded=!0;n.version='2.0';n.queue=[];t=b.createElement(e);t.async=!0;\nt.src=v;s=b.getElementsByTagName(e)[0];s.parentNode.insertBefore(t,s)}(window,\ndocument,'script','//connect.facebook.net/en_US/fbevents.js');\n\ntry{\n  fbq('init', '111649226022273');\n  fbq('track', \"CardView\");\n\n}catch(err) {\n  console.log('Facebook Track Error:', err);\n}\n\n  });"),
		// ),
	)
}
