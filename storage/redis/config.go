package redis

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func NewConfigFromEnv() (Config, error) {
	c := Config{
		Addr:     "localhost:6379",
		Username: "redis",
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	if v := os.Getenv("REDIS_ADDR"); v != "" {
		c.Addr = v
	}

	if v := os.Getenv("REDIS_USERNAME"); v != "" {
		c.Password = v
	}

	if v := os.Getenv("REDIS_PASSWORD"); v != "" {
		c.Password = v
	}

	if v := os.Getenv("REDIS_DB"); v != "" {
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return c, fmt.Errorf("invalid config value for redis DB: '%v'", v)
		}
		c.DB = int(val)
	}

	return c, nil
}
