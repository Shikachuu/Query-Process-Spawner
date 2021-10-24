package cmd

import (
	"github.com/spf13/cobra"
)

type projectOptions struct {
	numberOfWorkers   int
	commandForWorkers []string
}

func RootCommand() *cobra.Command {
	var (
		projectOpts projectOptions
		version     bool
	)

	cmd := &cobra.Command{
		Use:              "qpp",
		Short:            "Start the QPP to spawn processes based on queue messages",
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if version {
				printVersion()
				return nil
			}
			return cmd.Help()
		},
	}
	// Flags
	cmd.Flags().IntVar(&projectOpts.numberOfWorkers, "max-workers", 1, "Set the max number of running processes")
	cmd.Flags().StringSliceVarP(&projectOpts.commandForWorkers, "command", "c", []string{""}, "Set the command that the workers will execute")
	cmd.Flags().BoolVarP(&version, "version", "v", false, "Show the version information")

	// Commands
	cmd.AddCommand(VersionCommand())
	cmd.AddCommand(redisListCommand(&projectOpts))

	return cmd
}
