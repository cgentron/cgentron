package cgentron

import (
	"context"

	"github.com/andersnormal/pkg/debug"
	"github.com/andersnormal/pkg/server"
	"github.com/spf13/cobra"
)

type proxy struct {
	opts     *Opts
	cmd      *cobra.Command
	listener Listener
}

// Proxy ...
type Proxy interface {
	// Start ...
	Start(context.Context) error
}

// Listener
type Listener server.Listener

// New ..
func New(l Listener, opts *Opts) Proxy {
	p := new(proxy)
	p.opts = opts
	p.listener = l

	return p
}

// WithContext ...
func (p *proxy) Start(ctx context.Context) error {
	// create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create server
	s, _ := server.WithContext(ctx)

	if p.opts.Debug {
		// debug listener
		d := debug.New(
			debug.WithPprof(),
			debug.WithStatusAddr(p.opts.StatusAddr),
		)

		s.Listen(d, true)
	}

	// listen for grpc
	s.Listen(p.listener, true)

	// listen for the server and wait for it to fail,
	// or for sys interrupts
	if err := s.Wait(); err != nil {
		return err
	}

	// noop
	return nil
}

// +private

func configure(p *proxy, opts ...Opt) error {
	for _, o := range opts {
		o(p.opts)
	}

	return nil
}
