package cgentron

import (
	"go.uber.org/zap"

	pb "github.com/cgentron/api/proto"
)

// ResolverRules ...
type ResolverRules map[string]*pb.ResolverRule

const (
	// DefaultAddr ...
	DefaultAddr = "0.0.0.0:9090"
	// DefaultStatusAddr ...
	DefaultStatusAddr = ":8443"
	// DefaultLogFormat ...
	DefaultLogFormat = "text"
	// DefaultLogLevel ...
	DefaultLogLevel = "warn"
	// DefaultVerbose ...
	DefaultVerbose = false
	// DefaultDebug ...
	DefaultDebug = false
)

// Opts ...
type Opts struct {
	// Verbose ...
	Verbose bool
	// Debug ...
	Debug bool
	// Addr ...
	Addr string
	// StatusAddr ...
	StatusAddr string
	// LogFormat ...
	LogFormat string
	// LogLevel ...
	LogLevel string
	// Logger ...
	Logger *zap.Logger
	// Resolvers ...
	Resolvers ResolverRules
}

// Opt ...
type Opt func(*Opts)

// DefaultOpts ...
func NewDefaultOpts() *Opts {
	return &Opts{
		Addr:       DefaultAddr,
		StatusAddr: DefaultStatusAddr,
		LogFormat:  DefaultLogFormat,
		LogLevel:   DefaultLogLevel,
		Debug:      DefaultDebug,
		Verbose:    DefaultVerbose,
		Resolvers:  make(ResolverRules),
	}
}

// NewOpts ...
func NewOpts(opts ...Opt) *Opts {
	o := NewDefaultOpts()
	o.Configure(opts...)

	return o
}

// Configure ...
func (s *Opts) Configure(opts ...Opt) error {
	for _, o := range opts {
		o(s)
	}

	return nil
}

// ConfigureLogger ...
func ConfigureLogger() error {
	return nil
}

// WithLogger ...
func WithLogger(logger *zap.Logger) Opt {
	return func(opts *Opts) {
		opts.Logger = logger
	}
}

// WithVerbose ...
func WithVerbose() Opt {
	return func(opts *Opts) {
		opts.Verbose = true
	}
}

// WithDebug ...
func WithDebug() Opt {
	return func(opts *Opts) {
		opts.Debug = true
	}
}

// WithAddr ...
func WithAddr(addr string) Opt {
	return func(opts *Opts) {
		opts.Addr = addr
	}
}

// WithStatusAddr ...
func WithStatusAddr(addr string) Opt {
	return func(opts *Opts) {
		opts.StatusAddr = addr
	}
}

// WithResolvers ...
func WithResolvers(resolvers ResolverRules) Opt {
	return func(opts *Opts) {
		for k, r := range resolvers {
			opts.Resolvers[k] = r
		}
	}
}
