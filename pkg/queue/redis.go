package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Queue struct {
	client *redis.Client
	ctx    context.Context
	list   string
}

func NewRedisQueue(host string, password string, db int, list string) (*Queue, error) {
	q := &Queue{
		ctx: context.Background(),
		client: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password, // no password set
			DB:       db,       // use default DB
		}),
		list: list,
	}
	return q, q.client.Ping(q.ctx).Err()
}

func (q *Queue) Receive() (string, error) {
	return q.client.RPop(q.ctx, q.list).Result()
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
