package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/rs/zerolog"
)

type GoMiddleware struct {
	appCtx *fiber.App
	appLog *zerolog.Logger
}

// CORS Middleware
func (m *GoMiddleware) CORS() fiber.Handler {
	crs := os.Getenv("CORS_WHITELISTS")

	if crs == "*" {
		return cors.New(cors.Config{
			AllowOrigins:  "*",
			AllowHeaders:  "Content-Type, Accept, Authorization",
			AllowMethods:  "GET, HEAD, PUT, PATCH, POST, DELETE",
			ExposeHeaders: "*",
		})
	}

	return cors.New(cors.Config{
		AllowOrigins:     crs,
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
		ExposeHeaders:    "*",
	})
}

func (m *GoMiddleware) LOGGER() fiber.Handler {
	// Inisialisasi logger jika belum tersedia
	if m.appLog == nil {
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		m.appLog = &logger
	}

	return func(c *fiber.Ctx) error {
		startTime := time.Now()
		body := c.Body()

		// Lanjutkan ke handler berikutnya
		err := c.Next()

		// Hitung durasi response
		responseTime := time.Since(startTime)

		event := m.appLog.Info()

		if err != nil {
			event = m.appLog.Error().Err(err)
		}

		event.
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("ip", c.IP()).
			Str("user_agent", c.Get("User-Agent")).
			Str("payload", string(body)).
			Interface("params", c.Queries()).
			Int("status", c.Response().StatusCode()).
			Dur("response_time", responseTime).
			Msg("HTTP request log")

		return err
	}
}

// Rate Limiter Middleware
func (m *GoMiddleware) RateLimiter() fiber.Handler {
	limiterCfg := limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			authKey := c.Get("Authorization")
			if authKey == "" {
				return c.IP() + c.Get("User-Agent")
			}
			return authKey
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}
	return limiter.New(limiterCfg)
}

func (m *GoMiddleware) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header Authorization
		authHeader := c.Get("Authorization")

		// Pastikan formatnya "Bearer <TOKEN>"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "Missing or malformed token",
			})
		}

		// Ambil token setelah "Bearer "
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Bandingkan dengan SECRET_KEY di environment
		expectedToken := os.Getenv("SECRET_KEY")
		if token != expectedToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "Invalid token",
			})
		}

		// Tambahkan header tambahan
		c.Set("X-Service-Name", "AuthService")

		// Lanjutkan ke handler berikutnya
		return c.Next()
	}
}

// Init Middleware
func InitMiddleware(ctx *fiber.App, appLog *zerolog.Logger) *GoMiddleware {
	return &GoMiddleware{appCtx: ctx, appLog: appLog}
}
