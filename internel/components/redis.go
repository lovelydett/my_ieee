package components

import (
	"github.com/go-redis/redis"
)

// _RedisClient is a wrapper around the redis client
type RedisClient struct {
	Client   *redis.Client
	address  string
	password string
	db       int
}

var redisClient = &RedisClient{nil, "", "", 0}

// Func GetRedisClient creates a new redis client (or reuse the existing one)
func GetRedisClient(address, password string, db int) (*RedisClient, error) {
	if address == redisClient.address && password == redisClient.password && db == redisClient.db {
		return redisClient, nil
	}

	// create a new redis client
	redisClient.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	client := redisClient.Client
	// ping the server to check if the connection is established
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
