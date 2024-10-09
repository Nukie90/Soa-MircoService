package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
)

func AuthenticateUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Add user ID to context
	c.Locals("userId", claims["user_id"])
	return c.Next()
}
