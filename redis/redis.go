package chremoas_redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Client struct {
	Client *redis.Client
	Prefix string
}

type Servers struct {
	Host      string
	Sentinels []string
}

var Nil = redis.Nil

func Init(servers Servers, password string, db int, prefix string) *Client {
	var c *redis.Client

	if servers.Sentinels != nil {
		c = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "master",
			SentinelAddrs: servers.Sentinels,
		})
	} else {
		c = redis.NewClient(&redis.Options{
			Addr:     servers.Host,
			Password: password, // no password set
			DB:       db,       // use default DB
		})
	}

	return &Client{Client: c, Prefix: prefix}
}

func (r Client) KeyName(key string) string {
	return fmt.Sprintf("%s:%s", r.Prefix, key)
}
