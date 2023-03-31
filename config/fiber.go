package config

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"net/http"
	"strings"
	"time"
)

var allowOrigin = []string{
	"http://localhost:3000",
	"http://localhost:3001",
	"*",
}

var allowHeaders = []string{
	"Content-Type",
	"Content-Length",
	"Accept-Encoding",
	"X-CSRF-Token",
	"Authorization",
	"Origin",
	"Cache-Control",
	"X-Requested-With",
	"Authorization",
	"Accept",
}

func (c *Config) InitFiber() {
	log.Println("Init app engine . . .")

	fiberSingleton.Do(func() {
		const timeoutDuration = 10

		engine := fiber.New(fiber.Config{
			ReadTimeout: time.Second * time.Duration(timeoutDuration),
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError

				var e *fiber.Error
				if errors.As(err, &e) {
					code = e.Code
				}

				return ctx.Status(code).JSON(fiber.Map{
					"error":   true,
					"code":    code,
					"message": http.StatusText(code),
					"data":    err.Error(),
				})
			},
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		})

		engine.Use(
			cors.New(cors.Config{
				AllowOrigins:     strings.Join(allowOrigin, ","),
				AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
				AllowHeaders:     strings.Join(allowHeaders, ","),
				AllowCredentials: true,
			}),
			logger.New(),
			recover.New(),
		)

		FiberApp = engine

		log.Printf("App engine: ready (gofiber v%s)", fiber.Version)
	})
}
