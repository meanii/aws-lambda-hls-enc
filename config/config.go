package config

import (
	"net/url"
	"os"
	"time"
)

type config struct {
	Origin           string
	KeyID            string
	PrivateKeyBase64 string
	ExpireTime       time.Duration
}

var ConfigEnv *config

// InitConfig initialize the configuration
func InitConfig() {
	ConfigEnv = load()
}

// load configuration from environment variables
func load() (c *config) {
	origin, _ := url.Parse(os.Getenv("CLOUDFRONT_ORIGIN"))
	expireTime, _ := time.ParseDuration(os.Getenv("EXPIRE_TIME"))
	c = &config{
		Origin:           origin.Host,
		KeyID:            os.Getenv("CLOUDFRONT_ACCESS_KEY_ID"),
		PrivateKeyBase64: os.Getenv("CLOUDFRONT_PRIVATE_KEY_BASE64"),
		ExpireTime:       expireTime,
	}
	return
}
