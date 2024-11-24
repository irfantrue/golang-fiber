package config

import (
	"golang-fiber/internal/middleware"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.App {
	return *fiber.New(fiber.Config{
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		ServerHeader:  "GolangFiberAPI/1.0",
		AppName:       "GolangFiberAPI/1.0",
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   30 * time.Second,
		BodyLimit:     10 * 1024 * 1024,
		ErrorHandler:  middleware.ErrorHandler,
		// Prefork:       true, // Multicore
	})
}
