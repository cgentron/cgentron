package templates

import (
	"text/template"

	"github.com/cgentron/cgentron/templates/shared"

	pgs "github.com/lyft/protoc-gen-star"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.Register(tpl, params)

	fn(tpl, params)
	return tpl
}

func File(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("file", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(fileTpl)
}
