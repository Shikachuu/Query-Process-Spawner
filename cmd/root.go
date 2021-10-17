package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	var (
		version         bool
		redisHost       string
		redisPassword   string
		redisDb         int
		numberOfWorkers int
	)

	cmd := &cobra.Command{
		Use:   "runner",
		Short: "Start the worker multiplexer",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if version {
				return VersionCommand().Execute()
			}

			fmt.Println("App goes brrr")
			return nil
		},
	}
	// Flags
	cmd.Flags().StringVar(&redisHost, "redis.host", "127.0.0.1:6379", "Set the address for the redis host")
	cmd.Flags().StringVar(&redisPassword, "redis.password", "", "Set the password for the redis host")
	cmd.Flags().IntVar(&redisDb, "redis.db", 0, "Set the DB for the redis host")
	cmd.Flags().IntVar(&numberOfWorkers, "workers.max", 0, "Set the DB for the redis host")

	cmd.Flags().BoolVarP(&version, "version", "v", false, "Show the version information")
	// Commands
	cmd.AddCommand(VersionCommand())

	return cmd
}
