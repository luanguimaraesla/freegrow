package cache

import (
	"fmt"
	"time"

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

// Setex sets an entry with TTL
func Setex(key string, value interface{}, duration time.Duration) error {
	c, err := Connect()
	if err != nil {
		return err
	}

	defer c.Close()

	if _, err := c.Do("SETEX", key, fmt.Sprintf("%.0f", duration.Seconds()), value); err != nil {
		return err
	}

	return nil
}

// GetString gets an string from cache
func GetString(key string) (string, error) {
	c, err := Connect()
	if err != nil {
		return "", err
	}

	defer c.Close()

	return redis.String(c.Do("GET", key))
}

// Renew set a new TTL for a key
func Renew(key string, duration time.Duration) error {
	c, err := Connect()
	if err != nil {
		return err
	}

	defer c.Close()

	if _, err := c.Do("EXPIRE", key, fmt.Sprintf("%.0f", duration.Seconds())); err != nil {
		return err
	}

	return nil
}
