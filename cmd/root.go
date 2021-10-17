package cmd

import (
	"github.com/Shikachuu/php-process-redis-list/pkg/queue"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	var (
		version         bool
		redisHost       string
		redisPassword   string
		redisDb         int
		redisList       string
		numberOfWorkers int
	)

	cmd := &cobra.Command{
		Use:   "runner",
		Short: "Start the worker multiplexer",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if version {
				printVersion()
				return nil
			}
			q, err := queue.NewRedisQueue(redisHost, redisPassword, redisDb, redisList)
			if err != nil {
				return err
			}
			qmc := make(chan string)
			go func() {
				_ = q.Listen(qmc)
			}()
			<-qmc
			return nil
		},
	}
	// Flags
	cmd.Flags().StringVar(&redisHost, "redis.host", "127.0.0.1:6379", "Set the address for the redis host")
	cmd.Flags().StringVar(&redisPassword, "redis.password", "", "Set the password for the redis host")
	cmd.Flags().IntVar(&redisDb, "redis.db", 0, "Set the DB for the redis host")
	cmd.Flags().StringVar(&redisList, "redis.list", "", "Set the list that the app listens on")
	cmd.Flags().IntVar(&numberOfWorkers, "workers.max", 1, "Set the DB for the redis host")

	cmd.Flags().BoolVarP(&version, "version", "v", false, "Show the version information")
	// Commands
	cmd.AddCommand(VersionCommand())

	return cmd
}
