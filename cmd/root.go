package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runner",
		Short: "Start the worker multiplexer",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println("App goes brrr")
		},
	}

	cmd.AddCommand(VersionCommand())

	return cmd
}
