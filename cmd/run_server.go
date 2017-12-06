package cmd

import (
	"github.com/kbence/rendr/rpc"
	"github.com/spf13/cobra"
)

// NewRunServerCommand creates the object for command `serve`
func NewRunServerCommand() *cobra.Command {
	return &cobra.Command{
		Use: "run-server",
		Run: func(cmd *cobra.Command, args []string) {
			rpc.Serve()
		},
	}
}
