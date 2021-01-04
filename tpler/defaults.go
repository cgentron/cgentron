package tpler

import (
	"text/template"

	"github.com/cgentron/cgentron/tpler/iface"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type sharedFuncs struct {
	pgsgo.Context
}

func createContext(m pgs.Method)

func Defaults() iface.TplerRegisterFn {
	return func(tpl *template.Template, params pgs.Parameters) {
		fns := sharedFuncs{pgsgo.InitContext(params)}

		tpl.Funcs(map[string]interface {
		}{
			"pkg":  fns.PackageName,
			"name": fns.Name,
		})
	}
}
