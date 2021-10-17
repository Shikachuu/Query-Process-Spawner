package cmd

import (
	"fmt"
	"github.com/Shikachuu/php-process-redis-list/internal"
	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version information",
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, _ []string) {
			printVersion()
		},
	}
}

func printVersion() {
	fmt.Printf("Version: %s\tCommit: %s\n", internal.Version, internal.GitCommit)
}
