package requestid

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// Config defines the config for middleware.
	Config struct {
		// Next defines a function to skip this middleware when returned true.
		//
		// Optional. Default: nil
		Next func(c *fiber.Ctx) bool

		// Header is the header key where to get/set the unique request ID
		//
		// Optional. Default: "X-Request-ID"
		Header string

		// Generator defines a function to generate the unique identifier.
		//
		// Optional. Default: utils.UUID
		Generator func() string
	}
)
