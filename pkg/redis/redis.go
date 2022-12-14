package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/haoran-mc/gcbbs/pkg/config"
)

var RD *redis.Client

func init() {
	r := config.Conf.Redis
	RD = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, gconv.Int(r.Port)),
		Password: r.Pass,
		DB:       gconv.Int(r.DB),
		PoolSize: 10,
	})
	if str, err := RD.Ping(context.Background()).Result(); err != nil || str != "PONG" {
		log.Fatalf("redis connect ping failed, err: %v", err)
	}
}
