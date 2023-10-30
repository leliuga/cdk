package requestid

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// ConfigDefault is the default config
// It uses a fast UUID generator which will expose the number of
// requests made to the server. To conceal this value for better
// privacy, use the "utils.UUIDv4" generator.
var (
	ConfigDefault = Config{
		Next:      nil,
		Header:    fiber.HeaderXRequestID,
		Generator: utils.UUID,
	}
)

// Helper function to set default values
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	c := config[0]

	if c.Header == "" {
		c.Header = ConfigDefault.Header
	}

	if c.Generator == nil {
		c.Generator = ConfigDefault.Generator
	}

	return c
}
