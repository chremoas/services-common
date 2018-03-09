package chremoas_redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type Client struct {
	Client *redis.Client
	Prefix string
}

var Nil = redis.Nil

func Init(addr string, password string, db int, prefix string) *Client {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,  // use default DB
	})

	return &Client{Client: c, Prefix: prefix}
}

func (r Client) Ping() (string, error) {
	result, err := r.Client.Ping().Result()
	return result, err
}

func (r Client) Get(key string, value interface{}) error {
	prekey := fmt.Sprintf("%s_%s", r.Prefix, key)
	result, err := r.Client.Get(prekey).Result()

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(result), value)

	return nil
}

func (r Client) Set(key string, value interface{}) error {
	prekey := fmt.Sprintf("%s_%s", r.Prefix, key)

	result, _ := json.Marshal(value)

	err := r.Client.Set(prekey, result, 0).Err()
	return err
}
