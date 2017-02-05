package main

import (
	"html/template"
	"io/ioutil"
	"os"

	"github.com/russross/blackfriday"
)

type tmplData struct {
	VendorCSS      template.CSS
	VendorThemeCSS template.CSS
	CustomCSS      template.CSS
	MainContent    template.HTML
	Favicon        template.HTML
}

func main() {
	// css
	vendorCSS, err := ioutil.ReadFile("./node_modules/hack/dist/hack.css")
	if err != nil {
		panic(err)
	}
	vendorThemeCSS, err := ioutil.ReadFile("./node_modules/hack/dist/dark.css")
	if err != nil {
		panic(err)
	}
	customCSS, err := ioutil.ReadFile("./css/index.css")
	if err != nil {
		panic(err)
	}

	// favicon
	favicon, err := ioutil.ReadFile("./templates/favicon.html")
	if err != nil {
		panic(err)
	}

	// content
	content, err := ioutil.ReadFile("./content/index.md")
	if err != nil {
		panic(err)
	}
	contentHTML := blackfriday.MarkdownCommon(content)

	// template
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./www/index.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(f, &tmplData{
		VendorCSS:      template.CSS(string(vendorCSS)),
		VendorThemeCSS: template.CSS(string(vendorThemeCSS)),
		CustomCSS:      template.CSS(string(customCSS)),
		Favicon:        template.HTML(string(favicon)),
		MainContent:    template.HTML(string(contentHTML)),
	})
	if err != nil {
		panic(err)
	}
	f.Close()
}
