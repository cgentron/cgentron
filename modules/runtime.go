package modules

import (
	"github.com/spf13/cobra"
)

type ModuleHandler func()

type Runtime struct {
	Opts *Opts

	cmd *cobra.Command
}

func New(opts ...Opt) *Runtime {
	options := NewOpts(opts...)

	r := new(Runtime)
	r.Opts = options

	cmd := &cobra.Command{}
	cmd.Flags().StringVar(&r.Opts.Input, "input", r.Opts.Input, "input")
	cmd.Flags().BoolVar(&r.Opts.Verbose, "verbose", r.Opts.Verbose, "verbose")

	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	r.cmd = cmd

	return r
}

func (r *Runtime) Execute(handler ModuleHandler) error {
	cmd := r.cmd
	cmd.Run = r.run(handler)

	return cmd.Execute()
}

func (r *Runtime) run(handler ModuleHandler) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		handler()
	}
}
