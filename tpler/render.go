package tpler

import (
	"bytes"
	"text/template"

	"github.com/cgentron/cgentron/tpler/iface"

	pgs "github.com/lyft/protoc-gen-star"
)

var _ iface.TplerContext = (*templateContext)(nil)

type templateContext struct {
	m   pgs.Method
	typ string
}

// Method ...
func (c *templateContext) Method() pgs.Method {
	return c.m
}

// Typ ...
func (c *templateContext) Typ() string {
	return c.typ
}

// RenderFunc ...
func RenderFunc(templates map[string]string, fn iface.TplerContextFn) iface.TplerRegisterFn {
	return func(tpl *template.Template, params pgs.Parameters) {
		for n, t := range templates {
			template.Must(tpl.New(n).Parse(t))
		}

		tpl.Funcs(map[string]interface {
		}{
			"context": fn,
			"render":  Render(tpl),
		})
	}
}

// Render ..
func Render(tpl *template.Template) func(ctx interface{}) (string, error) {
	return func(ctx interface{}) (string, error) {
		buf := &bytes.Buffer{}
		err := tpl.ExecuteTemplate(buf, ctx.(iface.TplerContext).Typ(), ctx)
		return buf.String(), err
	}
}
