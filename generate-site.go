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

	// index
	indexContent, err := ioutil.ReadFile("./content/index.md")
	if err != nil {
		panic(err)
	}
	indexContentHTML := blackfriday.MarkdownCommon(indexContent)
	indexData := &tmplData{
		VendorCSS:      template.CSS(string(vendorCSS)),
		VendorThemeCSS: template.CSS(string(vendorThemeCSS)),
		CustomCSS:      template.CSS(string(customCSS)),
		Favicon:        template.HTML(string(favicon)),
		MainContent:    template.HTML(string(indexContentHTML)),
	}
	err = parseTemplate("./templates/index.html", "./www/index.html", indexData)
	if err != nil {
		panic(err)
	}

	// 404
	notFoundContent, err := ioutil.ReadFile("./content/404.md")
	if err != nil {
		panic(err)
	}
	notFoundContentHTML := blackfriday.MarkdownCommon(notFoundContent)
	notFoundData := &tmplData{
		VendorCSS:      template.CSS(string(vendorCSS)),
		VendorThemeCSS: template.CSS(string(vendorThemeCSS)),
		CustomCSS:      template.CSS(string(customCSS)),
		Favicon:        template.HTML(string(favicon)),
		MainContent:    template.HTML(string(notFoundContentHTML)),
	}
	err = parseTemplate("./templates/index.html", "./www/404.html", notFoundData)
	if err != nil {
		panic(err)
	}
}

func parseTemplate(templateFile, destFile string, data *tmplData) error {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}
	f, err := os.Create(destFile)
	if err != nil {
		return err
	}
	err = t.Execute(f, data)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
