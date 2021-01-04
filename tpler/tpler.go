package tpler

import (
	"text/template"

	"github.com/cgentron/cgentron/tpler/iface"
	pgs "github.com/lyft/protoc-gen-star"
)

var _ iface.Tpler = (*tpler)(nil)

// Opts ...
type Opts struct {
}

// Opt ...
type Opt func(*Opts)

// Configure ...
func (s *Opts) Configure(opts ...Opt) error {
	for _, o := range opts {
		o(s)
	}

	return nil
}

type tpler struct {
	opts *Opts
}

// New ...
func New(opts ...Opt) iface.Tpler {
	options := new(Opts)
	options.Configure(opts...)

	t := new(tpler)
	t.opts = options

	return t
}

func (t *tpler) Make(name string, tpl string, params pgs.Parameters, fns ...iface.TplerRegisterFn) (*template.Template, error) {
	tt := template.New(name)

	for _, fn := range fns {
		fn(tt, params)
	}

	return tt.Parse(tpl)
}
