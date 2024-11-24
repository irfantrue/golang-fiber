package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func Compress() fiber.Handler {
	return compress.New(compress.Config{
		Next:  nil,
		Level: 1,
	})
}
