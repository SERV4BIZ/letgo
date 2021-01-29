package utility

import (
	"errors"
	"regexp"

	minify "github.com/tdewolff/minify"
	minify_css "github.com/tdewolff/minify/css"
	minify_html "github.com/tdewolff/minify/html"
	minify_js "github.com/tdewolff/minify/js"
	minify_json "github.com/tdewolff/minify/json"
	minify_svg "github.com/tdewolff/minify/svg"
	minify_xml "github.com/tdewolff/minify/xml"
)

var m *minify.M = nil

// Minify is compress code
func Minify(exttype string, buffer []byte) ([]byte, error) {
	if m == nil {
		m = minify.New()
		m.AddFunc("text/css", minify_css.Minify)
		m.AddFunc("text/html", minify_html.Minify)
		m.AddFunc("image/svg+xml", minify_svg.Minify)
		m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), minify_js.Minify)
		m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), minify_json.Minify)
		m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), minify_xml.Minify)
	}

	if exttype == "css" {
		return m.Bytes("text/css", buffer)
	} else if exttype == "html" {
		return m.Bytes("text/html", buffer)
	} else if exttype == "svg" {
		return m.Bytes("image/svg+xml", buffer)
	} else if exttype == "js" {
		return m.Bytes("application/javascript", buffer)
	} else if exttype == "json" {
		return m.Bytes("application/json", buffer)
	} else if exttype == "xml" {
		return m.Bytes("text/xml", buffer)
	}
	return buffer, errors.New("Not support")
}
