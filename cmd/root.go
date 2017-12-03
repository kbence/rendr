package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "rendr",
	Short: "rendr is a rendering farm server/worker/client for Blender",
}

func init() {
	rootCmd.AddCommand(NewRunServerCommand())
	rootCmd.AddCommand(NewRunWorkerCommand())
}

// Execute runs randr command/subcommands
func Execute() {
	rootCmd.Execute()
}
