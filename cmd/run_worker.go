package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

// NewRunWorkerCommand creates the object for command `serve`
func NewRunWorkerCommand() *cobra.Command {
	return &cobra.Command{
		Use: "run-worker",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				time.Sleep(1000)
			}
		},
	}
}
