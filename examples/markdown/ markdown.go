package main

import (
	"fmt"

	"github.com/fsuhrau/gjs-material/bind"
	"github.com/fsuhrau/gjs-material/material"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	vecty.SetTitle("Markdown Demo")
	vecty.AddStylesheet("../assets/css/material-kit.css?v=2.0.4")
	vecty.AddStylesheet("https://maxcdn.bootstrapcdn.com/font-awesome/latest/css/font-awesome.min.css")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Roboto+Slab:400,700|Material+Icons")
	model := Model{
		Content: `# Markdown Example
This is a live editor, try editing the Markdown on the right of the page.
`,
	}

	vecty.RenderBody(&PageView{
		Model: model,
	})
}

type Model struct {
	Content string
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
	Model Model
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Style("float", "right"),
			),
			&material.TextArea{
				Rows: 14,
				Cols: 70,
				Value: bind.ValueRerender(&p.Model.Content, func() {
					fmt.Println(p.Model)
					vecty.Rerender(p)
				}),
			},
		),

		// Render the markdown.
		&Markdown{
			Input: p.Model.Content,
		},
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
		// 	vecty.Text("$(document).ready(function(){\n    \n    \n    var _gaq = _gaq || [];\n_gaq.push(['_setAccount', 'UA-46172202-1']);\n_gaq.push(['_trackPageview']);\n\n(function() {\n    var ga = document.createElement('script');\n    ga.type = 'text/javascript';\n    ga.async = true;\n    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';\n    var s = document.getElementsByTagName('script')[0];\n    s.parentNode.insertBefore(ga, s);\n})();\n\n    // Facebook Pixel Code Don't Delete\n!function(f,b,e,v,n,t,s){if(f.fbq)return;n=f.fbq=function(){n.callMethod?\nn.callMethod.apply(n,arguments):n.queue.push(arguments)};if(!f._fbq)f._fbq=n;\nn.push=n;n.loaded=!0;n.version='2.0';n.queue=[];t=b.createElement(e);t.async=!0;\nt.src=v;s=b.getElementsByTagName(e)[0];s.parentNode.insertBefore(t,s)}(window,\ndocument,'script','//connect.facebook.net/en_US/fbevents.js');\n\ntry{\n  fbq('init', '111649226022273');\n  fbq('track', \"PageView\");\n\n}catch(err) {\n  console.log('Facebook Track Error:', err);\n}\n\n  });"),
		// ),
	)
}

// Markdown is a simple component which renders the Input markdown as sanitized
// HTML into a div.
type Markdown struct {
	vecty.Core
	Input string `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (m *Markdown) Render() vecty.ComponentOrHTML {
	// Render the markdown input into HTML using Blackfriday.
	unsafeHTML := blackfriday.MarkdownCommon([]byte(m.Input))

	// Sanitize the HTML.
	safeHTML := string(bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML))

	// Return the HTML, which we guarantee to be safe / sanitized.
	return elem.Div(
		vecty.Markup(
			vecty.UnsafeHTML(safeHTML),
		),
	)
}
