package cache

import (
	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

// Init starts a new redis pool
func Init(redisURL string) {
	pool = newPool(redisURL)
}

// newPool defines a new redis pool
func newPool(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(url)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

// Connect creates a new redis pool connection member
func Connect() (redis.Conn, error) {
	conn := pool.Get()

	if err := conn.Err(); err != nil {
		return conn, err
	}

	return conn, nil
}

// Ping tests redis connection
func Ping() error {
	c, err := Connect()
	defer c.Close()

	return err
}
