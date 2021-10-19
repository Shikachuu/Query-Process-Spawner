package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Queue struct {
	client *redis.Client
	ctx    context.Context
	queue  string
}

func NewRedisQueue(host string, password string, db int, queue string) (*Queue, error) {
	q := &Queue{
		ctx: context.Background(),
		client: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password, // no password set
			DB:       db,       // use default DB
		}),
		queue: queue,
	}
	return q, q.client.Ping(q.ctx).Err()
}

func (q *Queue) Receive() (string, error) {
	return q.client.RPop(q.ctx, q.queue).Result()
}

func (q *Queue) Listen(i chan string) error {
	for {
		msg, err := q.Receive()
		if err != nil {
			if err.Error() == "redis: nil" {
				continue
			}
			return err
		}
		i <- msg
	}
}
