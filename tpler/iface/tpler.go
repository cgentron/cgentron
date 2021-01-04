package iface

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

// TplerRegisterFn ...
type TplerRegisterFn func(tpl *template.Template, params pgs.Parameters)

// TplerContext ...
type TplerContext interface {
	// Typ ...
	Typ() string
	// Method ...
	Method() pgs.Method
}

// TplerContextFn ...
type TplerContextFn func(m pgs.Method) (TplerContext, error)

// Tpler ...
type Tpler interface {
	Make(string, string, pgs.Parameters, ...TplerRegisterFn) (*template.Template, error)
}
