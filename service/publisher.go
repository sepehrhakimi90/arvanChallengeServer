package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"

	"github.com/sepehrhakimi90/arvanChallengeServer/entity"
	"github.com/sepehrhakimi90/arvanChallengeServer/utils"
)

var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
	redisChannel = os.Getenv("REDIS_CHANNEL")
)

type Publisher interface {
	Publish(rule *entity.Rule) error
	Close()
}

type redisPublisher struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisPublisher(ctx context.Context) Publisher {
	return &redisPublisher{
		rdb: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
			Password: "",
			DB: 0,
		},
		),
		ctx: ctx,
	}
}

func (r *redisPublisher) Publish(rule *entity.Rule) error{
	data, err := json.Marshal(&rule.RuleData)
	if err != nil {
		utils.LogError("redisPublisher", "Publish", err)
		return err
	}
	err = r.rdb.Publish(r.ctx, redisChannel, string(data)).Err()
	if err != nil {
		utils.LogError("redisPublisher", "Publish", err)
		return err
	}
	return nil
}

func (r *redisPublisher) Close() {
	panic("implement me")
}
