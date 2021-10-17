package root

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
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Printf("Version: %s\n", internal.Version)
			return nil
		},
	}
}
