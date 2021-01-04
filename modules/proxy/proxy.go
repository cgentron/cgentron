package proxy

import (
	"bytes"
	"io"

	"github.com/cgentron/protoc-gen-cgentron/pkg/templates"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func New() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Module) Name() string { return "cgen" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	for _, f := range targets {
		m.genProxy(f, buf, m.Parameters())
	}

	return m.Artifacts()
}

func (m *Module) genProxy(f pgs.File, buf *bytes.Buffer, params pgs.Parameters) {
	m.Push(f.Name().String())
	defer m.Pop()

	buf.Reset()
	v := initCgenVisitor(m, buf, "", params)
	m.CheckErr(pgs.Walk(v, f), "unable to generate AST tree")

	out := buf.String()

	m.AddGeneratorFile(
		m.ctx.OutputPath(f).SetExt(".cgen.go").String(),
		out,
	)
}

type CgenVisitor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	pgs.Parameters
	prefix string
	w      io.Writer
}

func initCgenVisitor(d pgs.DebuggerCommon, w io.Writer, prefix string, params pgs.Parameters) CgenVisitor {
	p := CgenVisitor{
		prefix:         prefix,
		w:              w,
		Parameters:     params,
		DebuggerCommon: d,
	}

	p.Visitor = pgs.PassThroughVisitor(&p)

	return p
}

// VisitFile ...
func (p CgenVisitor) VisitFile(f pgs.File) (pgs.Visitor, error) {
	tpl, err := templates.File(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, f)
	if err != nil {
		return nil, err
	}

	return p, err
}

var _ pgs.Module = (*Module)(nil)
