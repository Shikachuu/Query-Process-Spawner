package cmd

import (
	"github.com/Shikachuu/php-process-redis-list/pkg/process"
	"github.com/Shikachuu/php-process-redis-list/pkg/queue"
	"github.com/spf13/cobra"
)

func redisListCommand(opts *projectOptions) *cobra.Command {
	var (
		redisHost     string
		redisPassword string
		redisDb       int
		redisList     string
	)
	cmd := &cobra.Command{
		Use:   "redis-list",
		Short: "Start qpp in redis-list listener mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			q, err := queue.NewRedisQueue(redisHost, redisPassword, redisDb, redisList)
			if err != nil {
				return err
			}

			qmc := make(chan string)
			go func() {
				_ = q.Listen(qmc)
			}()
			for {
				process.RunProcesses(qmc, opts.commandForWorkers, opts.numberOfWorkers)
			}
		},
	}
	cmd.Flags().StringVar(&redisHost, "host", "127.0.0.1:6379", "Set the address for the redis host")
	cmd.Flags().StringVar(&redisPassword, "password", "", "Set the password for the redis host")
	cmd.Flags().IntVar(&redisDb, "db", 0, "Set the DB for the redis host")
	cmd.Flags().StringVar(&redisList, "list", "", "Set the list that the app listens on")
	return cmd
}
